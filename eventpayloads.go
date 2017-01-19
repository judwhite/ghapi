package ghapi

import "time"

// GitHubEventPayload contains Repository and Sender information. It's not sent by any GitHub events but is used
// as an embedded type in other event payloads.
type GitHubEventPayload struct {
	Sender     User `json:"sender"`
	Repository struct {
		ID               int       `json:"id"`
		Name             string    `json:"name"`
		FullName         string    `json:"full_name"`
		Owner            User      `json:"owner"`
		Private          bool      `json:"private"`
		HTMLURL          string    `json:"html_url"`
		Description      string    `json:"description"`
		Fork             bool      `json:"fork"`
		URL              string    `json:"url"`
		ForksURL         string    `json:"forks_url"`
		KeysURL          string    `json:"keys_url"`
		CollaboratorsURL string    `json:"collaborators_url"`
		TeamsURL         string    `json:"teams_url"`
		HooksURL         string    `json:"hooks_url"`
		IssueEventsURL   string    `json:"issue_events_url"`
		EventsURL        string    `json:"events_url"`
		AssigneesURL     string    `json:"assignees_url"`
		BranchesURL      string    `json:"branches_url"`
		TagsURL          string    `json:"tags_url"`
		BlobsURL         string    `json:"blobs_url"`
		GitTagsURL       string    `json:"git_tags_url"`
		GitRefsURL       string    `json:"git_refs_url"`
		TreesURL         string    `json:"trees_url"`
		StatusesURL      string    `json:"statuses_url"`
		LanguagesURL     string    `json:"languages_url"`
		StargazersURL    string    `json:"stargazers_url"`
		ContributorsURL  string    `json:"contributors_url"`
		SubscribersURL   string    `json:"subscribers_url"`
		SubscriptionURL  string    `json:"subscription_url"`
		CommitsURL       string    `json:"commits_url"`
		GitCommitsURL    string    `json:"git_commits_url"`
		CommentsURL      string    `json:"comments_url"`
		IssueCommentURL  string    `json:"issue_comment_url"`
		ContentsURL      string    `json:"contents_url"`
		CompareURL       string    `json:"compare_url"`
		MergesURL        string    `json:"merges_url"`
		ArchiveURL       string    `json:"archive_url"`
		DownloadsURL     string    `json:"downloads_url"`
		IssuesURL        string    `json:"issues_url"`
		PullsURL         string    `json:"pulls_url"`
		MilestonesURL    string    `json:"milestones_url"`
		NotificationsURL string    `json:"notifications_url"`
		LabelsURL        string    `json:"labels_url"`
		ReleasesURL      string    `json:"releases_url"`
		DeploymentsURL   string    `json:"deployments_url"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		PushedAt         time.Time `json:"pushed_at"`
		GitURL           string    `json:"git_url"`
		SSHURL           string    `json:"ssh_url"`
		CloneURL         string    `json:"clone_url"`
		SVNURL           *string   `json:"svn_url"`
		Homepage         string    `json:"homepage"`
		Size             int       `json:"size"`
		StargazersCount  int       `json:"stargazers_count"`
		WatchersCount    int       `json:"watchers_count"`
		Language         *string   `json:"language"`
		HasIssues        bool      `json:"has_issues"`
		HasDownloads     bool      `json:"has_downloads"`
		HasWiki          bool      `json:"has_wiki"`
		HasPages         bool      `json:"has_pages"`
		ForksCount       int       `json:"forks_count"`
		MirrorURL        *string   `json:"mirror_url"`
		OpenIssuesCount  int       `json:"open_issues_count"`
		Forks            int       `json:"forks"`
		OpenIssues       int       `json:"open_issues"`
		Watchers         int       `json:"watchers"`
		DefaultBranch    string    `json:"default_branch"`
	} `json:"repository"`
}

// PingEventPayload is received from the Ping Event.
// See https://developer.github.com/webhooks/#ping-event.
type PingEventPayload struct {
	GitHubEventPayload
	Zen    string `json:"zen"`
	HookID int    `json:"hook_id"`
	Hook   struct {
		Type   string   `json:"type"`
		ID     int      `json:"id"`
		Name   string   `json:"name"`
		Active bool     `json:"active"`
		Events []string `json:"events"`
		Config struct {
			ContentType string `json:"content_type"`
			InsecureSSL string `json:"insecure_ssl"`
			Secret      string `json:"secret"`
			URL         string `json:"url"`
		} `json:"config"`
		UpdatedAt    time.Time `json:"updated_at"`
		CreatedAt    time.Time `json:"created_at"`
		URL          string    `json:"url"`
		TestURL      string    `json:"test_url"`
		PingURL      string    `json:"ping_url"`
		LastResponse struct {
			Code    *int    `json:"code"`
			Status  string  `json:"status"`
			Message *string `json:"message"`
		} `json:"last_response"`
	} `json:"hook"`
}

// PullRequestEventPayload is received from the Pull Request Event.
// See https://developer.github.com/v3/activity/events/types/#pullrequestevent.
type PullRequestEventPayload struct {
	GitHubEventPayload
	Action      PullRequestAction   `json:"action"`
	Number      int                 `json:"number"`
	PullRequest PullRequestResponse `json:"pull_request"`
}

// IssueCommentEventPayload is received from the Issue Comment Event.
// See https://developer.github.com/v3/activity/events/types/#issuecommentevent.
type IssueCommentEventPayload struct {
	GitHubEventPayload
	Action  string               `json:"action"`
	Issue   IssueResponse        `json:"issue"`
	Comment IssueCommentResponse `json:"comment"`
}

// PushEventPayload is received from the Push Event.
// See https://developer.github.com/v3/activity/events/types/#pushevent.
type PushEventPayload struct {
	Sender  User    `json:"sender"`
	Ref     string  `json:"ref"`
	Before  string  `json:"before"`
	After   string  `json:"after"`
	Created bool    `json:"created"`
	Deleted bool    `json:"deleted"`
	Forced  bool    `json:"forced"`
	BaseRef *string `json:"base_ref"`
	Compare string  `json:"compare"`
	Commits []struct {
		ID        string `json:"id"`
		Distinct  bool   `json:"distinct"`
		Message   string `json:"message"`
		Timestamp string `json:"timestamp"`
		URL       string `json:"url"`
		Author    struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"author"`
		Committer struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"committer"`
		Added    []string `json:"added"`
		Removed  []string `json:"removed"`
		Modified []string `json:"modified"`
	} `json:"commits"`
	HeadCommit struct {
		ID        string `json:"id"`
		Distinct  bool   `json:"distinct"`
		Message   string `json:"message"`
		Timestamp string `json:"timestamp"`
		URL       string `json:"url"`
		Author    struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"author"`
		Committer struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"committer"`
		Added    []string `json:"added"`
		Removed  []string `json:"removed"`
		Modified []string `json:"modified"`
	} `json:"head_commit"`
	Repository struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Owner    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"owner"`
		Private          bool      `json:"private"`
		HTMLURL          string    `json:"html_url"`
		Description      string    `json:"description"`
		Fork             bool      `json:"fork"`
		URL              string    `json:"url"`
		ForksURL         string    `json:"forks_url"`
		KeysURL          string    `json:"keys_url"`
		CollaboratorsURL string    `json:"collaborators_url"`
		TeamsURL         string    `json:"teams_url"`
		HooksURL         string    `json:"hooks_url"`
		IssueEventsURL   string    `json:"issue_events_url"`
		EventsURL        string    `json:"events_url"`
		AssigneesURL     string    `json:"assignees_url"`
		BranchesURL      string    `json:"branches_url"`
		TagsURL          string    `json:"tags_url"`
		BlobsURL         string    `json:"blobs_url"`
		GitTagsURL       string    `json:"git_tags_url"`
		GitRefsURL       string    `json:"git_refs_url"`
		TreesURL         string    `json:"trees_url"`
		StatusesURL      string    `json:"statuses_url"`
		LanguagesURL     string    `json:"languages_url"`
		StargazersURL    string    `json:"stargazers_url"`
		ContributorsURL  string    `json:"contributors_url"`
		SubscribersURL   string    `json:"subscribers_url"`
		SubscriptionURL  string    `json:"subscription_url"`
		CommitsURL       string    `json:"commits_url"`
		GitCommitsURL    string    `json:"git_commits_url"`
		CommentsURL      string    `json:"comments_url"`
		IssueCommentURL  string    `json:"issue_comment_url"`
		ContentsURL      string    `json:"contents_url"`
		CompareURL       string    `json:"compare_url"`
		MergesURL        string    `json:"merges_url"`
		ArchiveURL       string    `json:"archive_url"`
		DownloadsURL     string    `json:"downloads_url"`
		IssuesURL        string    `json:"issues_url"`
		PullsURL         string    `json:"pulls_url"`
		MilestonesURL    string    `json:"milestones_url"`
		NotificationsURL string    `json:"notifications_url"`
		LabelsURL        string    `json:"labels_url"`
		ReleasesURL      string    `json:"releases_url"`
		DeploymentsURL   string    `json:"deployments_url"`
		CreatedAt        int       `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		PushedAt         int       `json:"pushed_at"`
		GitURL           string    `json:"git_url"`
		SSHURL           string    `json:"ssh_url"`
		CloneURL         string    `json:"clone_url"`
		SVNURL           string    `json:"svn_url"`
		Homepage         *string   `json:"homepage"`
		Size             int       `json:"size"`
		StargazersCount  int       `json:"stargazers_count"`
		WatchersCount    int       `json:"watchers_count"`
		Language         *string   `json:"language"`
		HasIssues        bool      `json:"has_issues"`
		HasDownloads     bool      `json:"has_downloads"`
		HasWiki          bool      `json:"has_wiki"`
		HasPages         bool      `json:"has_pages"`
		ForksCount       int       `json:"forks_count"`
		MirrorURL        *string   `json:"mirror_url"`
		OpenIssuesCount  int       `json:"open_issues_count"`
		Forks            int       `json:"forks"`
		OpenIssues       int       `json:"open_issues"`
		Watchers         int       `json:"watchers"`
		DefaultBranch    string    `json:"default_branch"`
		Stargazers       int       `json:"stargazers"`
		MasterBranch     string    `json:"master_branch"`
		Organization     string    `json:"organization"`
	} `json:"repository"`
	Pusher struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"pusher"`
}
