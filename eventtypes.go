package ghapi

// GitHubEventType represents the value of the "X-Github-Event" header from a GitHub event.
// See https://developer.github.com/webhooks/#events.
type GitHubEventType string

const (
	// CommitCommentEventType occurs when a Commit is commented on.
	CommitCommentEventType GitHubEventType = "commit_comment"
	// CreateEventType occurs when a Branch or Tag is created.
	CreateEventType GitHubEventType = "create"
	// DeleteEventType occurs when a Branch or Tag is deleted.
	DeleteEventType GitHubEventType = "delete"
	// DeploymentEventType occurs when a Repository has a new deployment created from the API.
	DeploymentEventType GitHubEventType = "deployment"
	// DeploymentStatusEventType occurs when a deployment for a Repository has a status update from the API.
	DeploymentStatusEventType GitHubEventType = "deployment_status"
	// ForkEventType occurs when a Repository is forked.
	ForkEventType GitHubEventType = "fork"
	// GollumEventType occurs when a Wiki page is updated.
	GollumEventType GitHubEventType = "gollum"
	// IssueCommentEventType occurs when an Issue or Pull Request is commented on.
	IssueCommentEventType GitHubEventType = "issue_comment"
	// IssuesEventType occurs when an Issue is assigned, unassigned, labeled, unlabeled, opened,
	// closed, or reopened.
	IssuesEventType GitHubEventType = "issues"
	// LabelEventType occurs when a label is created, edited, or deleted.
	LabelEventType GitHubEventType = "label"
	// MemberEventType occurs when a User is added or removed as a collaborator to a non-Organization Repository.
	MemberEventType GitHubEventType = "member"
	// MembershipEventType occurs when a User is added or removed from a team. Organization hooks only.
	MembershipEventType GitHubEventType = "membership"
	// MilestoneEventType occurs when a Milestone is created, closed, opened, edited, or deleted.
	MilestoneEventType GitHubEventType = "milestone"
	// PageBuildEventType occurs when a Pages site is built or results in a failed build.
	PageBuildEventType GitHubEventType = "page_build"
	// PublicEventType occurs when a Repository changes from private to public.
	PublicEventType GitHubEventType = "public"
	// PullRequestReviewCommentEventType occurs when a comment is created, edited, or deleted (in the Files
	// Changed tab).
	PullRequestReviewCommentEventType GitHubEventType = "pull_request_review_comment"
	// PullRequestReviewEventType occurs when a Pull Request Review is submitted.
	PullRequestReviewEventType GitHubEventType = "pull_request_review"
	// PullRequestEventType occurs when a Pull Request is assigned, unassigned, labeled, unlabeled, opened, closed,
	// reopened, or synchronized (updated due to a new push inthe branch that the pull request is tracking).
	PullRequestEventType GitHubEventType = "pull_request"
	// PushEventType occurs when any Git push to a Repository occurs, including editing tags or branches.
	// Commits via API actions that update references are also counted.
	PushEventType GitHubEventType = "push"
	// RepositoryEventType occurs when a Repository is created. Organization hooks only.
	RepositoryEventType GitHubEventType = "repository"
	// ReleaseEventType occurs when a Release is published in a Repository.
	ReleaseEventType GitHubEventType = "release"
	// StatusEventType occurs when a Repository has a status update from the API.
	StatusEventType GitHubEventType = "status"
	// TeamEventType occurs when a team is created, deleted, modified, or added to or removed from a repository.
	// Organization hooks only.
	TeamEventType GitHubEventType = "team"
	// TeamAddEventType occurs when a team is added or modified on a Repository.
	TeamAddEventType GitHubEventType = "team_add"
	// WatchEventType occurs any time a User watches a Repository.
	WatchEventType GitHubEventType = "watch"
	// PingEventType occurs when you create a new webhook. GitHub will send you a simple ping event to letyou know
	// you've set up the webhook correctly. This event isn't stored so it isn't retrievable via the Events API. You
	// can trigger a ping again by calling the ping endpoint. See https://developer.github.com/webhooks/#ping-event
	// and https://developer.github.com/v3/repos/hooks/#ping-a-hook.
	PingEventType GitHubEventType = "ping"
)
