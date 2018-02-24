package ghapi

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// APIInfo contains the API URL and Authtoken for making API requests. APIInfo is embedded in all other API structs.
type APIInfo struct {
	BaseURL     string
	OAuth2Token string
}

// RepositoryInfo contains APIInfo, the repository owner, and the repository name.
type RepositoryInfo struct {
	APIInfo
	Owner      string
	Repository string
}

// GitHubAPI is returned by the NewGitHubAPI function and contains fields for making API calls.
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
	Refs         RefsAPI
}

// IssueAPI is used to get information about a repository's issues. Note Pull Requests are treated as issues in some
// cases, for example labeling a Pull Request is done through IssueAPI.
type IssueAPI struct {
	RepositoryInfo
}

// UserAPI is used to get information about a user.
type UserAPI struct {
	APIInfo
}

// OrganizationAPI is used to get information about an organization.
type OrganizationAPI struct {
	APIInfo
	Organization string
}

// PullRequestsAPI is used to get information about a repository's pull requests.
type PullRequestsAPI struct {
	RepositoryInfo
}

// StatusAPI is used to get commit status information in a repository.
type StatusAPI struct {
	RepositoryInfo
}

// BranchesAPI is used to get information about a repository's branches.
type BranchesAPI struct {
	RepositoryInfo
}

// RepositoryAPI is used to get information about a repository.
type RepositoryAPI struct {
	RepositoryInfo
}

// ContentsAPI is used to get information about the contents of a file in a repository.
type ContentsAPI struct {
	RepositoryInfo
}

// RefsAPI is used to get, create, update and delete refs in a repository.
type RefsAPI struct {
	RepositoryInfo
}

// AuthenticatedUser contains information about the current authenticated user.
type AuthenticatedUser struct {
	Login             string    `json:"login"`
	ID                int       `json:"id"`
	AvatarURL         string    `json:"avatar_url"`
	GravatarID        string    `json:"gravatar_id"`
	URL               string    `json:"url"`
	HTMLURL           string    `json:"html_url"`
	FollowersURL      string    `json:"followers_url"`
	FollowingURL      string    `json:"following_url"`
	GistsURL          string    `json:"gists_url"`
	StarredURL        string    `json:"starred_url"`
	SubscriptionsURL  string    `json:"subscriptions_url"`
	OrganizationsURL  string    `json:"organizations_url"`
	ReposURL          string    `json:"repos_url"`
	EventsURL         string    `json:"events_url"`
	ReceivedEventsURL string    `json:"received_events_url"`
	Type              string    `json:"type"`
	SiteAdmin         bool      `json:"site_admin"`
	Name              string    `json:"name"`
	Company           string    `json:"company"`
	Blog              string    `json:"blog"`
	Location          string    `json:"location"`
	Email             string    `json:"email"`
	Hireable          bool      `json:"hireable"`
	Bio               string    `json:"bio"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	TotalPrivateRepos int       `json:"total_private_repos"`
	OwnedPrivateRepos int       `json:"owned_private_repos"`
	PrivateGists      int       `json:"private_gists"`
	DiskUsage         int       `json:"disk_usage"`
	Collaborators     int       `json:"collaborators"`
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
	gitHubAPI.Refs = RefsAPI{RepositoryInfo: repositoryInfo}

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

// GetOAuthScopes returns the current authenticated user's OAuth scopes.
func GetOAuthScopes(baseURL, authToken string) ([]string, error) {
	apiInfo := APIInfo{BaseURL: baseURL, OAuth2Token: authToken}
	url := apiInfo.addBaseURL("/user")
	resp, err := apiInfo.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if _, err = io.Copy(ioutil.Discard, resp.Body); err != nil {
		return nil, err
	}

	scopesHeader := resp.Header.Get("x-oauth-scopes")
	var scopes []string
	for _, s := range strings.Split(scopesHeader, ",") {
		s = strings.TrimSpace(s)
		if s != "" {
			scopes = append(scopes, s)
		}
	}

	return scopes, nil
}

// OrgSummary contains organization summary information.
type OrgSummary struct {
	Login            string `json:"login"`
	ID               int    `json:"id"`
	URL              string `json:"url"`
	ReposURL         string `json:"repos_url"`
	EventsURL        string `json:"events_url"`
	HooksURL         string `json:"hooks_url"`
	IssuesURL        string `json:"issues_url"`
	MembersURL       string `json:"members_url"`
	PublicMembersURL string `json:"public_members_url"`
	AvatarURL        string `json:"avatar_url"`
	Description      string `json:"description"`
}

// GetOrganizations returns Organization summary information from /organizations.
//
// See https://developer.github.com/v3/orgs/#list-all-organizations.
func GetOrganizations(baseURL, authToken string, since int) ([]OrgSummary, error) {
	apiInfo := APIInfo{BaseURL: baseURL, OAuth2Token: authToken}
	url := apiInfo.addBaseURL(fmt.Sprintf("/organizations?since=%d", since))
	resp, err := apiInfo.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var orgs []OrgSummary

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&orgs); err != nil {
		return nil, err
	}

	return orgs, nil
}

// Is404 returns true if the error is an HTTP 404.
func Is404(err error) bool {
	return IsHTTPError(err, 404)
}

// IsHTTPError returns true if the error is an HTTP error with the specified status code.
func IsHTTPError(err error, statusCode int) bool {
	if err == nil {
		return false
	}
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

func (apiInfo *APIInfo) doHTTPRequest(method, url string, body *string, acceptHeader string) (*http.Response, error) {
	req, err := apiInfo.getHTTPRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if acceptHeader != "" {
		req.Header.Set("Accept", acceptHeader)
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
	return apiInfo.doHTTPRequest("DELETE", url, nil, "")
}

func (apiInfo *APIInfo) httpGet(url string) (*http.Response, error) {
	return apiInfo.doHTTPRequest("GET", url, nil, "")
}

func (apiInfo *APIInfo) httpPatch(url, body string) (*http.Response, error) {
	return apiInfo.doHTTPRequest("PATCH", url, &body, "")
}

func (apiInfo *APIInfo) httpPost(url, body string) (*http.Response, error) {
	return apiInfo.doHTTPRequest("POST", url, &body, "")
}
