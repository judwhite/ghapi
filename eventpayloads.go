package ghapi

type GithubEventPayload struct {
	Repository RepositoryPayload `json:"repository"`
	Sender     User              `json:"sender"`
}

type PullRequestEventPayload struct {
	GithubEventPayload
	Action      PullRequestAction  `json:"action"`
	Number      int                `json:"number"`
	PullRequest PullRequestPayload `json:"pull_request"`
}

type IssueCommentEventPayload struct {
	GithubEventPayload
	Action  string              `json:"action"`
	Issue   IssuePayload        `json:"issue"`
	Comment IssueCommentPayload `json:"comment"`
}
