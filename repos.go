package ghapi

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"time"
)

/*const ctLayout = "2006-01-02T15:04:05Z"

// CustomTime embeds time.Time and supports unmarshalling from either date/time formats or epoch offset in seconds.
type CustomTime struct {
	time.Time
}

// UnmarshalJSON unmarshals the token to a time.Time value.
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
}*/

// RepositoryResponse contains information about the repository
type RepositoryResponse struct {
	ID               int       `json:"id"`
	Owner            User      `json:"owner"`
	Name             string    `json:"name"`
	FullName         string    `json:"full_name"`
	Description      string    `json:"description"`
	Private          bool      `json:"private"`
	Fork             bool      `json:"fork"`
	URL              string    `json:"url"`
	HTMLURL          string    `json:"html_url"`
	ArchiveURL       string    `json:"archive_url"`
	AssigneesURL     string    `json:"assignees_url"`
	BlobsURL         string    `json:"blobs_url"`
	BranchesURL      string    `json:"branches_url"`
	CloneURL         string    `json:"clone_url"`
	CollaboratorsURL string    `json:"collaborators_url"`
	CommentsURL      string    `json:"comments_url"`
	CommitsURL       string    `json:"commits_url"`
	CompareURL       string    `json:"compare_url"`
	ContentsURL      string    `json:"contents_url"`
	ContributorsURL  string    `json:"contributors_url"`
	DeploymentsURL   string    `json:"deployments_url"`
	DownloadsURL     string    `json:"downloads_url"`
	EventsURL        string    `json:"events_url"`
	ForksURL         string    `json:"forks_url"`
	GitCommitsURL    string    `json:"git_commits_url"`
	GitRefsURL       string    `json:"git_refs_url"`
	GitTagsURL       string    `json:"git_tags_url"`
	GitURL           string    `json:"git_url"`
	HooksURL         string    `json:"hooks_url"`
	IssueCommentURL  string    `json:"issue_comment_url"`
	IssueEventsURL   string    `json:"issue_events_url"`
	IssuesURL        string    `json:"issues_url"`
	KeysURL          string    `json:"keys_url"`
	LabelsURL        string    `json:"labels_url"`
	LanguagesURL     string    `json:"languages_url"`
	MergesURL        string    `json:"merges_url"`
	MilestonesURL    string    `json:"milestones_url"`
	MirrorURL        *string   `json:"mirror_url"`
	NotificationsURL string    `json:"notifications_url"`
	PullsURL         string    `json:"pulls_url"`
	ReleasesURL      string    `json:"releases_url"`
	SSHURL           string    `json:"ssh_url"`
	StargazersURL    string    `json:"stargazers_url"`
	StatusesURL      string    `json:"statuses_url"`
	SubscribersURL   string    `json:"subscribers_url"`
	SubscriptionURL  string    `json:"subscription_url"`
	SvnURL           string    `json:"svn_url"`
	TagsURL          string    `json:"tags_url"`
	TeamsURL         string    `json:"teams_url"`
	TreesURL         string    `json:"trees_url"`
	Homepage         *string   `json:"homepage"`
	Language         *string   `json:"language"`
	ForksCount       int       `json:"forks_count"`
	StargazersCount  int       `json:"stargazers_count"`
	WatchersCount    int       `json:"watchers_count"`
	Size             int       `json:"size"`
	DefaultBranch    string    `json:"default_branch"`
	OpenIssuesCount  int       `json:"open_issues_count"`
	HasIssues        bool      `json:"has_issues"`
	HasWiki          bool      `json:"has_wiki"`
	HasPages         bool      `json:"has_pages"`
	HasDownloads     bool      `json:"has_downloads"`
	PushedAt         time.Time `json:"pushed_at"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Permissions      struct {
		Admin bool `json:"admin"`
		Push  bool `json:"push"`
		Pull  bool `json:"pull"`
	} `json:"permissions"`
	SubscribersCount int `json:"subscribers_count"`
	Organization     struct {
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
	} `json:"organization"`
	Parent struct {
		ID               int       `json:"id"`
		Owner            User      `json:"owner"`
		Name             string    `json:"name"`
		FullName         string    `json:"full_name"`
		Description      string    `json:"description"`
		Private          bool      `json:"private"`
		Fork             bool      `json:"fork"`
		URL              string    `json:"url"`
		HTMLURL          string    `json:"html_url"`
		ArchiveURL       string    `json:"archive_url"`
		AssigneesURL     string    `json:"assignees_url"`
		BlobsURL         string    `json:"blobs_url"`
		BranchesURL      string    `json:"branches_url"`
		CloneURL         string    `json:"clone_url"`
		CollaboratorsURL string    `json:"collaborators_url"`
		CommentsURL      string    `json:"comments_url"`
		CommitsURL       string    `json:"commits_url"`
		CompareURL       string    `json:"compare_url"`
		ContentsURL      string    `json:"contents_url"`
		ContributorsURL  string    `json:"contributors_url"`
		DeploymentsURL   string    `json:"deployments_url"`
		DownloadsURL     string    `json:"downloads_url"`
		EventsURL        string    `json:"events_url"`
		ForksURL         string    `json:"forks_url"`
		GitCommitsURL    string    `json:"git_commits_url"`
		GitRefsURL       string    `json:"git_refs_url"`
		GitTagsURL       string    `json:"git_tags_url"`
		GitURL           string    `json:"git_url"`
		HooksURL         string    `json:"hooks_url"`
		IssueCommentURL  string    `json:"issue_comment_url"`
		IssueEventsURL   string    `json:"issue_events_url"`
		IssuesURL        string    `json:"issues_url"`
		KeysURL          string    `json:"keys_url"`
		LabelsURL        string    `json:"labels_url"`
		LanguagesURL     string    `json:"languages_url"`
		MergesURL        string    `json:"merges_url"`
		MilestonesURL    string    `json:"milestones_url"`
		MirrorURL        *string   `json:"mirror_url"`
		NotificationsURL string    `json:"notifications_url"`
		PullsURL         string    `json:"pulls_url"`
		ReleasesURL      string    `json:"releases_url"`
		SSHURL           string    `json:"ssh_url"`
		StargazersURL    string    `json:"stargazers_url"`
		StatusesURL      string    `json:"statuses_url"`
		SubscribersURL   string    `json:"subscribers_url"`
		SubscriptionURL  string    `json:"subscription_url"`
		SvnURL           string    `json:"svn_url"`
		TagsURL          string    `json:"tags_url"`
		TeamsURL         string    `json:"teams_url"`
		TreesURL         string    `json:"trees_url"`
		Homepage         *string   `json:"homepage"`
		Language         *string   `json:"language"`
		ForksCount       int       `json:"forks_count"`
		StargazersCount  int       `json:"stargazers_count"`
		WatchersCount    int       `json:"watchers_count"`
		Size             int       `json:"size"`
		DefaultBranch    string    `json:"default_branch"`
		OpenIssuesCount  int       `json:"open_issues_count"`
		HasIssues        bool      `json:"has_issues"`
		HasWiki          bool      `json:"has_wiki"`
		HasPages         bool      `json:"has_pages"`
		HasDownloads     bool      `json:"has_downloads"`
		PushedAt         time.Time `json:"pushed_at"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		Permissions      struct {
			Admin bool `json:"admin"`
			Push  bool `json:"push"`
			Pull  bool `json:"pull"`
		} `json:"permissions"`
	} `json:"parent"`
	Source struct {
		ID               int       `json:"id"`
		Owner            User      `json:"owner"`
		Name             string    `json:"name"`
		FullName         string    `json:"full_name"`
		Description      string    `json:"description"`
		Private          bool      `json:"private"`
		Fork             bool      `json:"fork"`
		URL              string    `json:"url"`
		HTMLURL          string    `json:"html_url"`
		ArchiveURL       string    `json:"archive_url"`
		AssigneesURL     string    `json:"assignees_url"`
		BlobsURL         string    `json:"blobs_url"`
		BranchesURL      string    `json:"branches_url"`
		CloneURL         string    `json:"clone_url"`
		CollaboratorsURL string    `json:"collaborators_url"`
		CommentsURL      string    `json:"comments_url"`
		CommitsURL       string    `json:"commits_url"`
		CompareURL       string    `json:"compare_url"`
		ContentsURL      string    `json:"contents_url"`
		ContributorsURL  string    `json:"contributors_url"`
		DeploymentsURL   string    `json:"deployments_url"`
		DownloadsURL     string    `json:"downloads_url"`
		EventsURL        string    `json:"events_url"`
		ForksURL         string    `json:"forks_url"`
		GitCommitsURL    string    `json:"git_commits_url"`
		GitRefsURL       string    `json:"git_refs_url"`
		GitTagsURL       string    `json:"git_tags_url"`
		GitURL           string    `json:"git_url"`
		HooksURL         string    `json:"hooks_url"`
		IssueCommentURL  string    `json:"issue_comment_url"`
		IssueEventsURL   string    `json:"issue_events_url"`
		IssuesURL        string    `json:"issues_url"`
		KeysURL          string    `json:"keys_url"`
		LabelsURL        string    `json:"labels_url"`
		LanguagesURL     string    `json:"languages_url"`
		MergesURL        string    `json:"merges_url"`
		MilestonesURL    string    `json:"milestones_url"`
		MirrorURL        *string   `json:"mirror_url"`
		NotificationsURL string    `json:"notifications_url"`
		PullsURL         string    `json:"pulls_url"`
		ReleasesURL      string    `json:"releases_url"`
		SSHURL           string    `json:"ssh_url"`
		StargazersURL    string    `json:"stargazers_url"`
		StatusesURL      string    `json:"statuses_url"`
		SubscribersURL   string    `json:"subscribers_url"`
		SubscriptionURL  string    `json:"subscription_url"`
		SvnURL           string    `json:"svn_url"`
		TagsURL          string    `json:"tags_url"`
		TeamsURL         string    `json:"teams_url"`
		TreesURL         string    `json:"trees_url"`
		Homepage         *string   `json:"homepage"`
		Language         *string   `json:"language"`
		ForksCount       int       `json:"forks_count"`
		StargazersCount  int       `json:"stargazers_count"`
		WatchersCount    int       `json:"watchers_count"`
		Size             int       `json:"size"`
		DefaultBranch    string    `json:"default_branch"`
		OpenIssuesCount  int       `json:"open_issues_count"`
		HasIssues        bool      `json:"has_issues"`
		HasWiki          bool      `json:"has_wiki"`
		HasPages         bool      `json:"has_pages"`
		HasDownloads     bool      `json:"has_downloads"`
		PushedAt         time.Time `json:"pushed_at"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		Permissions      struct {
			Admin bool `json:"admin"`
			Push  bool `json:"push"`
			Pull  bool `json:"pull"`
		} `json:"permissions"`
	} `json:"source"`
}

// ForkResponse is returned by RepositoryAPI.Fork and RepositoryAPI.ForkAsync.
type ForkResponse struct {
	ID               int       `json:"id"`
	Owner            User      `json:"owner"`
	Name             string    `json:"name"`
	FullName         string    `json:"full_name"`
	Description      string    `json:"description"`
	Private          bool      `json:"private"`
	Fork             bool      `json:"fork"`
	URL              string    `json:"url"`
	HTMLURL          string    `json:"html_url"`
	ArchiveURL       string    `json:"archive_url"`
	AssigneesURL     string    `json:"assignees_url"`
	BlobsURL         string    `json:"blobs_url"`
	BranchesURL      string    `json:"branches_url"`
	CloneURL         string    `json:"clone_url"`
	CollaboratorsURL string    `json:"collaborators_url"`
	CommentsURL      string    `json:"comments_url"`
	CommitsURL       string    `json:"commits_url"`
	CompareURL       string    `json:"compare_url"`
	ContentsURL      string    `json:"contents_url"`
	ContributorsURL  string    `json:"contributors_url"`
	DeploymentsURL   string    `json:"deployments_url"`
	DownloadsURL     string    `json:"downloads_url"`
	EventsURL        string    `json:"events_url"`
	ForksURL         string    `json:"forks_url"`
	GitCommitsURL    string    `json:"git_commits_url"`
	GitRefsURL       string    `json:"git_refs_url"`
	GitTagsURL       string    `json:"git_tags_url"`
	GitURL           string    `json:"git_url"`
	HooksURL         string    `json:"hooks_url"`
	IssueCommentURL  string    `json:"issue_comment_url"`
	IssueEventsURL   string    `json:"issue_events_url"`
	IssuesURL        string    `json:"issues_url"`
	KeysURL          string    `json:"keys_url"`
	LabelsURL        string    `json:"labels_url"`
	LanguagesURL     string    `json:"languages_url"`
	MergesURL        string    `json:"merges_url"`
	MilestonesURL    string    `json:"milestones_url"`
	MirrorURL        string    `json:"mirror_url"`
	NotificationsURL string    `json:"notifications_url"`
	PullsURL         string    `json:"pulls_url"`
	ReleasesURL      string    `json:"releases_url"`
	SSHURL           string    `json:"ssh_url"`
	StargazersURL    string    `json:"stargazers_url"`
	StatusesURL      string    `json:"statuses_url"`
	SubscribersURL   string    `json:"subscribers_url"`
	SubscriptionURL  string    `json:"subscription_url"`
	SvnURL           string    `json:"svn_url"`
	TagsURL          string    `json:"tags_url"`
	TeamsURL         string    `json:"teams_url"`
	TreesURL         string    `json:"trees_url"`
	Homepage         string    `json:"homepage"`
	Language         string    `json:"language"`
	ForksCount       int       `json:"forks_count"`
	StargazersCount  int       `json:"stargazers_count"`
	WatchersCount    int       `json:"watchers_count"`
	Size             int       `json:"size"`
	DefaultBranch    string    `json:"default_branch"`
	OpenIssuesCount  int       `json:"open_issues_count"`
	HasIssues        bool      `json:"has_issues"`
	HasWiki          bool      `json:"has_wiki"`
	HasPages         bool      `json:"has_pages"`
	HasDownloads     bool      `json:"has_downloads"`
	PushedAt         time.Time `json:"pushed_at"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Permissions      struct {
		Admin bool `json:"admin"`
		Push  bool `json:"push"`
		Pull  bool `json:"pull"`
	} `json:"permissions"`
}

