package ghapi

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const ctLayout = "2006-01-02T15:04:05Z"

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
		parsedTime, err := time.Parse(ctLayout, string(b))
		if err != nil {
			return err
		}
		ct.Time = parsedTime
	} else if len(b) > 0 {
		secs, err := strconv.Atoi(string(b))
		if err != nil {
			return err
		}
		ct.Time = time.Unix(int64(secs), 0)
	}
	return nil
}

// RepositoryPayload contains information about the repository
type RepositoryPayload struct {
	ID int `json:"id"`
	// The short name of the repository.
	Name string `json:"name"`
	// The name of the repository including the owner (user/organization).
	FullName string `json:"full_name"`
	// TODO: What if the owner is an organization? Owner.Type field?
	Owner            User       `json:"owner"`
	Private          bool       `json:"private"`
	HTMLURL          string     `json:"html_url"`
	Fork             bool       `json:"fork"`
	URL              string     `json:"url"`
	ForksURL         string     `json:"forks_url"`
	KeysURL          string     `json:"keys"`
	CollaboratorsURL string     `json:"collaborators_url"`
	TeamsURL         string     `json:"teams_url"`
	HooksURL         string     `json:"hooks_url"`
	IssueEventsURL   string     `json:"issue_events_url"`
	EventURL         string     `json:"events_url"`
	AssigneesURL     string     `json:"assignees_url"`
	BranchesURL      string     `json:"branches_url"`
	TagsURL          string     `json:"tags_url"`
	BlobsURL         string     `json:"blobs_url"`
	GitTagsURL       string     `json:"git_tags_url"`
	GitRefsURL       string     `json:"git_refs_url"`
	TreesURL         string     `json:"trees_url"`
	StatusesURL      string     `json:"statuses_url"`
	LanguagesURL     string     `json:"languages_url"`
	StargazersURL    string     `json:"stargazers_url"`
	ContributorsURL  string     `json:"contributors_url"`
	SubscribersURL   string     `json:"subscribers_url"`
	SubscriptionURL  string     `json:"subscription_url"`
	CommitsURL       string     `json:"commits_url"`
	GitCommitsURL    string     `json:"git_commits_url"`
	CommentsURL      string     `json:"comments_url"`
	IssueCommentURL  string     `json:"issue_comment_url"`
	ContentsURL      string     `json:"contents_url"`
	CompareURL       string     `json:"compare_url"`
	MergesURL        string     `json:"merges_url"`
	ArchiveURL       string     `json:"archive_url"`
	DownloadsURL     string     `json:"downloads_url"`
	IssuesURL        string     `json:"issues_url"`
	PullsURL         string     `json:"pulls_url"`
	MilestonesURL    string     `json:"milestones_url"`
	NotificationsURL string     `json:"notifications_url"`
	LabelsURL        string     `json:"labels_url"`
	ReleasesURL      string     `json:"releases_url"`
	CreatedAt        CustomTime `json:"created_at"`
	UpdatedAt        CustomTime `json:"updated_at"`
	PushedAt         CustomTime `json:"pushed_at"`
	GitURL           string     `json:"git_url"`
	SSHURL           string     `json:"ssh_url"`
	CloneURL         string     `json:"clone_url"`
	SVNURL           string     `json:"svn_url"`
	// TODO: can be null
	HomePage        string `json:"homepage"`
	Size            int    `json:"size"`
	StargazersCount int    `json:"stargazers_count"`
	WatchersCount   int    `json:"watchers_count"`
	// TODO: can be null
	Language     string `json:"language"`
	HasIssues    bool   `json:"has_issues"`
	HasDownloads bool   `json:"has_downloads"`
	HasWiki      bool   `json:"has_wiki"`
	HasPages     bool   `json:"has_pages"`
	ForksCount   int    `json:"forks_count"`
	// TODO: can be null
	MirrorURL       string `json:"mirror_url"`
	OpenIssuesCount int    `json:"open_issues_count"`
	Forks           int    `json:"forks"`
	OpenIssues      int    `json:"open_issues"`
	Watchers        int    `json:"watchers"`
	DefaultBranch   string `json:"default_branch"`
}

