package ghapi

import (
	"errors"
	"fmt"
	"strings"
)

// ErrHTTPError is returned when a non-200 status code is returned from a GitHub API call.
type ErrHTTPError struct {
	Message      string
	Status       string
	StatusCode   int
	Method       string
	RequestBody  string
	ResponseBody string
	URL          string
}

func (e ErrHTTPError) Error() string {
	message := fmt.Sprintf("%s %s", e.Status, e.Message)
	message = strings.TrimSpace(message)
	return fmt.Sprintf("%s\n%s %s\nRequest Body:\n%s\nResponse Body:\n%s", message, e.Method, e.URL, e.RequestBody, e.ResponseBody)
}

// ErrSignatureNotFound is returned when the "X-Hub-Signature" header is not found in a GitHub event.
var ErrSignatureNotFound = errors.New("\"X-Hub-Signature\" header not found")

// ErrSignatureMismatch is returned when the "X-Hub-Signature" in a GitHub event does not match the expected signature.
var ErrSignatureMismatch = errors.New("signature mismatch")

// ErrGitHubEventNotFound is returned when the "X-Github-Event" header is not found in a GitHub event.
var ErrGitHubEventNotFound = errors.New("\"X-Github-Event\" header not found")

// ErrSignatureMarkerNotFound is returned when the "sha1=" marker is not found in the "X-Hub-Signature" header in
// a GitHub event.
var ErrSignatureMarkerNotFound = errors.New("\"sha1=\" marker not found")

// ErrHTTPRequestBodyNil is returned when the request body is nil from a GitHub event.
var ErrHTTPRequestBodyNil = errors.New("http.Request Body is nil")
