package ghapi

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// PullRequestAction represents an action from a "pull_request" GitHub event.
type PullRequestAction string

const (
	// Assigned means a pull request was assigned.
	Assigned PullRequestAction = "assigned"
	// Unassigned means a pull request was unassigned.
	Unassigned PullRequestAction = "unassigned"
	// Labeled means a pull request was labeled.
	Labeled PullRequestAction = "labeled"
	// Unlabeled means a pull request was unlabeled.
	Unlabeled PullRequestAction = "unlabeled"
	// Opened means a pull request was opened.
	Opened PullRequestAction = "opened"
	// Closed means a pull request was closed.
	Closed PullRequestAction = "closed"
	// Reopened means a pull request was reopened.
	Reopened PullRequestAction = "reopened"
	// Synchronize means a pull request was synchronized (updated due to a new push in the branch that the pull
	// request is tracking).
	Synchronize PullRequestAction = "synchronize"
)

// MergeMethod represents the state of GitHub's "merge_method" (merge, squash, rebase).
type MergeMethod string

const (
	// Merge commit method; "merge".
	Merge MergeMethod = "merge"
	// Squash commit method; "squash".
	Squash MergeMethod = "squash"
	// Rebase commit method; "rebase".
	Rebase MergeMethod = "rebase"
)

// MergeRequestResponse contains information about an attempted merge request.
type MergeRequestResponse struct {
	SHA     string `json:"sha"`
	Merged  bool   `json:"merged"`
	Message string `json:"message"`
}

// PullRequestResponse contains information about a pull request.
type PullRequestResponse struct {
	URL            string     `json:"url"`
	ID             int        `json:"id"`
	HTMLURL        string     `json:"html_url"`
	DiffURL        string     `json:"diff_url"`
	PatchURL       string     `json:"patch_url"`
	IssueURL       string     `json:"issue_url"`
	Number         int        `json:"number"`
	State          string     `json:"state"`
	Locked         bool       `json:"locked"`
	Title          string     `json:"title"`
	User           User       `json:"user"`
	Body           string     `json:"body"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	ClosedAt       *time.Time `json:"closed_at"`
	MergedAt       *time.Time `json:"merged_at"`
	MergeCommitSHA string     `json:"merge_commit_sha"`
	Assignee       *User      `json:"assignee"`
	Assignees      []User     `json:"assignees"`
	Milestone      *struct {
		URL          string     `json:"url"`
		HTMLURL      string     `json:"html_url"`
		LabelsURL    string     `json:"labels_url"`
		ID           int        `json:"id"`
		Number       int        `json:"number"`
		State        string     `json:"state"`
		Title        string     `json:"title"`
		Description  string     `json:"description"`
		Creator      User       `json:"creator"`
		OpenIssues   int        `json:"open_issues"`
		ClosedIssues int        `json:"closed_issues"`
		CreatedAt    time.Time  `json:"created_at"`
		UpdatedAt    time.Time  `json:"updated_at"`
		ClosedAt     *time.Time `json:"closed_at"`
		DueOn        *time.Time `json:"due_on"`
	} `json:"milestone"`
	CommitsURL        string `json:"commits_url"`
	ReviewCommentsURL string `json:"review_comments_url"`
	ReviewCommentURL  string `json:"review_comment_url"`
	CommentsURL       string `json:"comments_url"`
	StatusesURL       string `json:"statuses_url"`
	Head              struct {
		Label string `json:"label"`
		Ref   string `json:"ref"`
		SHA   string `json:"sha"`
		User  User   `json:"user"`
		Repo  struct {
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
			SvnURL           string    `json:"svn_url"`
			Homepage         string    `json:"homepage"`
			Size             int       `json:"size"`
			StargazersCount  int       `json:"stargazers_count"`
			WatchersCount    int       `json:"watchers_count"`
			Language         string    `json:"language"`
			HasIssues        bool      `json:"has_issues"`
			HasDownloads     bool      `json:"has_downloads"`
			HasWiki          bool      `json:"has_wiki"`
			HasPages         bool      `json:"has_pages"`
			ForksCount       int       `json:"forks_count"`
			MirrorURL        string    `json:"mirror_url"`
			OpenIssuesCount  int       `json:"open_issues_count"`
			Forks            int       `json:"forks"`
			OpenIssues       int       `json:"open_issues"`
			Watchers         int       `json:"watchers"`
			DefaultBranch    string    `json:"default_branch"`
		} `json:"repo"`
	} `json:"head"`
	Base struct {
		Label string `json:"label"`
		Ref   string `json:"ref"`
		SHA   string `json:"sha"`
		User  User   `json:"user"`
		Repo  struct {
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
			SvnURL           string    `json:"svn_url"`
			Homepage         string    `json:"homepage"`
			Size             int       `json:"size"`
			StargazersCount  int       `json:"stargazers_count"`
			WatchersCount    int       `json:"watchers_count"`
			Language         string    `json:"language"`
			HasIssues        bool      `json:"has_issues"`
			HasDownloads     bool      `json:"has_downloads"`
			HasWiki          bool      `json:"has_wiki"`
			HasPages         bool      `json:"has_pages"`
			ForksCount       int       `json:"forks_count"`
			MirrorURL        string    `json:"mirror_url"`
			OpenIssuesCount  int       `json:"open_issues_count"`
			Forks            int       `json:"forks"`
			OpenIssues       int       `json:"open_issues"`
			Watchers         int       `json:"watchers"`
			DefaultBranch    string    `json:"default_branch"`
		} `json:"repo"`
	} `json:"base"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		Issue struct {
			Href string `json:"href"`
		} `json:"issue"`
		Comments struct {
			Href string `json:"href"`
		} `json:"comments"`
		ReviewComments struct {
			Href string `json:"href"`
		} `json:"review_comments"`
		ReviewComment struct {
			Href string `json:"href"`
		} `json:"review_comment"`
		Commits struct {
			Href string `json:"href"`
		} `json:"commits"`
		Statuses struct {
			Href string `json:"href"`
		} `json:"statuses"`
	} `json:"_links"`
	Merged         bool   `json:"merged"`
	Mergeable      *bool  `json:"mergeable"`
	MergeableState string `json:"mergeable_state"`
	MergedBy       *User  `json:"merged_by"`
	Comments       int    `json:"comments"`
	ReviewComments int    `json:"review_comments"`
	Commits        int    `json:"commits"`
	Additions      int    `json:"additions"`
	Deletions      int    `json:"deletions"`
	ChangedFiles   int    `json:"changed_files"`
}

