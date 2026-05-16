package webhook

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
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

	CreatedAt  time.Time
	LastUpdate time.Time
}

func (w Webhook) MatchRequest(ctx context.Context, r *http.Request) (event *eventsv1.WebhookEvent, err error) {
	params, trailing, matches := ParseWebhookPath(w.Path, r.URL.Path)
	if !matches {
		return nil, nil
	}

	content, err := io.ReadAll(r.Body)

	evt := &eventsv1.WebhookEvent{
		WebhookPath:    w.Path,
		RequestPath:    r.URL.EscapedPath() + r.URL.Query().Encode(),
		HttpHeaders:    make(map[string]*eventsv1.HeaderValues),
		Content:        content,
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

// ParseWebhookPath checks if the passed URL matches the webhooks path definition.
// If matches, path parameters are extracted and returned.
func ParseWebhookPath(pattern string, urlPath string) (params map[string]string, trailing string, matches bool) {
	patternParts := strings.Split(pattern, "/")
	urlParts := strings.Split(urlPath, "/")

	// quick check
	if len(urlParts) < len(patternParts) {
		log.Println("len mismatch")
		return nil, "", false
	}

	if len(urlParts) > len(patternParts) && patternParts[len(patternParts)-1] != "{#}" {
		return nil, "", false
	}

	params = make(map[string]string)
	for idx, p := range patternParts {
		if len(p) > 0 && p[0] == '{' && p[len(p)-1] == '}' {
			// parse path parameter
			paramName := p[1 : len(p)-1]
			if paramName == "" {
				// this is actual an error and should have been caught before
				// so just return "no-match"
				return nil, "", false
			}

			if idx == len(patternParts)-1 && paramName == "#" {
				return params, strings.Join(urlParts[idx:], "/"), true
			}

			params[paramName] = urlParts[idx]
		} else if p != urlParts[idx] {
			log.Printf("%q does not match %q", p, urlParts[idx])
			// the path parts don't match
			return nil, "", false
		}
	}

	// if we get here, there's no trailing part
	return params, "", true
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
