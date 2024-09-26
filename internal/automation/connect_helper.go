package automation

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/bufbuild/connect-go"
	"github.com/dop251/goja"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func withProtoMessage[T proto.Message, R any](fn func(T) (R, error)) func(obj *goja.Object) (R, error) {
	return func(obj *goja.Object) (R, error) {
		var (
			params T
			result R
		)

		pType := reflect.TypeOf(params).Elem()
		value := reflect.New(pType).Interface()

		m := value.(proto.Message)

		blob, err := json.Marshal(obj)
		if err != nil {
			return result, fmt.Errorf("failed to convert goja.Object to JSON: %w", err)
		}

		if err := protojson.Unmarshal(blob, m); err != nil {
			return result, fmt.Errorf("failed to convert goja.Object to proto.Message: %w", err)
		}

		return fn(m.(T))
	}
}

func convertProtoMessage(msg proto.Message) (any, error) {
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

type UnaryMethod[T any, R any] func(context.Context, *connect.Request[T]) (*connect.Response[R], error)

func wrapConnectMethod[T any, R any](fn UnaryMethod[T, R]) func(obj *goja.Object) (any, error) {
	return func(obj *goja.Object) (any, error) {
		var (
			params T
		)

		pType := reflect.TypeOf(params)
		value := reflect.New(pType).Interface()

		m := value.(proto.Message)

		blob, err := json.Marshal(obj)
		if err != nil {
			return nil, fmt.Errorf("failed to convert goja.Object to JSON: %w", err)
		}

		log.Printf("trying to unmashal blob to message of type %v", m)
		log.Println(string(blob))

		if err := protojson.Unmarshal(blob, m); err != nil {
			return nil, fmt.Errorf("failed to convert goja.Object to proto.Message: (%s) %w", string(blob), err)
		}

		res, err := fn(context.Background(), connect.NewRequest(value.(*T)))
		if err != nil {
			return nil, nil
		}

		var rpb any = res.Msg
		jsonBlob, err := protojson.Marshal(rpb.(proto.Message))
		if err != nil {
			return nil, err
		}

		var result map[string]any
		if err := json.Unmarshal(jsonBlob, &result); err != nil {
			return nil, err
		}

		return result, nil
	}
}
