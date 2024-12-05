package connect

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/dop251/goja"
	"github.com/hashicorp/go-multierror"
	"github.com/tierklinik-dobersberg/apis/pkg/cli"
	"github.com/tierklinik-dobersberg/apis/pkg/discovery"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/common"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
	"github.com/tierklinik-dobersberg/pbtype-server/pkg/protoresolve"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

type ConnectModule struct{}

func (*ConnectModule) Name() string { return "services" }

func (*ConnectModule) NewModuleInstance(vu modules.VU) (*goja.Object, error) {
	cfg := vu.Config()

	exports := vu.Runtime().NewObject()

	if cfg.TypeServerURL == "" {
		return nil, nil
	}

	merr := new(multierror.Error)

	for _, serviceName := range cfg.ConnectServices {

		serviceParts := strings.Split(serviceName, ".")

		jsServiceName := strings.ToLower(
			strings.TrimSuffix(serviceParts[len(serviceParts)-1], "Service"),
		)

		slog.Info("creating connect service", "jsModule", jsServiceName, "service-name", serviceName)
		makeServiceClient(vu.Discoverer(), vu.TypeResolver(), jsServiceName, exports, vu, serviceName, merr)
	}

	return exports, merr.ErrorOrNil()
}

func makeServiceClient(disc discovery.Discoverer, resolver protoresolve.Resolver, pkgname string, obj *goja.Object, vu modules.VU, serviceName string, merr *multierror.Error) {
	serviceObj := vu.Runtime().NewObject()

	d, err := resolver.FindDescriptorByName(protoreflect.FullName(serviceName))
	if err != nil {
		merr.Errors = append(merr.Errors, fmt.Errorf("%s: %w", serviceName, err))
		return
	}

	desc, ok := d.(protoreflect.ServiceDescriptor)
	if !ok {
		merr.Errors = append(merr.Errors, fmt.Errorf("%s: expected a service descriptor but got %T", serviceName, d))
		return
	}

	for mi := 0; mi < desc.Methods().Len(); mi++ {
		mdesc := desc.Methods().Get(mi)

		// Streaming is not yet supported
		if mdesc.IsStreamingClient() || mdesc.IsStreamingServer() {
			continue
		}

		methodName := strings.ToLower(string(mdesc.Name()[0])) + string(mdesc.Name()[1:])

		cli := &client{
			service:  string(desc.FullName()),
			method:   string(mdesc.Name()),
			disc:     disc,
			request:  mdesc.Input(),
			response: mdesc.Output(),
			cli:      cli.NewInsecureHttp2Client(),
			rt:       vu.Runtime(),
		}

		serviceObj.Set(methodName, cli.do)
	}

	obj.Set(pkgname, serviceObj)
}

type client struct {
	service  string
	method   string
	disc     discovery.Discoverer
	request  protoreflect.MessageDescriptor
	response protoreflect.MessageDescriptor

	rt *goja.Runtime

	cli *http.Client
}

func (cli *client) resolveEndpoint() (string, error) {
	parts := strings.Split(cli.service, ".")

	slog.Info("resolving service instance", "service", cli.service)

	queries := make([]string, 0, len(parts))
	for idx := 1; idx <= len(parts); idx++ {
		queries = append(queries, strings.Join(parts[:idx], "."))
	}
	slices.Reverse(queries)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, q := range queries {
		res, err := cli.disc.Discover(ctx, q)
		if err == nil && len(res) > 0 {
			ep := fmt.Sprintf("http://%s/%s/%s", res[0].Address, cli.service, cli.method)
			return ep, nil
		}

		slog.Info("failed to resolve service instance", "query", q, "error", err)
	}

	return "", fmt.Errorf("failed to find a healthy service instance")
}

func (c *client) do(in *goja.Object, options *goja.Object) any {
	payload, err := ObjectToProto(in, c.request)
	if err != nil {
		common.Throw(c.rt, err)
	}

	blob, err := protojson.Marshal(payload)
	if err != nil {
		common.Throw(c.rt, err)
	}

	ep, err := c.resolveEndpoint()
	if err != nil {
		common.Throw(c.rt, err)
	}

	req, err := http.NewRequest(http.MethodPost, ep, bytes.NewReader(blob))
	if err != nil {
		common.Throw(c.rt, err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if options != nil {
		slog.Info("preparing custom HTTP headers")
		headerObj := options.Get("headers")
		if headerObj != nil {
			if headers, ok := headerObj.(*goja.Object); ok {
				for _, key := range headers.Keys() {
					headerValue := headers.Get(key)
					val := headerValue.Export()

					switch v := val.(type) {
					case string:
						req.Header.Add(key, v)

					case []any:
						for _, el := range v {
							if s, ok := el.(string); ok {
								req.Header.Add(key, s)
							} else {
								slog.Warn("ignoring custom header value", "header", key, "value", el)
							}
						}

					default:
						slog.Warn("ignoring custom header value", "header", key, "value", val)
					}
				}
			} else {
				slog.Warn("ignoring custom request headers, unsupported type", "type", fmt.Sprintf("%T", headerObj))
			}
		}

	}

	response, err := c.cli.Do(req)
	if err != nil {
		common.Throw(c.rt, err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		err := fmt.Errorf("unexpected status code for %s %s: %s", req.Method, req.URL.String(), response.Status)
		common.Throw(c.rt, err)
	}

	res := dynamicpb.NewMessage(c.response)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		common.Throw(c.rt, err)
	}

	if err := protojson.Unmarshal(body, res); err != nil {
		common.Throw(c.rt, err)
	}

	// create a goja value

	protoBlob, err := protojson.Marshal(res)
	if err != nil {
		common.Throw(c.rt, err)
	}

	m := make(map[string]any)
	if err := json.Unmarshal(protoBlob, &m); err != nil {
		common.Throw(c.rt, err)
	}

	return m
}

func ObjectToProto(in *goja.Object, out protoreflect.MessageDescriptor) (proto.Message, error) {
	msg := dynamicpb.NewMessage(out)

	blob, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("failed to convert goja.Object to JSON: %w", err)
	}

	if err := protojson.Unmarshal(blob, msg); err != nil {
		return nil, fmt.Errorf("failed to convert goja.Object to proto.Message: (%s) %w", string(blob), err)
	}

	return msg, nil
}

func ConvertProtoMessage(msg proto.Message) (any, error) {
	jsonBlob, err := protojson.Marshal(msg)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(jsonBlob, &result); err != nil {
		return nil, err
	}

	return result, nil
}