// CreatePullRequestResponse is returned by PullRequestsAPI.Create.
type CreatePullRequestResponse struct {
	ID                int    `json:"id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	DiffURL           string `json:"diff_url"`
	PatchURL          string `json:"patch_url"`
	IssueURL          string `json:"issue_url"`
	CommitsURL        string `json:"commits_url"`
	ReviewCommentsURL string `json:"review_comments_url"`
	ReviewCommentURL  string `json:"review_comment_url"`
	CommentsURL       string `json:"comments_url"`
	StatusesURL       string `json:"statuses_url"`
	Number            int    `json:"number"`
	State             string `json:"state"`
	Title             string `json:"title"`
	Body              string `json:"body"`
	Assignee          *User  `json:"assignee"`
	Milestone         struct {
		URL          string    `json:"url"`
		HTMLURL      string    `json:"html_url"`
		LabelsURL    string    `json:"labels_url"`
		ID           int       `json:"id"`
		Number       int       `json:"number"`
		State        string    `json:"state"`
		Title        string    `json:"title"`
		Description  string    `json:"description"`
		Creator      User      `json:"creator"`
		OpenIssues   int       `json:"open_issues"`
		ClosedIssues int       `json:"closed_issues"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		ClosedAt     time.Time `json:"closed_at"`
		DueOn        time.Time `json:"due_on"`
	} `json:"milestone"`
	Locked    bool      `json:"locked"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ClosedAt  time.Time `json:"closed_at"`
	MergedAt  time.Time `json:"merged_at"`
	Head      struct {
		Label string `json:"label"`
		Ref   string `json:"ref"`
		SHA   string `json:"sha"`
		User  User   `json:"user"`
		Repo  struct {
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
		} `json:"repo"`
	} `json:"head"`
	Base struct {
		Label string `json:"label"`
		Ref   string `json:"ref"`
		SHA   string `json:"sha"`
		User  User   `json:"user"`
		Repo  struct {
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
		} `json:"repo"`
	} `json:"base"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		Issue struct {
			Href string `json:"href"`
		} `json:"issue"`
		Comments struct {
			Href string `json:"href"`
		} `json:"comments"`
		ReviewComments struct {
			Href string `json:"href"`
		} `json:"review_comments"`
		ReviewComment struct {
			Href string `json:"href"`
		} `json:"review_comment"`
		Commits struct {
			Href string `json:"href"`
		} `json:"commits"`
		Statuses struct {
			Href string `json:"href"`
		} `json:"statuses"`
	} `json:"_links"`
	User User `json:"user"`
}

