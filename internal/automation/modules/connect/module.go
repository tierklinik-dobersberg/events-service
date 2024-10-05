package connect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/dop251/goja"
	"github.com/hashicorp/go-multierror"
	"github.com/tierklinik-dobersberg/apis/pkg/cli"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
)

type ConnectModule struct{}

func (*ConnectModule) Name() string { return "connect" }

func (*ConnectModule) NewModuleInstance(vu modules.VU) (*goja.Object, error) {
	cfg := vu.Config()

	merr := new(multierror.Error)

	if cfg.IdmURL != "" {
		makeServiceClient("users", vu, cfg.IdmURL, "tkd.idm.v1.UserService", merr)
		makeServiceClient("roles", vu, cfg.IdmURL, "tkd.idm.v1.RoleService", merr)
		makeServiceClient("notify", vu, cfg.IdmURL, "tkd.idm.v1.NotiyService", merr)
	}

	if cfg.RosterURL != "" {
		makeServiceClient("roster", vu, cfg.RosterURL, "tkd.roster.v1.RosterService", merr)
		makeServiceClient("offtime", vu, cfg.RosterURL, "tkd.roster.v1.OffTimeService", merr)
		makeServiceClient("workshift", vu, cfg.RosterURL, "tkd.roster.v1.WorkShiftService", merr)
	}

	if cfg.TaskServiceURL != "" {
		makeServiceClient("tasks", vu, cfg.TaskServiceURL, "tkd.tasks.v1.TaskService", merr)
		makeServiceClient("boards", vu, cfg.TaskServiceURL, "tkd.tasks.v1.BoardService", merr)
	}

	return nil, merr.ErrorOrNil()
}

func makeServiceClient(pkgname string, vu modules.VU, ep string, serviceName string, merr *multierror.Error) {
	serviceObj := vu.Runtime().NewObject()

	d, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(serviceName))
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
			endpoint: ep + "/" + string(desc.FullName()) + "/" + string(mdesc.Name()),
			request:  mdesc.Input(),
			response: mdesc.Output(),
			cli:      cli.NewInsecureHttp2Client(),
		}

		serviceObj.Set(methodName, cli.do)
	}

	vu.Runtime().Set(pkgname, serviceObj)
}

type client struct {
	endpoint string
	request  protoreflect.MessageDescriptor
	response protoreflect.MessageDescriptor

	cli *http.Client
}

func (c *client) do(in *goja.Object) (any, error) {
	payload, err := ObjectToProto(in, c.request)
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

	// create a goja value

	protoBlob, err := protojson.Marshal(res)
	if err != nil {
		return nil, err
	}

	m := make(map[string]any)
	if err := json.Unmarshal(protoBlob, &m); err != nil {
		return nil, err
	}

	return m, nil
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
