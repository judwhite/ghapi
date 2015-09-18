package ghapi

import (
	"fmt"
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
	ApiInfo
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

func (apiInfo *RepositoryInfo) getUrl(url string) string {
	url = strings.Replace(url, ":owner", apiInfo.Owner, 1)
	url = strings.Replace(url, ":repo", apiInfo.Repository, 1)
	url = apiInfo.BaseUrl + url
	return url
}

func (apiInfo *ApiInfo) httpRequest(method, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Content-Type", "application/json")
	if apiInfo.OAuth2Token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("token %s", apiInfo.OAuth2Token))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (apiInfo *ApiInfo) httpDelete(url string) (*http.Response, error) {
	return apiInfo.httpRequest("DELETE", url)
}

func (apiInfo *ApiInfo) httpGet(url string) (*http.Response, error) {
	return apiInfo.httpRequest("GET", url)
}