type ForkResponse struct {
	ID    int `json:"id"`
	Owner struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"owner"`
	Name             string     `json:"name"`
	FullName         string     `json:"full_name"`
	Description      string     `json:"description"`
	Private          bool       `json:"private"`
	Fork             bool       `json:"fork"`
	URL              string     `json:"url"`
	HTMLURL          string     `json:"html_url"`
	ArchiveURL       string     `json:"archive_url"`
	AssigneesURL     string     `json:"assignees_url"`
	BlobsURL         string     `json:"blobs_url"`
	BranchesURL      string     `json:"branches_url"`
	CloneURL         string     `json:"clone_url"`
	CollaboratorsURL string     `json:"collaborators_url"`
	CommentsURL      string     `json:"comments_url"`
	CommitsURL       string     `json:"commits_url"`
	CompareURL       string     `json:"compare_url"`
	ContentsURL      string     `json:"contents_url"`
	ContributorsURL  string     `json:"contributors_url"`
	DeploymentsURL   string     `json:"deployments_url"`
	DownloadsURL     string     `json:"downloads_url"`
	EventsURL        string     `json:"events_url"`
	ForksURL         string     `json:"forks_url"`
	GitCommitsURL    string     `json:"git_commits_url"`
	GitRefsURL       string     `json:"git_refs_url"`
	GitTagsURL       string     `json:"git_tags_url"`
	GitURL           string     `json:"git_url"`
	HooksURL         string     `json:"hooks_url"`
	IssueCommentURL  string     `json:"issue_comment_url"`
	IssueEventsURL   string     `json:"issue_events_url"`
	IssuesURL        string     `json:"issues_url"`
	KeysURL          string     `json:"keys_url"`
	LabelsURL        string     `json:"labels_url"`
	LanguagesURL     string     `json:"languages_url"`
	MergesURL        string     `json:"merges_url"`
	MilestonesURL    string     `json:"milestones_url"`
	MirrorURL        string     `json:"mirror_url"`
	NotificationsURL string     `json:"notifications_url"`
	PullsURL         string     `json:"pulls_url"`
	ReleasesURL      string     `json:"releases_url"`
	SSHURL           string     `json:"ssh_url"`
	StargazersURL    string     `json:"stargazers_url"`
	StatusesURL      string     `json:"statuses_url"`
	SubscribersURL   string     `json:"subscribers_url"`
	SubscriptionURL  string     `json:"subscription_url"`
	SvnURL           string     `json:"svn_url"`
	TagsURL          string     `json:"tags_url"`
	TeamsURL         string     `json:"teams_url"`
	TreesURL         string     `json:"trees_url"`
	Homepage         string     `json:"homepage"`
	Language         string     `json:"language"`
	ForksCount       int        `json:"forks_count"`
	StargazersCount  int        `json:"stargazers_count"`
	WatchersCount    int        `json:"watchers_count"`
	Size             int        `json:"size"`
	DefaultBranch    string     `json:"default_branch"`
	OpenIssuesCount  int        `json:"open_issues_count"`
	HasIssues        bool       `json:"has_issues"`
	HasWiki          bool       `json:"has_wiki"`
	HasPages         bool       `json:"has_pages"`
	HasDownloads     bool       `json:"has_downloads"`
	PushedAt         CustomTime `json:"pushed_at"`
	CreatedAt        CustomTime `json:"created_at"`
	UpdatedAt        CustomTime `json:"updated_at"`
	Permissions      struct {
		Admin bool `json:"admin"`
		Push  bool `json:"push"`
		Pull  bool `json:"pull"`
	} `json:"permissions"`
}

type RepositoryCommit struct {
	URL         string `json:"url"`
	SHA         string `json:"sha"`
	HTMLURL     string `json:"html_url"`
	CommentsURL string `json:"comments_url"`
	Commit      struct {
		URL    string `json:"url"`
		Author struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"author"`
		Committer struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    struct {
			URL string `json:"url"`
			SHA string `json:"sha"`
		} `json:"tree"`
		CommentCount int `json:"comment_count"`
		Verification struct {
			Verified  bool   `json:"verified"`
			Reason    string `json:"reason"`
			Signature string `json:"signature"`
			Payload   string `json:"payload"`
		} `json:"verification"`
	} `json:"commit"`
	Author struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	Committer struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"committer"`
	Parents []struct {
		URL string `json:"url"`
		SHA string `json:"sha"`
	} `json:"parents"`
}

// Get returns the repository information.
func (api *RepositoryAPI) Get() (*RepositoryPayload, error) {
	url := api.getURL("/repos/:owner/:repo")

	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repository RepositoryPayload

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&repository); err != nil {
		return nil, err
	}

	return &repository, nil
}

// ForkAsync forks the repository into the user's account.
// Forking a Repository happens asynchronously.
// Therefore, you may have to wait a short period before accessing the git objects.
func (api *RepositoryAPI) ForkAsync() (*ForkResponse, error) {
	url := api.getURL("/repos/:owner/:repo/forks")

	resp, err := api.httpPost(url, "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var forkResponse ForkResponse

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&forkResponse); err != nil {
		return nil, err
	}

	return &forkResponse, nil
}

// Fork performs a synchronous fork operation with a timeout period. See ForkAsync and IsReady methods for
// more details.
func (api *RepositoryAPI) Fork(timeout time.Duration) (*ForkResponse, error) {
	timeoutChan := time.After(timeout)

	forkResponse, err := api.ForkAsync()
	if err != nil {
		return nil, err
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if ready, err := api.IsReady(); err != nil {
				return forkResponse, err
			} else if ready {
				return forkResponse, nil
			}
		case <-timeoutChan:
			return forkResponse, fmt.Errorf("timeout (%v) waiting for fork of %s/%s to complete",
				timeout, api.Owner, api.Repository)
		}
	}
}

// GetCommits returns commits for the repository.
func (api *RepositoryAPI) GetCommits(page int) ([]RepositoryCommit, error) {
	url := api.getURL("/repos/:owner/:repo/commits")

	// TODO (judwhite)
	// Link: <https://api.github.com/resource?page=2>; rel="next",
	//       <https://api.github.com/resource?page=5>; rel="last"

	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	commits := make([]RepositoryCommit, 0)

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&commits); err != nil {
		return nil, err
	}

	return commits, nil
}

// Exists returns true if the repository exists.
func (api *RepositoryAPI) Exists() (bool, error) {
	_, err := api.Get()
	if err != nil {
		if Is404(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// IsReady returns true if the repository is "ready", which usually means a fork operation has completed.
// It does this by trying to list the first page of commits on the repository; if "/repos/:owner/:repo/commits"
// returns 409, the repository is not yet ready.
//
// This method is used internally by the Fork method. It can be used to check if a call to ForkAsync has completed
// the fork operation.
func (api *RepositoryAPI) IsReady() (bool, error) {
	_, err := api.GetCommits(1)
	if err != nil {
		if IsHTTPError(err, 409) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
