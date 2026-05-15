package provet

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tierklinik-dobersberg/events-service/internal/automation"
	"github.com/tierklinik-dobersberg/events-service/internal/automation/modules"
	"github.com/tierklinik-dobersberg/events-service/internal/config"
	"github.com/tierklinik-dobersberg/provet-go/provet"
)

type fakeClient struct{}

func str(s string) *string {
	return &s
}

type testLogger struct{}

// Log logs an info level message. It implements the console.Printer interface.
func (*testLogger) Log(msg string) {
	log.Printf("INFO: %s", msg)
}

// Warn logs an warn level message. It implements the console.Printer interface.
func (*testLogger) Warn(msg string) {
	log.Printf("WARN: %s", msg)
}

// Error logs an error level message. It implements the console.Printer interface.
func (*testLogger) Error(msg string) {
	log.Printf("ERROR: %s", msg)
}

func (*fakeClient) Do(req *http.Request) (*http.Response, error) {
	log.Printf("doing request: %s", req.URL.String())

	blob, _ := json.Marshal(provet.PaginatedClientList{
		Results: []provet.Client{
			{
				Firstname: str("Musterfrau"),
				ZipCode:   str("3843"),
			},
		},
	})

	return &http.Response{
		StatusCode: 200,
		Status:     "Test",
		Body:       io.NopCloser(bytes.NewReader(blob)),
	}, nil
}

func TestProvetModule(t *testing.T) {
	m := &Module{}
	reg := &modules.Registry{}
	reg.Register(m)

	rt, err := automation.New("test", config.Config{
		ProvetID:           1,
		ProvetClientID:     "test",
		ProvetClientSecret: "test",
	}, nil, automation.WithModulsRegistry(reg), automation.WithConsole(&testLogger{}))

	m.client.Client = &fakeClient{}

	require.NoError(t, err)

	_, err = rt.RunScript(`
		const provet = require("provet")

		var res = provet.clientList({
			zip_code__in: "3843,3842",
			firstname__is: "Musterfrau",
		})

		if (res.results[0].firstname != "Musterfrau" || res.results[0].zip_code != "3843") {
			throw new Error("unexpected result")
		}
	`)

	require.NoError(t, err)
}
