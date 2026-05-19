package fetch

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/dop251/goja"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/common"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
)

type Module struct{}

func (*Module) Name() string { return "fetch" }

func (*Module) NewModuleInstance(vu modules.VU) (*goja.Object, error) {
	vm := vu.Runtime()

	vm.Set("fetch", func(url string, options map[string]any) goja.Value {
		method := "GET"
		if m, ok := options["method"].(string); ok && m != "" {
			method = m
		}

		var body io.Reader
		if b, ok := options["body"].(string); ok {
			body = strings.NewReader(b)
		}

		req, err := http.NewRequest(method, url, body)
		if err != nil {
			common.Throw(vm, err)
		}

		if headers, ok := options["headers"].(map[string]any); ok {
			for key, hval := range headers {
				switch v := hval.(type) {
				case []any:
					for _, anyVal := range v {
						if s, ok := anyVal.(string); ok {
							req.Header.Add(key, s)
						}
					}
				case string:
					req.Header.Add(key, v)
				}
			}
		}

		if user, ok := options["user"].(string); ok {
			password, _ := options["password"].(string)

			req.SetBasicAuth(user, password)
		}

		promise, resolve, reject := vm.NewPromise()

		go func() {
			vu.Log().Log(context.Background(), slog.LevelInfo, "fetch: sending request", "url", req.URL.String(), "headers", req.Header)

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				reject(err)
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				reject(fmt.Errorf("failed to read response body: %w", err))
				return
			}

			// prepare response headers
			resHeaders := map[string]any{}
			for k, v := range resp.Header {
				resHeaders[k] = v
			}

			resolve(map[string]any{
				"ok":         resp.StatusCode >= 200 && resp.StatusCode < 300,
				"status":     resp.StatusCode,
				"statusText": resp.Status,
				"url":        resp.Request.URL.String(),
				"headers":    resHeaders,
				"text":       func() string { return string(body) },
				"json": func() (any, error) {
					var data any
					err := json.Unmarshal(body, &data)
					return data, err
				},
				"arrayBuffer": func() any { return vm.NewArrayBuffer(body) },
				"bytes": func() (any, error) {
					return vm.New(vm.Get("Uint8Array"), vm.ToValue(vm.NewArrayBuffer(body)))
				},
			})
		}()

		return vm.ToValue(promise)
	})

	return nil, nil
}