// RepositoryCommit represents a commit summary in a repository.
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
	Author    User `json:"author"`
	Committer User `json:"committer"`
	Parents   []struct {
		URL string `json:"url"`
		SHA string `json:"sha"`
	} `json:"parents"`
}

// Commit represents commit details.
type Commit struct {
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
	Author    User `json:"author"`
	Committer User `json:"committer"`
	Parents   []struct {
		URL string `json:"url"`
		SHA string `json:"sha"`
	} `json:"parents"`
	Stats struct {
		Additions int `json:"additions"`
		Deletions int `json:"deletions"`
		Total     int `json:"total"`
	} `json:"stats"`
	Files []struct {
		Filename  string `json:"filename"`
		Additions int    `json:"additions"`
		Deletions int    `json:"deletions"`
		Changes   int    `json:"changes"`
		Status    string `json:"status"`
		RawURL    string `json:"raw_url"`
		BlobURL   string `json:"blob_url"`
		Patch     string `json:"patch"`
	} `json:"files"`
}

// Get returns the repository information.
func (api *RepositoryAPI) Get() (*RepositoryResponse, error) {
	url := api.getURL("/repos/:owner/:repo")

	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repository RepositoryResponse

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

	commits := []RepositoryCommit{}

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&commits); err != nil {
		return nil, err
	}

	return commits, nil
}

