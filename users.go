package ghapi

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	// LDAP Distinguished Name
	LdapDn string `json:"ldap_dn"`
}

// The user who performed the event.
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

func (api *UserApi) GetUser(userName string) (*UserFull, error) {
	url := api.addBaseUrl("/users/" + userName)
	return api.GetUserByUrl(url)
}

func (api *UserApi) GetUserByUrl(url string) (*UserFull, error) {
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

// List public organization memberships for the specified user.
//
// This method only lists public memberships, regardless of authentication. If
// you need to fetch all of the organization memberships (public and private)
// for the authenticated user, use the List your organizations API instead.
func (api *UserApi) GetOrganizations(userName string) ([]UserOrganizationPayload, error) {
	url := api.addBaseUrl(fmt.Sprintf("/users/%s/orgs", userName))
	return api.GetOrganizationsByUrl(url)
}

// List public organization memberships for the specified user.
//
// This method only lists public memberships, regardless of authentication. If
// you need to fetch all of the organization memberships (public and private)
// for the authenticated user, use the List your organizations API instead.
func (api *UserApi) GetOrganizationsByUrl(url string) ([]UserOrganizationPayload, error) {
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