// PullRequestCommit contains commit information. This value is returned by PullRequestsAPI.GetCommits.
type PullRequestCommit struct {
	SHA    string `json:"sha"`
	Commit struct {
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
			SHA string `json:"sha"`
			URL string `json:"url"`
		} `json:"tree"`
		URL          string `json:"url"`
		CommentCount int    `json:"comment_count"`
	} `json:"commit"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	CommentsURL string `json:"comments_url"`
	Author      User   `json:"author"`
	Committer   User   `json:"committer"`
	Parents     []struct {
		SHA     string `json:"sha"`
		URL     string `json:"url"`
		HTMLURL string `json:"html_url"`
	} `json:"parents"`
}

// PullRequestReview contains a Review with a state and optional body comment on a PullRequest
type PullRequestReview struct {
	ID     int    `json:"id"`
	NodeID string `json:"node_id"`
	User   struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
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
	} `json:"user"`
	Body           string `json:"body"`
	CommitID       string `json:"commit_id"`
	State          string `json:"state"`
	HTMLURL        string `json:"html_url"`
	PullRequestURL string `json:"pull_request_url"`
	Links          struct {
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		PullRequest struct {
			Href string `json:"href"`
		} `json:"pull_request"`
	} `json:"_links"`
}

// PullRequestReviewComments contains the comments on a unified diff in a Pull Request
type PullRequestReviewComments struct {
	URL                 string `json:"url"`
	ID                  int    `json:"id"`
	NodeID              string `json:"node_id"`
	PullRequestReviewID int    `json:"pull_request_review_id"`
	DiffHunk            string `json:"diff_hunk"`
	Path                string `json:"path"`
	Position            int    `json:"position"`
	OriginalPosition    int    `json:"original_position"`
	CommitID            string `json:"commit_id"`
	OriginalCommitID    string `json:"original_commit_id"`
	InReplyToID         int    `json:"in_reply_to_id"`
	User                struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
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
	} `json:"user"`
	Body           string    `json:"body"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	HTMLURL        string    `json:"html_url"`
	PullRequestURL string    `json:"pull_request_url"`
	Links          struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		PullRequest struct {
			Href string `json:"href"`
		} `json:"pull_request"`
	} `json:"_links"`
}


// ListPullRequests lists pull requests for the repository.
// state is either "open", "closed", or "all".
func (api *PullRequestsAPI) ListPullRequests(state string) ([]PullRequestResponse, error) {
	var allPullRequests []PullRequestResponse
	for page := 1; ; page++ {
		url := api.getURL(fmt.Sprintf("/repos/:owner/:repo/pulls?state=%s&page=%d", state, page))

		resp, err := api.httpGet(url)
		//resp.Header["Link"] // TODO (judwhite), get next page until last
		//<url>; rel="last", <url>; rel="first", <url>; rel="prev", <url>; rel="next"
		if err != nil {
			return nil, err
		}

		pullRequests := []PullRequestResponse{}
		if err = json.NewDecoder(resp.Body).Decode(&pullRequests); err != nil {
			resp.Body.Close()
			return nil, err
		}
		resp.Body.Close()

		allPullRequests = append(allPullRequests, pullRequests...)
		if len(pullRequests) == 0 || resp.Header.Get("Link") == "" {
			break
		}
	}

	return allPullRequests, nil
}

// GetPullRequest get a pull request by PR number.
func (api *PullRequestsAPI) GetPullRequest(pullRequestNumber int) (*PullRequestResponse, error) {
	url := api.getURL("/repos/:owner/:repo/pulls/" + strconv.Itoa(pullRequestNumber))

	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pullRequest PullRequestResponse

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&pullRequest); err != nil {
		return nil, err
	}

	return &pullRequest, nil
}

