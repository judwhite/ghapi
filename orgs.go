package ghapi

import "time"

// The organization which the event occurred on.
type OrganizationPayload struct {
	Login             string                   `json:"login"`
	Id                int                      `json:"id"`
	Url               string                   `json:"url"`
	ReposUrl          string                   `json:"repos_url"`
	EventsUrl         string                   `json:"events_url"`
	MembersUrl        string                   `json:"members_url"`
	PublicMembersUrl  string                   `json:"public_members_url"`
	AvatarUrl         string                   `json:"avatar_url"`
	Description       string                   `json:"description"`
	PublicRepos       int                      `json:"public_repos"`
	PublicGists       int                      `json:"public_gists"`
	Followers         int                      `json:"followers"`
	Following         int                      `json:"following"`
	HtmlUrl           string                   `json:"html_url"`
	CreatedAt         time.Time                `json:"created_at"`
	UpdatedAt         time.Time                `json:"updated_at"`
	Type              string                   `json:"type"`
	TotalPrivateRepos int                      `json:"total_private_repos"`
	OwnedPrivateRepos int                      `json:"owned_private_repos"`
	PrivateGist       int                      `json:"private_gists"`
	DiskUsage         int64                    `json:"disk_usage"`
	Collaborators     int                      `json:"collaborators"`
	BillingEmail      string                   `json:"billing_email"`
	Plan              *OrganizationPlanPayload `json:"plan"`
}

type OrganizationPlanPayload struct {
	Name         string `json:"name"`
	Space        int64  `json:"space"`
	PrivateRepos int64  `json:"private_repos"`
}
