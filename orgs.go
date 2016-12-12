package ghapi

import "time"

// OrganizationPayload contains information about the Organization
type OrganizationPayload struct {
	UserOrganizationPayload
	ReposURL          string                   `json:"repos_url"`
	EventsURL         string                   `json:"events_url"`
	MembersURL        string                   `json:"members_url"`
	PublicMembersURL  string                   `json:"public_members_url"`
	PublicRepos       int                      `json:"public_repos"`
	PublicGists       int                      `json:"public_gists"`
	Followers         int                      `json:"followers"`
	Following         int                      `json:"following"`
	HTMLURL           string                   `json:"html_url"`
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

type UserOrganizationPayload struct {
	Login       string `json:"login"`
	ID          int    `json:"id"`
	URL         string `json:"url"`
	AvatarURL   string `json:"avatar_url"`
	Description string `json:"description"`
}

type OrganizationPlanPayload struct {
	Name         string `json:"name"`
	Space        int64  `json:"space"`
	PrivateRepos int64  `json:"private_repos"`
}
