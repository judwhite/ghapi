package ghapi

import (
	"errors"
	"fmt"
)

type ErrHTTPError struct {
	Status       string
	StatusCode   int
	Method       string
	RequestBody  string
	ResponseBody string
	URL          string
}

func (e ErrHTTPError) Error() string {
	return fmt.Sprintf("%s\n%s %s\nRequest Body:\n%s\nResponse Body:\n%s", e.Status, e.Method, e.URL, e.RequestBody, e.ResponseBody)
}

var ErrSignatureNotFound = errors.New("\"X-Hub-Signature\" header not found")

var ErrSignatureMismatch = errors.New("signature mismatch")

var ErrGitHubEventNotFound = errors.New("\"X-Github-Event\" header not found")

var ErrSignatureMarkerNotFound = errors.New("\"sha1=\" marker not found")

var ErrHTTPRequestBodyNil = errors.New("http.Request Body is nil")
