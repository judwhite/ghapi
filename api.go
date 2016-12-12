package ghapi

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type APIInfo struct {
	BaseURL     string
	OAuth2Token string
}

type RepositoryInfo struct {
	APIInfo
	Owner      string
	Repository string
}

type GitHubAPI struct {
	APIInfo
	Issue        IssueAPI
	User         UserAPI
	Organization OrganizationAPI
	PullRequest  PullRequestsAPI
	Status       StatusAPI
	Branch       BranchesAPI
}

type IssueAPI struct {
	RepositoryInfo
}

type UserAPI struct {
	APIInfo
}

type OrganizationAPI struct {
	APIInfo
}

type PullRequestsAPI struct {
	RepositoryInfo
}

type StatusAPI struct {
	RepositoryInfo
}

type BranchesAPI struct {
	RepositoryInfo
}

func NewGitHubAPI(baseURL, owner, repository, authToken string) GitHubAPI {
	apiInfo := APIInfo{BaseURL: baseURL, OAuth2Token: authToken}

	gitHubAPI := GitHubAPI{APIInfo: apiInfo}

	repositoryInfo := RepositoryInfo{
		APIInfo:    apiInfo,
		Owner:      owner,
		Repository: repository,
	}

	gitHubAPI.Issue = IssueAPI{RepositoryInfo: repositoryInfo}
	gitHubAPI.User = UserAPI{APIInfo: apiInfo}
	gitHubAPI.Organization = OrganizationAPI{APIInfo: apiInfo}
	gitHubAPI.PullRequest = PullRequestsAPI{RepositoryInfo: repositoryInfo}
	gitHubAPI.Status = StatusAPI{RepositoryInfo: repositoryInfo}
	gitHubAPI.Branch = BranchesAPI{RepositoryInfo: repositoryInfo}

	return gitHubAPI
}

func (apiInfo *APIInfo) addBaseURL(url string) string {
	return apiInfo.BaseURL + url
}

func (apiInfo *RepositoryInfo) getURL(url string) string {
	url = strings.Replace(url, ":owner", apiInfo.Owner, 1)
	url = strings.Replace(url, ":repo", apiInfo.Repository, 1)
	url = apiInfo.BaseURL + url
	return url
}

func (apiInfo *APIInfo) getHTTPRequest(method, url string, body *string) (*http.Request, error) {
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

func (apiInfo *APIInfo) doHTTPRequest(method, url string, body *string) (*http.Response, error) {
	req, err := apiInfo.getHTTPRequest(method, url, body)
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
		err = &ErrHTTPError{
			Status:       resp.Status,
			StatusCode:   resp.StatusCode,
			Method:       method,
			URL:          url,
			RequestBody:  requestBody,
			ResponseBody: responseBody,
		}
		return nil, err
	}

	return resp, nil
}

func (apiInfo *APIInfo) httpDelete(url string) (*http.Response, error) {
	return apiInfo.doHTTPRequest("DELETE", url, nil)
}

func (apiInfo *APIInfo) httpGet(url string) (*http.Response, error) {
	return apiInfo.doHTTPRequest("GET", url, nil)
}

func (apiInfo *APIInfo) httpPatch(url, body string) (*http.Response, error) {
	return apiInfo.doHTTPRequest("PATCH", url, &body)
}

func (apiInfo *APIInfo) httpPost(url, body string) (*http.Response, error) {
	return apiInfo.doHTTPRequest("POST", url, &body)
}
