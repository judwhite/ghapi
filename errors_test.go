package ghapi

import "testing"

func TestErrHttpError_Error(t *testing.T) {
	e := ErrHTTPError{
		Status:       "404 Not Found",
		StatusCode:   404,
		Method:       "POST",
		URL:          "http://example.org",
		RequestBody:  "{ id: \"1\" }",
		ResponseBody: "{ message: \"not found\" }",
	}

	const expected string = "404 Not Found\nPOST http://example.org\nRequest Body:\n{ id: \"1\" }\nResponse Body:\n{ message: \"not found\" }"

	expect(t, expected, e.Error(), "e.Error()")
}
