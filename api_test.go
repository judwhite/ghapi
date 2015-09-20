package ghapi

import (
	"fmt"
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
				"     - kind: '%v%'\n"+
				"     - actual:'%v'\n"+
				"     - reflect.ValueOf(actual):'%v'",
				msg, k, actual, value)
		}
	}
}

func expectNil(t *testing.T, actual interface{}, msg string) {
	if actual != nil && !reflect.ValueOf(actual).IsNil() {
		t.Fatalf("%s - '%v' expected to not be <nil>", msg, actual)
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
	expectedBaseUrl    string = "http://baseurl.org"
	expectedOwner             = "test_owner"
	expectedRepository        = "test_repository"
	expectedAuthToken         = "test_authtoken"
)

func makeGitHubApi() GitHubApi {
	return NewGitHubApi(expectedBaseUrl, expectedOwner, expectedRepository, expectedAuthToken)
}

func makeGitHubApiTestServer(handler func(w http.ResponseWriter, r *http.Request)) (*httptest.Server, GitHubApi, chan struct{}) {
	signal := make(chan struct{}, 1)
	httptestServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
		signal <- struct{}{}
	}))
	api := NewGitHubApi(httptestServer.URL, expectedOwner, expectedRepository, expectedAuthToken)
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

func TestNewGitHubApi(t *testing.T) {
	api := makeGitHubApi()

	expect(t, expectedBaseUrl, api.BaseUrl, "api.BaseUrl")
	expect(t, expectedAuthToken, api.OAuth2Token, "api.OAuth2Token")

	expect(t, expectedBaseUrl, api.Issue.BaseUrl, "api.Issue.BaseUrl")
	expect(t, expectedAuthToken, api.Issue.OAuth2Token, "api.Issue.OAuth2Token")
	expect(t, expectedOwner, api.Issue.Owner, "api.Issue.Owner")
	expect(t, expectedRepository, api.Issue.Repository, "api.Issue.Repository")

	expect(t, expectedBaseUrl, api.User.BaseUrl, "api.User.BaseUrl")
	expect(t, expectedAuthToken, api.User.OAuth2Token, "api.User.OAuth2Token")

	expect(t, expectedBaseUrl, api.Organization.BaseUrl, "api.Organization.BaseUrl")
	expect(t, expectedAuthToken, api.Organization.OAuth2Token, "api.Organization.OAuth2Token")
}

func TestApiInfo_addBaseUrl(t *testing.T) {
	api := makeGitHubApi()

	expected := "http://baseurl.org/suffix"

	actual := api.addBaseUrl("/suffix")

	expect(t, expected, actual, "addBaseUrl")
}

func TestApiInfo_doHttpRequest_ReturnsErrOnParseError(t *testing.T) {
	api := makeGitHubApi()
	resp, err := api.doHttpRequest("GET", ":/noscheme", nil)
	expectNil(t, resp, "resp")
	expectNotNil(t, err, "err")
	expect(t, "parse :/noscheme: missing protocol scheme", err.Error(), "err.Error()")
}

func TestApiInfo_doHttpRequest_ReturnsErrOnDoError(t *testing.T) {
	api := makeGitHubApi()
	resp, err := api.doHttpRequest("GET", "http://0.0.0.0:0/wat", nil)
	expectNil(t, resp, "resp")
	expectNotNil(t, err, "err")
	expect(t, "Get http://0.0.0.0:0/wat: dial tcp 0.0.0.0:0: connectex: The requested address is not valid in its context.", err.Error(), "err.Error()")
}

func TestApiInfo_doHttpRequest_HasHeadersSet(t *testing.T) {
	ts, api, signal := makeGitHubApiTestServer(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		if r.URL != nil && r.URL.Path == "/test_headers" {
			expect(t, "application/vnd.github.v3+json", r.Header.Get("Accept"), "r.Header[\"Accept\"]")
			expect(t, "application/json", r.Header.Get("Content-Type"), "r.Header[\"Content-Type\"]")
			expect(t, "token "+expectedAuthToken, r.Header.Get("Authorization"), "r.Header[\"Authorization\"]")
		} else {
			t.Fatalf("unexpected URL %v", r.URL)
		}
	})
	defer ts.Close()

	resp, err := api.doHttpRequest("GET", ts.URL+"/test_headers", nil)
	waitSignal(t, signal)

	expectNotNil(t, resp, "resp")
	expectNil(t, err, "err")
	defer resp.Body.Close()
}

func TestApiInfo_doHttpRequest_AuthorizationHeadersNotSetWhenAuthTokenIsEmpty(t *testing.T) {
	ts, api, signal := makeGitHubApiTestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	resp, err := api.doHttpRequest("GET", ts.URL+"/test_headers", nil)
	waitSignal(t, signal)

	expectNotNil(t, resp, "resp")
	expectNil(t, err, "err")
	defer resp.Body.Close()
}

func TestRepositoryInfo_getUrl(t *testing.T) {
	issueApi := makeGitHubApi().Issue

	expected := "http://baseurl.org/test_owner/test_repository/suffix"

	actual := issueApi.getUrl("/:owner/:repo/suffix")

	expect(t, expected, actual, "getUrl")
}

func TestRepositoryInfo_httpPatch(t *testing.T) {
	const expectedBody string = "expected body"

	ts, api, signal := makeGitHubApiTestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	ts, api, signal := makeGitHubApiTestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	expectedUrl := ts.URL + "/test_patch"
	resp, err := api.httpPatch(expectedUrl, expectedRequestBody)
	waitSignal(t, signal)

	expectNil(t, resp, "resp")
	expectNotNil(t, err, "err")

	e, ok := err.(*ErrHttpError)
	if !ok {
		t.Fatal("err is not of type *ErrHttpError")
	}

	expect(t, "404 Not Found", e.Status, "e.Status")
	expect(t, 404, e.StatusCode, "e.StatusCode")
	expect(t, "PATCH", e.Method, "e.Method")
	expect(t, expectedRequestBody, e.RequestBody, "e.RequestBody")
	expect(t, expectedResponseBody, e.ResponseBody, "e.ResponseBody")
	expect(t, expectedUrl, e.Url, "e.Url")
}