// MergePullRequest merges a pull request using the specified method.
func (api *PullRequestsAPI) MergePullRequest(pullRequestNumber int, method MergeMethod) (*MergeRequestResponse, error) {
	url := api.getURL("/repos/:owner/:repo/pulls/" + strconv.Itoa(pullRequestNumber) + "/merge")
	body := struct {
		MergeMethod string `json:"merge_method"`
	}{string(method)}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := api.httpPut(url, string(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var mergeRequest MergeRequestResponse

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&mergeRequest); err != nil {
		return nil, err
	}
	return &mergeRequest, err
}

// GetCommits gets all commits for a pull request by PR number.
func (api *PullRequestsAPI) GetCommits(pullRequestNumber int) ([]PullRequestCommit, error) {
	var allCommits []PullRequestCommit
	for page := 1; ; page++ {
		url := api.getURL(fmt.Sprintf("/repos/:owner/:repo/pulls/%d/commits?page=%d", pullRequestNumber, page))

		resp, err := api.httpGet(url)
		//resp.Header["Link"] // TODO (judwhite), get next page until last
		//<url>; rel="last", <url>; rel="first", <url>; rel="prev", <url>; rel="next"
		if err != nil {
			return nil, err
		}

		commits := []PullRequestCommit{}
		if err = json.NewDecoder(resp.Body).Decode(&commits); err != nil {
			resp.Body.Close()
			return nil, err
		}
		resp.Body.Close()

		allCommits = append(allCommits, commits...)
		if len(commits) == 0 || resp.Header.Get("Link") == "" {
			break
		}
	}

	return allCommits, nil
}

// Create creates a Pull Request using head (owner:branch) targeting the base (target branch).
// See https://developer.github.com/v3/pulls/#create-a-pull-request
func (api *PullRequestsAPI) Create(head, base, title, body string, maintainerCanModify bool) (*CreatePullRequestResponse, error) {
	post := struct {
		Head                string `json:"head"`
		Base                string `json:"base"`
		Title               string `json:"title"`
		Body                string `json:"body"`
		MaintainerCanModify bool   `json:"maintainer_can_modify"`
	}{
		Head:                head,
		Base:                base,
		Title:               title,
		Body:                body,
		MaintainerCanModify: maintainerCanModify,
	}

	url := api.getURL("/repos/:owner/:repo/pulls")

	b, err := json.Marshal(post)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpPost(url, string(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pr CreatePullRequestResponse

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&pr); err != nil {
		return nil, err
	}

	return &pr, nil
}

// ListPullRequestReviews gets all the PullRequestReviews for a Pull Request
func (api *PullRequestsAPI) ListPullRequestReviews(pullRequestNumber int) ([]PullRequestReview, error) {
	var allPRReviews []PullRequestReview
	for page := 1; ; page++ {
		url := api.getURL(fmt.Sprintf("/repos/:owner/:repo/pulls/%d/reviews?page=%d", pullRequestNumber, page))

		resp, err := api.httpGet(url)
		//resp.Header["Link"] // TODO (judwhite), get next page until last
		//<url>; rel="last", <url>; rel="first", <url>; rel="prev", <url>; rel="next"
		if err != nil {
			return nil, err
		}

		pullRequestReviews := []PullRequestReview{}
		if err = json.NewDecoder(resp.Body).Decode(&pullRequestReviews); err != nil {
			resp.Body.Close()
			return nil, err
		}
		resp.Body.Close()

		allPRReviews = append(allPRReviews, pullRequestReviews...)
		if len(pullRequestReviews) == 0 || resp.Header.Get("Link") == "" {
			break
		}
	}
	return allPRReviews, nil
}

// ListPullRequestReviewComments gets all PullRequestReviewComments for a Pull Request
func (api *PullRequestsAPI) ListPullRequestReviewComments(pullRequestNumber int) ([]PullRequestReviewComments, error) {
	var allPRReviewComments []PullRequestReviewComments
	for page := 1; ; page++ {
		url := api.getURL(fmt.Sprintf("/repos/:owner/:repo/pulls/%d/comments?page=%d", pullRequestNumber, page))

		resp, err := api.httpGet(url)
		//resp.Header["Link"] // TODO (judwhite), get next page until last
		//<url>; rel="last", <url>; rel="first", <url>; rel="prev", <url>; rel="next"
		if err != nil {
			return nil, err
		}

		pullRequestReviewComments := []PullRequestReviewComments{}
		if err = json.NewDecoder(resp.Body).Decode(&pullRequestReviewComments); err != nil {
			resp.Body.Close()
			return nil, err
		}
		resp.Body.Close()

		allPRReviewComments = append(allPRReviewComments, pullRequestReviewComments...)
		if len(pullRequestReviewComments) == 0 || resp.Header.Get("Link") == "" {
			break
		}
	}
	return allPRReviewComments, nil
}