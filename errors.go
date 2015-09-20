package ghapi

import (
	"errors"
	"fmt"
)

type ErrHttpError struct {
	Status       string
	StatusCode   int
	Method       string
	RequestBody  string
	ResponseBody string
	Url          string
}

func (e ErrHttpError) Error() string {
	return fmt.Sprintf("%s\n%s %s\nRequest Body:\n%s\nResponse Body:\n%s", e.Status, e.Method, e.Url, e.RequestBody, e.ResponseBody)
}

var ErrSignatureNotFound = errors.New("\"X-Hub-Signature\" header not found")

var ErrSignatureMismatch = errors.New("signature mismatch")

var ErrGitHubEventNotFound = errors.New("\"X-Github-Event\" header not found")

var ErrSignatureMarkerNotFound = errors.New("\"sha1=\" marker not found")

var ErrHttpRequestBodyNil = errors.New("http.Request Body is nil")
