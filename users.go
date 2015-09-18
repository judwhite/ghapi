package ghapi

import (
	"encoding/json"
	"fmt"
)

// The user who performed the event.
type UserPayload struct {
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
	LdapDn            string `json:"ldap_dn"`
}

func (api *UserApi) GetUser(userName string) (*UserPayload, error) {
	url := api.addBaseUrl("/users/" + userName)
	return api.GetUserByUrl(url)
}

func (api *UserApi) GetUserByUrl(url string) (*UserPayload, error) {
	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	user := &UserPayload{}

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (api *UserApi) GetOrganizations(userName string) ([]OrganizationPayload, error) {
	url := fmt.Sprintf("/users/%s/orgs", userName)
	return api.GetOrganizationsByUrl(url)
}

func (api *UserApi) GetOrganizationsByUrl(url string) ([]OrganizationPayload, error) {
	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	orgs := []OrganizationPayload{}

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&orgs); err != nil {
		return nil, err
	}

	return orgs, nil
}
