package ghapi

import (
	"encoding/json"
	"fmt"
	"time"
)

// User contains information about a GitHub user; it's used as a field in many other responses. For full user
// information, see UserFull.
type User struct {
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
	// LDAPDN is the LDAP Distinguished Name (GitHub Enterprise)
	LDAPDN string `json:"ldap_dn"`
}

// UserFull contains additional information about a user beyong the User struct. It's the response type of
// UserAPI.GetUser and UserAPI.GetUserByURL.
type UserFull struct {
	User
	Name        string    `json:"name"`
	Company     string    `json:"company"`
	Blog        string    `json:"blog"`
	Location    string    `json:"location"`
	Email       string    `json:"email"`
	Hireable    bool      `json:"hireable"`
	Bio         string    `json:"bio"`
	PublicRepos int       `json:"public_repos"`
	PublicGists int       `json:"public_gists"`
	Followers   int       `json:"followers"`
	Following   int       `json:"following"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// UserPublicOrganizationResponse contains information about an organization a user belongs to, which is publicly visible.
type UserPublicOrganizationResponse struct {
	Login            string `json:"login"`
	ID               int    `json:"id"`
	URL              string `json:"url"`
	ReposURL         string `json:"repos_url"`
	EventsURL        string `json:"events_url"`
	HooksURL         string `json:"hooks_url"`
	IssuesURL        string `json:"issues_url"`
	MembersURL       string `json:"members_url"`
	PublicMembersURL string `json:"public_members_url"`
	AvatarURL        string `json:"avatar_url"`
	Description      string `json:"description"`
}

// GetUser returns user information by login name.
func (api *UserAPI) GetUser(userName string) (*UserFull, error) {
	url := api.addBaseURL("/users/" + userName)
	return api.GetUserByURL(url)
}

// GetUserByURL returns user information by URL.
func (api *UserAPI) GetUserByURL(url string) (*UserFull, error) {
	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user UserFull

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetPublicOrganizations gets public organization memberships for the specified user.
//
// This method only lists public memberships, regardless of authentication. If
// you need to fetch all of the organization memberships (public and private)
// for the authenticated user, use the List your organizations API instead.
func (api *UserAPI) GetPublicOrganizations(userName string) ([]UserPublicOrganizationResponse, error) {
	url := api.addBaseURL(fmt.Sprintf("/users/%s/orgs", userName))
	return api.GetPublicOrganizationsByURL(url)
}

// GetPublicOrganizationsByURL gets public organization memberships for the specified user.
//
// This method only lists public memberships, regardless of authentication. If
// you need to fetch all of the organization memberships (public and private)
// for the authenticated user, use the List your organizations API instead.
func (api *UserAPI) GetPublicOrganizationsByURL(url string) ([]UserPublicOrganizationResponse, error) {
	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	orgs := []UserPublicOrganizationResponse{}

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&orgs); err != nil {
		return nil, err
	}

	return orgs, nil
}
