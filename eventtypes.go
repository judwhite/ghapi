package ghapi

type GitHubEventType string

// Refer to https://developer.github.com/webhooks/#events
const (
	// Any time a Commit is commented on.
	CommitComment GitHubEventType = "commit_comment"
	// Any time a Branch or Tag is created.
	Create = "create"
	// Any time a Branch or Tag is deleted.
	Delete = "delete"
	// Any time a Repository has a new deployment created from the API.
	Deployment = "deployment"
	// Any time a deployment for a Repository has a status update from the API.
	DeploymentStatus = "deployment_status"
	// Any time a Repository is forked.
	Fork = "fork"
	// Any time a Wiki page is updated.
	Gollum = "gollum"
	// Any time an Issue or Pull Request is commented on.
	IssueComment = "issue_comment"
	// Any time an Issue is assigned, unassigned, labeled, unlabeled, opened,
	// closed, or reopened.
	Issues = "issues"
	// Any time a User is added as a collaborator to a non-Organization
	// Repository.
	Member = "member"
	// Any time a User is added or removed from a team. Organization hooks only.
	Membership = "membership"
	// Any time a Pages site is built or results in a failed build.
	PageBuild = "page_build"
	// Any time a Repository changes from private to public.
	Public = "public"
	// Any time a comment is created on a portion of the unified diff of a pull
	// request (the Files Changed tab).
	PullRequestReviewComment = "pull_request_review_comment"
	// Any time a Pull Request is assigned, unassigned, labeled, unlabeled,
	// opened, closed, reopened, or synchronized (updated due to a new push in
	// the branch that the pull request is tracking).
	PullRequest = "pull_request"
	// Any Git push to a Repository, including editing tags or branches.
	// Commits via API actions that update references are also counted.
	Push = "push"
	// Any time a Repository is created. Organization hooks only.
	Repository = "repository"
	// Any time a Release is published in a Repository.
	Release = "release"
	// Any time a Repository has a status update from the API.
	Status = "status"
	// Any time a team is added or modified on a Repository.
	TeamAdd = "team_add"
	// Any time a User watches a Repository.
	Watch = "watch"
	// When you create a new webhook, we’ll send you a simple ping event to let
	// you know you’ve set up the webhook correctly. This event isn’t stored so
	// it isn’t retrievable via the Events API. You can trigger a ping again by
	// calling the ping endpoint.
	Ping = "ping"
)
