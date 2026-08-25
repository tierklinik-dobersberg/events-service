package webhook

import (
	"fmt"
	"strings"
)

// ValidateWebhookPath validates a webhook path pattern for validity.
func ValidateWebhookPath(path string) error {
	if len(path) == 0 {
		return fmt.Errorf("missing webhook path")
	}

	parts := strings.Split(path, "/")
	names := make(map[string]struct{})
	for idx, p := range parts {
		if len(p) == 0 {
			return fmt.Errorf("empty path segment %d", idx)
		}

		if p[0] == '{' && p[len(p)-1] == '}' {
			paramName := p[1 : len(p)-1]
			if paramName == "" {
				return fmt.Errorf("missing parameter name in segment %d", idx)
			}

			if paramName == "#" && idx != len(parts)-1 {
				return fmt.Errorf("trailing parameter match must be the last segment, found at %d", idx)
			}

			if _, ok := names[paramName]; ok {
				return fmt.Errorf("duplicated parameter name %s in segment %d", paramName, idx)
			}
			names[paramName] = struct{}{}
		}
	}

	return nil
}

// ParseWebhookPath checks if the passed URL matches the webhooks path definition.
// If matches, path parameters are extracted and returned.
func ParseWebhookPath(pattern string, urlPath string) (params map[string]string, trailing string, matches bool) {
	pattern = strings.TrimPrefix(strings.TrimSuffix(pattern, "/"), "/")
	urlPath = strings.TrimPrefix(strings.TrimSuffix(urlPath, "/"), "/")

	patternParts := strings.Split(pattern, "/")
	urlParts := strings.Split(urlPath, "/")

	// quick check
	if len(urlParts) < len(patternParts) {
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
			// the path parts don't match
			return nil, "", false
		}
	}

	// if we get here, there's no trailing part
	return params, "", true
}
