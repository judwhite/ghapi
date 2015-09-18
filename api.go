package ghapi

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type ApiInfo struct {
	BaseUrl     string
	OAuth2Token string
}

type RepositoryInfo struct {
	ApiInfo
	Owner      string
	Repository string
}

type GitHubApi struct {
	ApiInfo
	Issue        IssueApi
	User         UserApi
	Organization OrganizationApi
}

type IssueApi struct {
	RepositoryInfo
}

type UserApi struct {
	ApiInfo
}

type OrganizationApi struct {
	ApiInfo
}

type PullRequestsApi struct {
	ApiInfo
	RepositoryInfo
}

func NewGitHubApi(baseUrl, owner, repository, authToken string) GitHubApi {
	apiInfo := ApiInfo{BaseUrl: baseUrl, OAuth2Token: authToken}

	gitHubApi := GitHubApi{ApiInfo: apiInfo}

	repositoryInfo := RepositoryInfo{
		ApiInfo:    apiInfo,
		Owner:      owner,
		Repository: repository,
	}

	gitHubApi.Issue = IssueApi{RepositoryInfo: repositoryInfo}
	gitHubApi.User = UserApi{ApiInfo: apiInfo}
	gitHubApi.Organization = OrganizationApi{ApiInfo: apiInfo}

	return gitHubApi
}

func (apiInfo *ApiInfo) addBaseUrl(url string) string {
	return apiInfo.BaseUrl + url
}

func (apiInfo *RepositoryInfo) getUrl(url string) string {
	url = strings.Replace(url, ":owner", apiInfo.Owner, 1)
	url = strings.Replace(url, ":repo", apiInfo.Repository, 1)
	url = apiInfo.BaseUrl + url
	return url
}

func (apiInfo *ApiInfo) getHttpRequest(method, url string, body *string) (*http.Request, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyReader = strings.NewReader(*body)
	}
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Content-Type", "application/json")
	if apiInfo.OAuth2Token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("token %s", apiInfo.OAuth2Token))
	}

	return req, nil
}

func (apiInfo *ApiInfo) doHttpRequest(method, url string, body *string) (*http.Response, error) {
	req, err := apiInfo.getHttpRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		defer resp.Body.Close()
		var requestBody, responseBody string
		b, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			responseBody = string(b)
		}
		if body != nil {
			requestBody = *body
		}
		return nil, ErrHttpError{Status: resp.Status, StatusCode: resp.StatusCode, Url: url, RequestBody: requestBody, ResponseBody: responseBody}
	}

	return resp, nil
}

func (apiInfo *ApiInfo) httpDelete(url string) (*http.Response, error) {
	return apiInfo.doHttpRequest("DELETE", url, nil)
}

func (apiInfo *ApiInfo) httpGet(url string) (*http.Response, error) {
	return apiInfo.doHttpRequest("GET", url, nil)
}

func (apiInfo *ApiInfo) httpPatch(url, body string) (*http.Response, error) {
	return apiInfo.doHttpRequest("PATCH", url, &body)
}
