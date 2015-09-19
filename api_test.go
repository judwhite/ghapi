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

func TestRepositoryInfo_getUrl(t *testing.T) {
	issueApi := makeGitHubApi().Issue

	expected := "http://baseurl.org/test_owner/test_repository/suffix"

	actual := issueApi.getUrl("/:owner/:repo/suffix")

	expect(t, expected, actual, "getUrl")
}
