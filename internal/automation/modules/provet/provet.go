package provet

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"reflect"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/dop251/goja"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
	"github.com/tierklinik-dobersberg/provet-go/provet"
)

type Module struct {
	client *provet.ProvetClient
}

func (*Module) Name() string { return "provet" }

func (m *Module) NewModuleInstance(vu modules.VU) (*goja.Object, error) {
	// Do nothing if Provet is not configured
	cfg := vu.Config()
	if cfg.ProvetID == 0 || cfg.ProvetClientID == "" || cfg.ProvetClientSecret == "" {
		return nil, nil
	}

	client, err := provet.NewWithClientCredentials(context.Background(), cfg.ProvetID, cfg.ProvetClientID, cfg.ProvetClientSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create provet API client: %w", err)
	}
	m.client = client.ClientInterface.(*provet.ProvetClient)

	clientType := reflect.ValueOf(m.client).Type()

	exports := vu.Runtime().NewObject()

	for idx := range clientType.NumMethod() {
		method := clientType.Method(idx)

		// only use methods that return *http.Response and accept a request struct.
		if strings.HasSuffix(method.Name, "WithResponse") || strings.HasSuffix(method.Name, "WithBody") || strings.HasSuffix(method.Name, "WithFormdataBody") {
			continue
		}

		if err := createMethod(vu, exports, m.client, method); err != nil {
			return nil, fmt.Errorf("")
		}
	}

	return exports, nil
}

func createMethod(vu modules.VU, obj *goja.Object, client *provet.ProvetClient, m reflect.Method) error {
	rt := vu.Runtime()

	obj.Set(jsFuncName(m.Name), func(args ...goja.Value) (goja.Value, error) {
		// make sure we got enough arguments
		if expected := m.Type.NumIn() - 3; len(args) < expected {
			return nil, fmt.Errorf("missing required parameters, got %d but expected %d", len(args), expected)
		}

		var callArgs []reflect.Value
		for mIdx := range m.Type.NumIn() {
			argIdx := mIdx - 2
			if mIdx <= 1 {
				// self and context.Context
				continue
			}
			if mIdx == m.Type.NumIn()-1 {
				// last one is variadic request-editor functions
				break
			}

			inputParam := m.Type.In(mIdx)

			// directly assign if that's possisble
			if args[argIdx].ExportType().AssignableTo(inputParam) {
				callArgs = append(callArgs, reflect.ValueOf(args[argIdx].Export()))
				continue
			}

			wasPointer := false
			if inputParam.Kind() == reflect.Ptr {
				wasPointer = true
				inputParam = inputParam.Elem()
			}

			// we only support structs for now
			if inputParam.Kind() != reflect.Struct {
				return nil, fmt.Errorf("parameter type %s is not yet supported", inputParam.Kind().String())
			}

			objArg, ok := args[argIdx].(*goja.Object)
			if !ok {
				return nil, fmt.Errorf("expected an object as parameter %d but got %s", argIdx, args[argIdx].ExportType().Name())
			}

			// marshal goja object as JSON
			blob, err := json.Marshal(objArg)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal parameter %d as JSON: %w", argIdx, err)
			}

			// unmarshal back into the expected parameter type
			p := reflect.New(inputParam)
			if err := json.Unmarshal(blob, p.Interface()); err != nil {
				return nil, fmt.Errorf("failed to unmarshal parameter %d into %s: %w", argIdx, inputParam.Name(), err)
			}

			spew.Dump(p.Interface())

			if wasPointer {
				callArgs = append(callArgs, p)
			} else {
				callArgs = append(callArgs, p.Elem())
			}
		}

		promise, resolve, reject := rt.NewPromise()

		go func() {
			defer func() {
				if x := recover(); x != nil {
					vu.Log().Log(context.Background(), slog.LevelError, "captured panic in provet API call", "reason", x)
				}
			}()
			// finnally, call the method
			out := m.Func.Call(append([]reflect.Value{
				reflect.ValueOf(client),
				reflect.ValueOf(context.Background()),
			}, callArgs...))

			// check returned arguments
			if !out[1].IsNil() {
				err, ok := out[1].Interface().(error)
				if !ok {
					rejectErr := fmt.Errorf("expected second return value to be error, got %T", out[1].Interface())
					if err := reject(rejectErr); err != nil {
						vu.Log().Log(context.Background(), slog.LevelError, "failed to reject provet API promise", "error", err)
					}
				}

				if err != nil {
					rejectErr := err
					if err := reject(rejectErr); err != nil {
						vu.Log().Log(context.Background(), slog.LevelError, "failed to reject provet API promise", "error", err)
					}
				}

				return
			}

			resp, ok := out[0].Interface().(*http.Response)
			if !ok {
				rejectErr := fmt.Errorf("expected first return value to be *http.Response, got %T", out[0].Interface())
				if err := reject(rejectErr); err != nil {
					vu.Log().Log(context.Background(), slog.LevelError, "failed to reject provet API promise", "error", err)
				}

				return
			}

			// check the status code
			if sc := resp.StatusCode; sc < 200 || sc >= 300 {
				rejectErr := fmt.Errorf("got unexpected status code %d: %s", sc, resp.Status)
				if err := reject(rejectErr); err != nil {
					vu.Log().Log(context.Background(), slog.LevelError, "failed to reject provet API promise", "error", err)
				}

				return
			}

			defer resp.Body.Close()

			content, err := io.ReadAll(resp.Body)
			if err != nil {
				rejectErr := fmt.Errorf("failed to read provet response body: %w", err)
				if err := reject(rejectErr); err != nil {
					vu.Log().Log(context.Background(), slog.LevelError, "failed to reject provet API promise", "error", err)
				}

				return
			}

			var result any
			if err := json.Unmarshal(content, &result); err != nil {
				rejectErr := fmt.Errorf("failed to unmarshal provet response body as JSON: %w", err)
				if err := reject(rejectErr); err != nil {
					vu.Log().Log(context.Background(), slog.LevelError, "failed to reject provet API promise", "error", err)
				}

				return
			}

			resolve(rt.ToValue(result))
		}()

		return rt.ToValue(promise), nil
	})

	return nil
}

func jsFuncName(name string) string {
	return strings.ToLower(name[:1]) + name[1:]
}

func init() {
	if err := modules.Register(&Module{}); err != nil {
		panic(err.Error())
	}
}
