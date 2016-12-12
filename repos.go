package ghapi

import (
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
		time, err := time.Parse(ctLayout, string(b))
		if err != nil {
			return err
		}
		ct.Time = time
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
