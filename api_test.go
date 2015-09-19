package ghapi

import (
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

		t.Fatalf("%s '%v' != '%v'", msg, expected, actual)
	}
}

func expectNotNil(t *testing.T, actual interface{}, msg string) {
	if actual == nil || reflect.ValueOf(actual).IsNil() {
		t.Fatalf("%s - expected to not be '%v'", msg, actual)
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

func makeGitHubApiTestServer(handler http.Handler) (*httptest.Server, GitHubApi) {
	httptestServer := httptest.NewServer(handler)
	api := NewGitHubApi(httptestServer.URL, expectedOwner, expectedRepository, expectedAuthToken)
	return httptestServer, api
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
	req, err := api.doHttpRequest("GET", ":/noscheme", nil)
	expectNil(t, req, "req")
	expectNotNil(t, err, "err")
	expect(t, "parse :/noscheme: missing protocol scheme", err.Error(), "err.Error()")
}

func TestApiInfo_doHttpRequest_ReturnsErrOnDoError(t *testing.T) {
	api := makeGitHubApi()
	req, err := api.doHttpRequest("GET", "http://0.0.0.0:0/wat", nil)
	expectNil(t, req, "req")
	expectNotNil(t, err, "err")
	expect(t, "Get http://0.0.0.0:0/wat: dial tcp 0.0.0.0:0: connectex: The requested address is not valid in its context.", err.Error(), "err.Error()")
}

func TestApiInfo_doHttpRequest_HasHeadersSet(t *testing.T) {
	signal := make(chan struct{}, 1)
	ts, api := makeGitHubApiTestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signal <- struct{}{}
		if r.URL != nil && r.URL.Path == "/test_headers" {
			expect(t, "application/vnd.github.v3+json", r.Header.Get("Accept"), "r.Header[\"Accept\"]")
			expect(t, "application/json", r.Header.Get("Content-Type"), "r.Header[\"Content-Type\"]")
			expect(t, "token "+expectedAuthToken, r.Header.Get("Authorization"), "r.Header[\"Authorization\"]")
		} else {
			t.Fatalf("unexpected URL %v", r.URL)
		}
	}))
	defer ts.Close()

	resp, err := api.doHttpRequest("GET", ts.URL+"/test_headers", nil)
	<-signal

	expectNotNil(t, resp, "resp")
	expectNil(t, err, "err")
}

func TestApiInfo_doHttpRequest_AuthorizationHeadersNotSetWhenAuthTokenIsEmpty(t *testing.T) {
	signal := make(chan struct{}, 1)
	ts, api := makeGitHubApiTestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signal <- struct{}{}
		if r.URL != nil && r.URL.Path == "/test_headers" {
			expect(t, "application/vnd.github.v3+json", r.Header.Get("Accept"), "r.Header[\"Accept\"]")
			expect(t, "application/json", r.Header.Get("Content-Type"), "r.Header[\"Content-Type\"]")
			expect(t, "", r.Header.Get("Authorization"), "r.Header[\"Authorization\"]")
		} else {
			t.Fatalf("unexpected URL %v", r.URL)
		}
	}))
	defer ts.Close()

	api.OAuth2Token = ""
	resp, err := api.doHttpRequest("GET", ts.URL+"/test_headers", nil)
	<-signal

	expectNotNil(t, resp, "resp")
	expectNil(t, err, "err")
}

func TestRepositoryInfo_getUrl(t *testing.T) {
	issueApi := makeGitHubApi().Issue

	expected := "http://baseurl.org/test_owner/test_repository/suffix"

	actual := issueApi.getUrl("/:owner/:repo/suffix")

	expect(t, expected, actual, "getUrl")
}
