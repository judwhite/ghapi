package ghapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func expect(t *testing.T, expected, actual interface{}, msg string) {
	if expected != actual {
		// handle comparing time.Time to *time.Time
		if expectedDate, ok := expected.(time.Time); ok {
			if actualDate, ok := actual.(*time.Time); ok {
				if actualDate != nil && expectedDate == *actualDate {
					return
				}
			}
		}

		t.Fatalf("%s '%v' (%T) != '%v' (%T)", msg, expected, expected, actual, actual)
	}
}

func expectNotNil(t *testing.T, actual interface{}, msg string) {
	if actual == nil {
		t.Fatalf("%s - expected to not be nil\n     - actual: '%v'", msg, actual)
	} else {
		value := reflect.ValueOf(actual)
		k := value.Kind()
		isNil := false
		switch k {
		case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
			isNil = value.IsNil()
		}

		if isNil {
			t.Fatalf("%s - expected to not be nil\n"+
				"     - kind: '%v'\n"+
				"     - actual:'%v'\n"+
				"     - reflect.ValueOf(actual):'%v'",
				msg, k, actual, value)
		}
	}
}

func expectNil(t *testing.T, actual interface{}, msg string) {
	if actual != nil && !reflect.ValueOf(actual).IsNil() {
		t.Fatalf("%s - '%v' expected to be <nil>", msg, actual)
	}
}

func expectErrHTTPError500(t *testing.T, err error) {
	if err != nil {
		if e, ok := err.(*ErrHTTPError); !ok {
			t.Fatalf("err is not of type *ErrHttpError, is %T", err)
		} else {
			expect(t, 500, e.StatusCode, "e.StatusCode")
		}
	} else {
		t.Fatal("expected error")
	}
}

func expectJSONSyntaxError(t *testing.T, err error, expectedMessage string) {
	if err != nil {
		if e, ok := err.(*json.SyntaxError); !ok {
			t.Fatalf("err is not of type *json.SyntaxError, is %T", err)
		} else {
			expect(t, expectedMessage, e.Error(), "e.Error()")
		}
	} else {
		t.Fatal("expected error")
	}
}

func date(value string) time.Time {
	date, err := time.Parse(time.RFC3339, value)
	if err != nil {
		panic(err)
	}
	return date
}

const (
	expectedBaseURL    string = "http://baseurl.org"
	expectedOwner             = "test_owner"
	expectedRepository        = "test_repository"
	expectedAuthToken         = "test_authtoken"
)

func makeGitHubAPI() GitHubAPI {
	return NewGitHubAPI(expectedBaseURL, expectedOwner, expectedRepository, expectedAuthToken)
}

func makeGitHubAPITestServer(handler func(w http.ResponseWriter, r *http.Request)) (*httptest.Server, GitHubAPI, chan struct{}) {
	signal := make(chan struct{}, 1)
	httptestServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
		signal <- struct{}{}
	}))
	api := NewGitHubAPI(httptestServer.URL, expectedOwner, expectedRepository, expectedAuthToken)
	return httptestServer, api, signal
}

func waitSignal(t *testing.T, signal <-chan struct{}) {
	ticker := time.NewTicker(5 * time.Second)
	select {
	case <-signal:
		return
	case <-ticker.C:
		t.Fatal("timeout")
	}
}

func TestNewGitHubAPI(t *testing.T) {
	api := makeGitHubAPI()

	expect(t, expectedBaseURL, api.BaseURL, "api.BaseUrl")
	expect(t, expectedAuthToken, api.OAuth2Token, "api.OAuth2Token")

	expect(t, expectedBaseURL, api.Issue.BaseURL, "api.Issue.BaseUrl")
	expect(t, expectedAuthToken, api.Issue.OAuth2Token, "api.Issue.OAuth2Token")
	expect(t, expectedOwner, api.Issue.Owner, "api.Issue.Owner")
	expect(t, expectedRepository, api.Issue.Repository, "api.Issue.Repository")

	expect(t, expectedBaseURL, api.User.BaseURL, "api.User.BaseUrl")
	expect(t, expectedAuthToken, api.User.OAuth2Token, "api.User.OAuth2Token")

	expect(t, expectedBaseURL, api.Organization.BaseURL, "api.Organization.BaseUrl")
	expect(t, expectedAuthToken, api.Organization.OAuth2Token, "api.Organization.OAuth2Token")
}

func TestApiInfo_addBaseUrl(t *testing.T) {
	api := makeGitHubAPI()

	expected := "http://baseurl.org/suffix"

	actual := api.addBaseURL("/suffix")

	expect(t, expected, actual, "addBaseUrl")
}

