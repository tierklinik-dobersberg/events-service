package automation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/dop251/goja"
	"github.com/tierklinik-dobersberg/apis/pkg/cli"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

type client struct {
	endpoint string
	request  protoreflect.MessageDescriptor
	response protoreflect.MessageDescriptor

	cli *http.Client
}

func WithConnectService(name, ep string, desc protoreflect.ServiceDescriptor) EngineOption {
	return func(e *Engine) {
		m := make(map[string]any)

		for mi := 0; mi < desc.Methods().Len(); mi++ {
			mdesc := desc.Methods().Get(mi)

			methodName := strings.ToLower(string(mdesc.Name()[0])) + string(mdesc.Name()[1:])

			cli := &client{
				endpoint: ep + "/" + string(desc.FullName()) + "/" + string(mdesc.Name()),
				request:  mdesc.Input(),
				response: mdesc.Output(),
				cli:      cli.NewInsecureHttp2Client(),
			}

			m[methodName] = cli.do
		}

		e.RegisterNativeModuleHelper(name, m)
	}
}

func (c *client) do(in *goja.Object) (any, error) {
	payload, err := objToProto(in, c.request)
	if err != nil {
		return nil, err
	}

	blob, err := protojson.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.endpoint, bytes.NewReader(blob))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	slog.Info("sending connect RPC request", "endpoint", c.endpoint)

	response, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %s", response.Status)
	}

	res := dynamicpb.NewMessage(c.response)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if err := protojson.Unmarshal(body, res); err != nil {
		return nil, err
	}

	slog.Info("received response from connect service", "body", fmt.Sprintf("%+v", res))

	return res, nil
}

func objToProto(in *goja.Object, out protoreflect.MessageDescriptor) (proto.Message, error) {
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