// GetCommit returns commit details for the specified SHA.
func (api *RepositoryAPI) GetCommit(sha string) (*Commit, error) {
	url := api.getURL("/repos/:owner/:repo/commits/" + sha)

	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var commit Commit

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&commit); err != nil {
		return nil, err
	}

	return &commit, nil
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

// CreateLabel creates a label in the repository. color is a 6 character hex code without the leading #.
func (api *RepositoryAPI) CreateLabel(name, color string) error {
	body := struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	}{name, color}

	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	url := api.getURL("/repos/:owner/:repo/labels")

	resp, err := api.httpPost(url, string(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(ioutil.Discard, resp.Body)
	return err
}

// UpdateLabel updates a label in the repository. color is a 6 character hex code without the leading #.
func (api *RepositoryAPI) UpdateLabel(origName, newName, color string) error {
	body := struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	}{newName, color}

	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	apiURL := api.getURL("/repos/:owner/:repo/labels/" + url.PathEscape(origName))

	resp, err := api.httpPatch(apiURL, string(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(ioutil.Discard, resp.Body)
	return err
}

// GetLabels returns all labels for the repository.
func (api *RepositoryAPI) GetLabels() ([]IssueLabel, error) {
	var allLabels []IssueLabel

	for page := 1; ; page++ {
		url := api.getURL(fmt.Sprintf("/repos/:owner/:repo/labels?page=%d", page))

		resp, err := api.httpGet(url)
		if err != nil {
			return nil, err
		}

		var labels []IssueLabel

		j := json.NewDecoder(resp.Body)
		if err = j.Decode(&labels); err != nil {
			resp.Body.Close()
			return nil, err
		}

		resp.Body.Close()

		if len(labels) == 0 {
			break
		}

		allLabels = append(allLabels, labels...)
	}

	return allLabels, nil
}
