package connect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/dop251/goja"
	"github.com/hashicorp/go-multierror"
	"github.com/tierklinik-dobersberg/apis/pkg/cli"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/common"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
	"github.com/tierklinik-dobersberg/pbtype-server/resolver"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

type ConnectModule struct{}

func (*ConnectModule) Name() string { return "connect" }

func (*ConnectModule) NewModuleInstance(vu modules.VU) (*goja.Object, error) {
	cfg := vu.Config()

	if cfg.TypeServerURL == "" {
		return nil, nil
	}

	resolver := resolver.New(cfg.TypeServerURL)

	merr := new(multierror.Error)

	for _, service := range cfg.ConnectServices {
		u, err := url.Parse(service)
		if err != nil {
			merr.Errors = append(merr.Errors, fmt.Errorf("failed to parse service url %q: %w", service, err))
			continue
		}

		parts := strings.Split(u.Path, "/")

		// the last part is expected to be the fully qualified service name
		serviceName := parts[len(parts)-1]
		serviceParts := strings.Split(serviceName, ".")

		jsServiceName := strings.ToLower(
			strings.TrimSuffix(serviceParts[len(serviceParts)-1], "Service"),
		)
		path := strings.Join(parts[:len(parts)-1], "/")

		serviceURL := fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, path)

		slog.Info("creating connect service", "jsModule", jsServiceName, "service-name", serviceName, "service-url", serviceURL)
		makeServiceClient(resolver, jsServiceName, vu, serviceURL, serviceName, merr)
	}

	return nil, merr.ErrorOrNil()
}

func makeServiceClient(resolver *resolver.Resolver, pkgname string, vu modules.VU, ep string, serviceName string, merr *multierror.Error) {
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

		serviceEndpoint := strings.TrimSuffix(ep, "/") + "/" + string(desc.FullName()) + "/" + string(desc.Name())

		cli := &client{
			endpoint: serviceEndpoint,
			request:  mdesc.Input(),
			response: mdesc.Output(),
			cli:      cli.NewInsecureHttp2Client(),
			rt:       vu.Runtime(),
		}

		serviceObj.Set(methodName, cli.do)
	}

	vu.Runtime().Set(pkgname, serviceObj)
}

type client struct {
	endpoint string
	request  protoreflect.MessageDescriptor
	response protoreflect.MessageDescriptor

	rt *goja.Runtime

	cli *http.Client
}

func (c *client) do(in *goja.Object) any {
	payload, err := ObjectToProto(in, c.request)
	if err != nil {
		common.Throw(c.rt, err)
	}

	blob, err := protojson.Marshal(payload)
	if err != nil {
		common.Throw(c.rt, err)
	}

	req, err := http.NewRequest(http.MethodPost, c.endpoint, bytes.NewReader(blob))
	if err != nil {
		common.Throw(c.rt, err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

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
