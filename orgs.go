package ghapi

import (
	"encoding/json"
	"fmt"
	"time"
)

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

// ListTeamsResponse is the response from OrganizationAPI.ListTeams.
type ListTeamsResponse struct {
	ID              int    `json:"id"`
	URL             string `json:"url"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Description     string `json:"description"`
	Privacy         string `json:"privacy"`
	Permission      string `json:"permission"`
	MembersURL      string `json:"members_url"`
	RepositoriesURL string `json:"repositories_url"`
}

// ListTeamMembersResponse is the response from OrganizationAPI.ListTeamMembers.
type ListTeamMembersResponse struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
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
}

// ListTeams lists and organization's teams. Note: to use this API call your authtoken must have org:read permission.
func (api *OrganizationAPI) ListTeams() ([]ListTeamsResponse, error) {
	var allTeams []ListTeamsResponse
	for page := 1; ; page++ {
		url := api.addBaseURL(fmt.Sprintf("/orgs/%s/teams?page=%d", api.Organization, page))

		resp, err := api.httpGet(url)
		//resp.Header["Link"] // TODO (judwhite), get next page until last
		//<url>; rel="last", <url>; rel="first", <url>; rel="prev", <url>; rel="next"
		if err != nil {
			switch val := err.(type) {
			case *ErrHTTPError:
				if val.StatusCode == 403 {
					val.Message = "does your authtoken have org:read permission?"
					return nil, val
				}
			}
			return nil, err
		}

		teams := make([]ListTeamsResponse, 0)
		if err = json.NewDecoder(resp.Body).Decode(&teams); err != nil {
			resp.Body.Close()
			return nil, err
		}
		resp.Body.Close()

		allTeams = append(allTeams, teams...)
		if len(teams) == 0 || resp.Header.Get("Link") == "" {
			break
		}
	}

	return allTeams, nil
}

// ListTeamMembers list team members for the specified teamID.
// role is "member" (normal members of the team), "maintainer" (team maintainers), or "all".
func (api *OrganizationAPI) ListTeamMembers(teamID int, role string) ([]ListTeamMembersResponse, error) {
	var allTeamMembers []ListTeamMembersResponse
	for page := 1; ; page++ {
		url := api.addBaseURL(fmt.Sprintf("/teams/%d/members?role=%s&page=%d", teamID, role, page))

		resp, err := api.httpGet(url)
		//resp.Header["Link"] // TODO (judwhite), get next page until last
		//<url>; rel="last", <url>; rel="first", <url>; rel="prev", <url>; rel="next"
		if err != nil {
			return nil, err
		}

		teamMembers := make([]ListTeamMembersResponse, 0)
		if err = json.NewDecoder(resp.Body).Decode(&teamMembers); err != nil {
			resp.Body.Close()
			return nil, err
		}
		resp.Body.Close()

		allTeamMembers = append(allTeamMembers, teamMembers...)
		if len(teamMembers) == 0 || resp.Header.Get("Link") == "" {
			break
		}
	}

	return allTeamMembers, nil
}
