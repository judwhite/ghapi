package ghapi

import (
	"testing"
)

func assert_string(t *testing.T, expected, actual, msg string) {
	if expected != actual {
		t.Fatalf("%s '%s' != '%s'", msg, expected, actual)
	}
}

const (
	expectedBaseUrl    string = "http://baseurl.org"
	expectedOwner             = "owner"
	expectedRepository        = "repository"
	expectedAuthToken         = "authtoken"
)

func makeGitHubApi() GitHubApi {
	return NewGitHubApi(expectedBaseUrl, expectedOwner, expectedRepository, expectedAuthToken)
}

func TestNewGitHubApi(t *testing.T) {
	api := makeGitHubApi()

	assert_string(t, expectedBaseUrl, api.BaseUrl, "api.BaseUrl")
	assert_string(t, expectedAuthToken, api.OAuth2Token, "api.OAuth2Token")

	assert_string(t, expectedBaseUrl, api.Issue.BaseUrl, "api.Issue.BaseUrl")
	assert_string(t, expectedAuthToken, api.Issue.OAuth2Token, "api.Issue.OAuth2Token")
	assert_string(t, expectedOwner, api.Issue.Owner, "api.Issue.Owner")
	assert_string(t, expectedRepository, api.Issue.Repository, "api.Issue.Repository")

	assert_string(t, expectedBaseUrl, api.User.BaseUrl, "api.User.BaseUrl")
	assert_string(t, expectedAuthToken, api.User.OAuth2Token, "api.User.OAuth2Token")

	assert_string(t, expectedBaseUrl, api.Organization.BaseUrl, "api.Organization.BaseUrl")
	assert_string(t, expectedAuthToken, api.Organization.OAuth2Token, "api.Organization.OAuth2Token")
}

func TestApiInfo_addBaseUrl(t *testing.T) {
	api := makeGitHubApi()

	expected := "http://baseurl.org/suffix"

	actual := api.addBaseUrl("/suffix")

	assert_string(t, expected, actual, "addBaseUrl")
}

func TestRepositoryInfo_getUrl(t *testing.T) {
	issueApi := makeGitHubApi().Issue

	expected := "http://baseurl.org/owner/repository/suffix"

	actual := issueApi.getUrl("/:owner/:repo/suffix")

	assert_string(t, expected, actual, "getUrl")
}