func TestApiInfo_doHttpRequest_ReturnsErrOnParseError(t *testing.T) {
	api := makeGitHubAPI()
	resp, err := api.doHTTPRequest("GET", ":/noscheme", nil, "")
	expectNil(t, resp, "resp")
	expectNotNil(t, err, "err")
	expect(t, "parse :/noscheme: missing protocol scheme", err.Error(), "err.Error()")
}

func TestAPIInfo_doHTTPRequest_ReturnsErrOnDoError(t *testing.T) {
	api := makeGitHubAPI()
	resp, err := api.doHTTPRequest("GET", "http://0.0.0.0:0/wat", nil, "")
	expectNil(t, resp, "resp")
	expectNotNil(t, err, "err")
	expect(t, "Get http://0.0.0.0:0/wat: dial tcp 0.0.0.0:0: connectex: The requested address is not valid in its context.", err.Error(), "err.Error()")
}

func TestAPIInfo_doHTTPRequest_HasHeadersSet(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/test_headers" {
			expect(t, "application/vnd.github.v3+json", r.Header.Get("Accept"), "r.Header[\"Accept\"]")
			expect(t, "application/json", r.Header.Get("Content-Type"), "r.Header[\"Content-Type\"]")
			expect(t, "token "+expectedAuthToken, r.Header.Get("Authorization"), "r.Header[\"Authorization\"]")
		} else {
			t.Fatalf("unexpected URL %v", r.URL)
		}
	})
	defer ts.Close()

	resp, err := api.doHTTPRequest("GET", ts.URL+"/test_headers", nil, "")
	waitSignal(t, signal)

	expectNotNil(t, resp, "resp")
	expectNil(t, err, "err")
	defer resp.Body.Close()
}

func TestAPIInfo_doHTTPRequest_AuthorizationHeadersNotSetWhenAuthTokenIsEmpty(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/test_headers" {
			expect(t, "application/vnd.github.v3+json", r.Header.Get("Accept"), "r.Header.Get(\"Accept\")")
			expect(t, "application/json", r.Header.Get("Content-Type"), "r.Header.Get(\"Content-Type\")")
			expect(t, 0, len(r.Header["Authorization"]), "len(r.Header[\"Authorization\"])")
		} else {
			t.Fatalf("unexpected URL %v", r.URL)
		}
	}))
	defer ts.Close()

	api.OAuth2Token = ""
	resp, err := api.doHTTPRequest("GET", ts.URL+"/test_headers", nil, "")
	waitSignal(t, signal)

	expectNotNil(t, resp, "resp")
	expectNil(t, err, "err")
	defer resp.Body.Close()
}

func TestRepositoryInfo_getURL(t *testing.T) {
	issueAPI := makeGitHubAPI().Issue

	expected := "http://baseurl.org/test_owner/test_repository/suffix"

	actual := issueAPI.getURL("/:owner/:repo/suffix")

	expect(t, expected, actual, "getUrl")
}

func TestRepositoryInfo_httpPatch(t *testing.T) {
	const expectedBody string = "expected body"

	ts, api, signal := makeGitHubAPITestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/test_patch" {
			expectNotNil(t, r.Body, "r.Body")
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
				return
			}
			expect(t, expectedBody, string(b), "b")
		} else {
			t.Fatalf("unexpected URL %v", r.URL)
		}
	}))
	defer ts.Close()

	resp, err := api.httpPatch(ts.URL+"/test_patch", expectedBody)
	waitSignal(t, signal)

	expectNotNil(t, resp, "resp")
	expectNil(t, err, "err")
	defer resp.Body.Close()
}

func TestRepositoryInfo_httpPatch_ErrorPopulated(t *testing.T) {
	const expectedRequestBody string = "expected request body"
	const expectedResponseBody string = "expected response body"

	ts, api, signal := makeGitHubAPITestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/test_patch" {
			w.WriteHeader(404)
			_, err := w.Write([]byte(expectedResponseBody))
			if err != nil {
				t.Fatal(err)
			}
		} else {
			t.Fatalf("unexpected URL %v", r.URL)
		}
	}))
	defer ts.Close()

	expectedURL := ts.URL + "/test_patch"
	resp, err := api.httpPatch(expectedURL, expectedRequestBody)
	waitSignal(t, signal)

	expectNil(t, resp, "resp")
	expectNotNil(t, err, "err")

	e, ok := err.(*ErrHTTPError)
	if !ok {
		t.Fatal("err is not of type *ErrHttpError")
	}

	expect(t, "404 Not Found", e.Status, "e.Status")
	expect(t, 404, e.StatusCode, "e.StatusCode")
	expect(t, "PATCH", e.Method, "e.Method")
	expect(t, expectedRequestBody, e.RequestBody, "e.RequestBody")
	expect(t, expectedResponseBody, e.ResponseBody, "e.ResponseBody")
	expect(t, expectedURL, e.URL, "e.Url")
}
