package ghapi

import "encoding/json"

type Branch struct {
	Name   string `json:"name"`
	Commit struct {
		SHA    string `json:"sha"`
		Commit struct {
			Author struct {
				Name  string `json:"name"`
				Date  string `json:"date"`
				Email string `json:"email"`
			} `json:"author"`
			URL     string `json:"url"`
			Message string `json:"message"`
			Tree    struct {
				Sha string `json:"sha"`
				URL string `json:"url"`
			} `json:"tree"`
			Committer struct {
				Name  string `json:"name"`
				Date  string `json:"date"`
				Email string `json:"email"`
			} `json:"committer"`
		} `json:"commit"`
		Author struct {
			GravatarID string `json:"gravatar_id"`
			AvatarURL  string `json:"avatar_url"`
			URL        string `json:"url"`
			ID         int    `json:"id"`
			Login      string `json:"login"`
		} `json:"author"`
		Parents []struct {
			Sha string `json:"sha"`
			URL string `json:"url"`
		} `json:"parents"`
		URL       string `json:"url"`
		Committer struct {
			GravatarID string `json:"gravatar_id"`
			AvatarURL  string `json:"avatar_url"`
			URL        string `json:"url"`
			ID         int    `json:"id"`
			Login      string `json:"login"`
		} `json:"committer"`
	} `json:"commit"`
	Links struct {
		HTML string `json:"html"`
		Self string `json:"self"`
	} `json:"_links"`
	Protected     bool   `json:"protected"`
	ProtectionURL string `json:"protection_url"`
}

func (api *BranchesAPI) GetBranch(branch string) (*Branch, error) {
	url := api.getURL("/repos/:owner/:repo/branches/" + branch)

	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &Branch{}

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}
