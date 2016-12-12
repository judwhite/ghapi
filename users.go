package ghapi

import (
	"encoding/json"
	"fmt"
	"time"
)

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
	// LDAP Distinguished Name
	LDAPDN string `json:"ldap_dn"`
}

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
	Following   int       `json:"fllowing"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (api *UserAPI) GetUser(userName string) (*UserFull, error) {
	url := api.addBaseURL("/users/" + userName)
	return api.GetUserByURL(url)
}

func (api *UserAPI) GetUserByURL(url string) (*UserFull, error) {
	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	user := &UserFull{}

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetOrganizations gets public organization memberships for the specified user.
//
// This method only lists public memberships, regardless of authentication. If
// you need to fetch all of the organization memberships (public and private)
// for the authenticated user, use the List your organizations API instead.
func (api *UserAPI) GetOrganizations(userName string) ([]UserOrganizationPayload, error) {
	url := api.addBaseURL(fmt.Sprintf("/users/%s/orgs", userName))
	return api.GetOrganizationsByURL(url)
}

// GetOrganizationsByURL gets public organization memberships for the specified user.
//
// This method only lists public memberships, regardless of authentication. If
// you need to fetch all of the organization memberships (public and private)
// for the authenticated user, use the List your organizations API instead.
func (api *UserAPI) GetOrganizationsByURL(url string) ([]UserOrganizationPayload, error) {
	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	orgs := []UserOrganizationPayload{}

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&orgs); err != nil {
		return nil, err
	}

	return orgs, nil
}
