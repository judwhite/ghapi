package ghapi

import (
	"encoding/json"
	"fmt"
)

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

		teams := []ListTeamsResponse{}
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
func (api *OrganizationAPI) ListTeamMembers(teamID int, role string) ([]User, error) {
	var allTeamMembers []User
	for page := 1; ; page++ {
		url := api.addBaseURL(fmt.Sprintf("/teams/%d/members?role=%s&page=%d", teamID, role, page))

		resp, err := api.httpGet(url)
		//resp.Header["Link"] // TODO (judwhite), get next page until last
		//<url>; rel="last", <url>; rel="first", <url>; rel="prev", <url>; rel="next"
		if err != nil {
			return nil, err
		}

		teamMembers := []User{}
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
