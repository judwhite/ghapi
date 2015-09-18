package ghapi

import "fmt"

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
