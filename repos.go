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

// The repository which the event occurred on.
type RepositoryPayload struct {
	Id int `json:"id"`
	// The short name of the repository.
	Name string `json:"name"`
	// The name of the repository including the owner (user/organization).
	FullName string `json:"full_name"`
	// TODO: What if the owner is an organization? Owner.Type field?
	Owner            User       `json:"owner"`
	Private          bool       `json:"private"`
	HtmlUrl          string     `json:"html_url"`
	Fork             bool       `json:"fork"`
	Url              string     `json:"url"`
	ForksUrl         string     `json:"forks_url"`
	KeysUrl          string     `json:"keys"`
	CollaboratorsUrl string     `json:"collaborators_url"`
	TeamsUrl         string     `json:"teams_url"`
	HooksUrl         string     `json:"hooks_url"`
	IssueEventsUrl   string     `json:"issue_events_url"`
	EventUrl         string     `json:"events_url"`
	AssigneesUrl     string     `json:"assignees_url"`
	BranchesUrl      string     `json:"branches_url"`
	TagsUrl          string     `json:"tags_url"`
	BlobsUrl         string     `json:"blobs_url"`
	GitTagsUrl       string     `json:"git_tags_url"`
	GitRefsUrl       string     `json:"git_refs_url"`
	TreesUrl         string     `json:"trees_url"`
	StatusesUrl      string     `json:"statuses_url"`
	LanguagesUrl     string     `json:"languages_url"`
	StargazersUrl    string     `json:"stargazers_url"`
	ContributorsUrl  string     `json:"contributors_url"`
	SubscribersUrl   string     `json:"subscribers_url"`
	SubscriptionUrl  string     `json:"subscription_url"`
	CommitsUrl       string     `json:"commits_url"`
	GitCommitsUrl    string     `json:"git_commits_url"`
	CommentsUrl      string     `json:"comments_url"`
	IssueCommentUrl  string     `json:"issue_comment_url"`
	ContentsUrl      string     `json:"contents_url"`
	CompareUrl       string     `json:"compare_url"`
	MergesUrl        string     `json:"merges_url"`
	ArchiveUrl       string     `json:"archive_url"`
	DownloadsUrl     string     `json:"downloads_url"`
	IssuesUrl        string     `json:"issues_url"`
	PullsUrl         string     `json:"pulls_url"`
	MilestonesUrl    string     `json:"milestones_url"`
	NotificationsUrl string     `json:"notifications_url"`
	LabelsUrl        string     `json:"labels_url"`
	ReleasesUrl      string     `json:"releases_url"`
	CreatedAt        CustomTime `json:"created_at"`
	UpdatedAt        CustomTime `json:"updated_at"`
	PushedAt         CustomTime `json:"pushed_at"`
	GitUrl           string     `json:"git_url"`
	SshUrl           string     `json:"ssh_url"`
	CloneUrl         string     `json:"clone_url"`
	SvnUrl           string     `json:"svn_url"`
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
	MirrorUrl       string `json:"mirror_url"`
	OpenIssuesCount int    `json:"open_issues_count"`
	Forks           int    `json:"forks"`
	OpenIssues      int    `json:"open_issues"`
	Watchers        int    `json:"watchers"`
	DefaultBranch   string `json:"default_branch"`
}
