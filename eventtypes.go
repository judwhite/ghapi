package ghapi

type GitHubEventType string

// Refer to https://developer.github.com/webhooks/#events
const (
	// Any time a Commit is commented on.
	CommitCommentEventType GitHubEventType = "commit_comment"
	// Any time a Branch or Tag is created.
	CreateEventType GitHubEventType = "create"
	// Any time a Branch or Tag is deleted.
	DeleteEventType GitHubEventType = "delete"
	// Any time a Repository has a new deployment created from the API.
	DeploymentEventType GitHubEventType = "deployment"
	// Any time a deployment for a Repository has a status update from the API.
	DeploymentStatusEventType GitHubEventType = "deployment_status"
	// Any time a Repository is forked.
	ForkEventType GitHubEventType = "fork"
	// Any time a Wiki page is updated.
	GollumEventType GitHubEventType = "gollum"
	// Any time an Issue or Pull Request is commented on.
	IssueCommentEventType GitHubEventType = "issue_comment"
	// Any time an Issue is assigned, unassigned, labeled, unlabeled, opened,
	// closed, or reopened.
	IssuesEventType GitHubEventType = "issues"
	// Any time a User is added as a collaborator to a non-Organization
	// Repository.
	MemberEventType GitHubEventType = "member"
	// Any time a User is added or removed from a team. Organization hooks only.
	MembershipEventType GitHubEventType = "membership"
	// Any time a Pages site is built or results in a failed build.
	PageBuildEventType GitHubEventType = "page_build"
	// Any time a Repository changes from private to public.
	PublicEventType GitHubEventType = "public"
	// Any time a comment is created on a portion of the unified diff of a pull
	// request (the Files Changed tab).
	PullRequestReviewCommentEventType GitHubEventType = "pull_request_review_comment"
	// Any time a Pull Request is assigned, unassigned, labeled, unlabeled,
	// opened, closed, reopened, or synchronized (updated due to a new push in
	// the branch that the pull request is tracking).
	PullRequestEventType GitHubEventType = "pull_request"
	// Any Git push to a Repository, including editing tags or branches.
	// Commits via API actions that update references are also counted.
	PushEventType GitHubEventType = "push"
	// Any time a Repository is created. Organization hooks only.
	RepositoryEventType GitHubEventType = "repository"
	// Any time a Release is published in a Repository.
	ReleaseEventType GitHubEventType = "release"
	// Any time a Repository has a status update from the API.
	StatusEventType GitHubEventType = "status"
	// Any time a team is added or modified on a Repository.
	TeamAddEventType GitHubEventType = "team_add"
	// Any time a User watches a Repository.
	WatchEventType GitHubEventType = "watch"
	// When you create a new webhook, we’ll send you a simple ping event to let
	// you know you’ve set up the webhook correctly. This event isn’t stored so
	// it isn’t retrievable via the Events API. You can trigger a ping again by
	// calling the ping endpoint.
	PingEventType GitHubEventType = "ping"
)
