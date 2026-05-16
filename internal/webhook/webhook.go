package webhook

import (
	"context"
	"fmt"
	"mime"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"

	eventsv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Webhook struct {
	Path                string
	TTL                 time.Duration
	ExpectedContentType string
	MaxContentLength    uint64
	ExpectedHeaders     map[string]string
	IPWhiteList         []net.IPNet

	expectedHeadersParsed map[string]*regexp.Regexp

	CreatedAt  time.Time
	LastUpdate time.Time
}

func (w *Webhook) Prepare() error {
	// validate the webhook path
	if err := ValidateWebhookPath(w.Path); err != nil {
		return fmt.Errorf("invalid webhook path definition: %w", err)
	}

	// check and prepare expected header regular expressions
	if len(w.ExpectedHeaders) > 0 {
		w.expectedHeadersParsed = make(map[string]*regexp.Regexp)

		for header, re := range w.ExpectedHeaders {
			if re == "" {
				w.expectedHeadersParsed[header] = nil
			} else {
				p, err := regexp.CompilePOSIX(re)
				if err != nil {
					return fmt.Errorf("invalid regular expression in expected_http_headers: %q: %w", header, err)
				}

				w.expectedHeadersParsed[header] = p
			}
		}
	}

	return nil
}

func (w Webhook) MatchRequest(ctx context.Context, r *http.Request, body []byte) (event *eventsv1.WebhookEvent, err error) {
	// make sure we ignore expired webhook
	if w.TTL > 0 && time.Now().After(w.LastUpdate.Add(w.TTL)) {
		return nil, nil
	}

	params, trailing, matches := ParseWebhookPath(w.Path, r.URL.Path)
	if !matches {
		return nil, nil
	}

	// TODO(ppacher):
	// 	- verify IP whiltelist

	// ensure request body length does not exceed the expected limit
	if w.MaxContentLength > 0 && w.MaxContentLength > uint64(r.ContentLength) {
		return nil, nil
	}

	// ensure the expected content type matches
	if w.ExpectedContentType != "" && !checkMimeTypes(w.ExpectedContentType, r.Header.Get("Content-Type")) {
		return nil, nil
	}

	// check expected headers and there values
	if len(w.expectedHeadersParsed) > 0 {
		for key, m := range w.expectedHeadersParsed {
			// if there is no regex, just check for header presence
			if m == nil && r.Header.Get(key) == "" {
				return nil, nil
			}

			if m != nil {
				for _, val := range r.Header.Values(key) {
					if !m.Match([]byte(val)) {
						return nil, nil
					}
				}
			}
		}
	}

	evt := &eventsv1.WebhookEvent{
		WebhookPath:    w.Path,
		RequestPath:    r.URL.EscapedPath() + r.URL.Query().Encode(),
		HttpHeaders:    make(map[string]*eventsv1.HeaderValues),
		Content:        body,
		PathParameters: params,
		ReceivedAt:     timestamppb.Now(),
	}
	// TODO(ppacher): add trailing
	_ = trailing

	for key, values := range r.Header {
		evt.HttpHeaders[key] = &eventsv1.HeaderValues{
			Values: values,
		}
	}

	return evt, nil
}

func checkMimeTypes(expected string, got string) bool {
	expectedMime, _, err := mime.ParseMediaType(expected)
	if err != nil {
		return false
	}

	gotMime, _, err := mime.ParseMediaType(got)
	if err != nil {
		return false
	}

	if expectedMime == gotMime {
		return true
	}

	if strings.HasSuffix(expectedMime, "/*") {
		expectedParts := strings.Split(expectedMime, "/")
		gotParts := strings.Split(gotMime, "/")

		if len(expectedParts) != len(gotParts) {
			return false
		}

		for idx := 0; idx < len(expectedParts)-1; idx++ {
			if expectedParts[idx] != gotParts[idx] {
				return false
			}
		}
	}

	return true
}

func (w Webhook) ToProto() *eventsv1.Webhook {
	pb := &eventsv1.Webhook{
		WebhookPath:         w.Path,
		ExpectedContentType: w.ExpectedContentType,
		MaxContentLength:    int64(w.MaxContentLength),
		CreatedAt:           timestamppb.New(w.CreatedAt),
		LastUpdate:          timestamppb.New(w.LastUpdate),
		TimeToLive:          durationpb.New(w.TTL),
		IpWhitelist:         make([]string, len(w.IPWhiteList)),
		ExpectedHttpHeaders: make(map[string]string, len(w.ExpectedHeaders)),
	}

	for idx, net := range w.IPWhiteList {
		pb.IpWhitelist[idx] = net.String()
	}

	for key, val := range w.ExpectedHeaders {
		// FIXME(ppacher): wrong protobuf definition
		pb.ExpectedHttpHeaders[key] = val
	}

	return pb
}
