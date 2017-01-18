package ghapi

import (
	"encoding/json"
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
	Repository   RepositoryAPI
	Contents     ContentsAPI
}

type IssueAPI struct {
	RepositoryInfo
}

type UserAPI struct {
	APIInfo
}

type OrganizationAPI struct {
	APIInfo
	Organization string
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

type RepositoryAPI struct {
	RepositoryInfo
}

type ContentsAPI struct {
	RepositoryInfo
}

type AuthenticatedUser struct {
	Login             string     `json:"login"`
	ID                int        `json:"id"`
	AvatarURL         string     `json:"avatar_url"`
	GravatarID        string     `json:"gravatar_id"`
	URL               string     `json:"url"`
	HTMLURL           string     `json:"html_url"`
	FollowersURL      string     `json:"followers_url"`
	FollowingURL      string     `json:"following_url"`
	GistsURL          string     `json:"gists_url"`
	StarredURL        string     `json:"starred_url"`
	SubscriptionsURL  string     `json:"subscriptions_url"`
	OrganizationsURL  string     `json:"organizations_url"`
	ReposURL          string     `json:"repos_url"`
	EventsURL         string     `json:"events_url"`
	ReceivedEventsURL string     `json:"received_events_url"`
	Type              string     `json:"type"`
	SiteAdmin         bool       `json:"site_admin"`
	Name              string     `json:"name"`
	Company           string     `json:"company"`
	Blog              string     `json:"blog"`
	Location          string     `json:"location"`
	Email             string     `json:"email"`
	Hireable          bool       `json:"hireable"`
	Bio               string     `json:"bio"`
	PublicRepos       int        `json:"public_repos"`
	PublicGists       int        `json:"public_gists"`
	Followers         int        `json:"followers"`
	Following         int        `json:"following"`
	CreatedAt         CustomTime `json:"created_at"`
	UpdatedAt         CustomTime `json:"updated_at"`
	TotalPrivateRepos int        `json:"total_private_repos"`
	OwnedPrivateRepos int        `json:"owned_private_repos"`
	PrivateGists      int        `json:"private_gists"`
	DiskUsage         int        `json:"disk_usage"`
	Collaborators     int        `json:"collaborators"`
	Plan              struct {
		Name          string `json:"name"`
		Space         int    `json:"space"`
		PrivateRepos  int    `json:"private_repos"`
		Collaborators int    `json:"collaborators"`
	} `json:"plan"`
}

// NewGitHubAPI returns a new GitHubAPI using the specified repository and authentication information.
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
	gitHubAPI.Organization = OrganizationAPI{APIInfo: apiInfo, Organization: owner}
	gitHubAPI.PullRequest = PullRequestsAPI{RepositoryInfo: repositoryInfo}
	gitHubAPI.Status = StatusAPI{RepositoryInfo: repositoryInfo}
	gitHubAPI.Branch = BranchesAPI{RepositoryInfo: repositoryInfo}
	gitHubAPI.Repository = RepositoryAPI{RepositoryInfo: repositoryInfo}
	gitHubAPI.Contents = ContentsAPI{RepositoryInfo: repositoryInfo}

	return gitHubAPI
}

// GetUser returns the current authenticated user.
func GetUser(baseURL, authToken string) (*AuthenticatedUser, error) {
	apiInfo := APIInfo{BaseURL: baseURL, OAuth2Token: authToken}
	url := apiInfo.addBaseURL("/user")
	resp, err := apiInfo.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user AuthenticatedUser

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// Is404 returns true if the error is an HTTP 404.
func Is404(err error) bool {
	return IsHTTPError(err, 404)
}

// IsHTTPError returns true if the error is an HTTP error with the specified status code.
func IsHTTPError(err error, statusCode int) bool {
	switch val := err.(type) {
	case *ErrHTTPError:
		if val.StatusCode == statusCode {
			return true
		}
	}
	return false
}

func (apiInfo *APIInfo) addBaseURL(url string) string {
	return apiInfo.BaseURL + url
}

func (apiInfo *RepositoryInfo) getURL(url string) string {
	url = strings.Replace(url, ":owner", apiInfo.Owner, 1)
	url = strings.Replace(url, ":repo", apiInfo.Repository, 1)
	return apiInfo.addBaseURL(url)
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
