package codec

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
)

type Resolver interface {
	protoregistry.ExtensionTypeResolver
	protoregistry.MessageTypeResolver
}

type CustomJSONCodec struct {
	resolver Resolver
}

func NewCustomJSONCodec(resolver Resolver) *CustomJSONCodec {
	if resolver == nil {
		resolver = protoregistry.GlobalTypes
	}

	return &CustomJSONCodec{
		resolver: resolver,
	}
}

func errNotProto(m any) error {
	return fmt.Errorf("expected a proto.Message but got %T", m)
}

var _ connect.Codec = (*CustomJSONCodec)(nil)

func (c *CustomJSONCodec) Name() string { return "json" }

func (c *CustomJSONCodec) Marshal(message any) ([]byte, error) {
	protoMessage, ok := message.(proto.Message)
	if !ok {
		return nil, errNotProto(message)
	}
	return protojson.MarshalOptions{
		Resolver: c.resolver,
	}.Marshal(protoMessage)
}

func (c *CustomJSONCodec) MarshalAppend(dst []byte, message any) ([]byte, error) {
	protoMessage, ok := message.(proto.Message)
	if !ok {
		return nil, errNotProto(message)
	}
	return protojson.MarshalOptions{
		Resolver: c.resolver,
	}.MarshalAppend(dst, protoMessage)
}

func (c *CustomJSONCodec) Unmarshal(binary []byte, message any) error {
	protoMessage, ok := message.(proto.Message)
	if !ok {
		return errNotProto(message)
	}
	if len(binary) == 0 {
		return errors.New("zero-length payload is not a valid JSON object")
	}
	// Discard unknown fields so clients and servers aren't forced to always use
	// exactly the same version of the schema.
	options := protojson.UnmarshalOptions{DiscardUnknown: true, Resolver: c.resolver}
	return options.Unmarshal(binary, protoMessage)
}

func (c *CustomJSONCodec) MarshalStable(message any) ([]byte, error) {
	// protojson does not offer a "deterministic" field ordering, but fields
	// are still ordered consistently by their index. However, protojson can
	// output inconsistent whitespace for some reason, therefore it is
	// suggested to use a formatter to ensure consistent formatting.
	// https://github.com/golang/protobuf/issues/1373
	messageJSON, err := c.Marshal(message)
	if err != nil {
		return nil, err
	}
	compactedJSON := bytes.NewBuffer(messageJSON[:0])
	if err = json.Compact(compactedJSON, messageJSON); err != nil {
		return nil, err
	}
	return compactedJSON.Bytes(), nil
}

func (c *CustomJSONCodec) IsBinary() bool {
	return false
}
