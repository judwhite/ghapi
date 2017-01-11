package ghapi

import "encoding/json"

type Contents struct {
	Type        string `json:"type"`
	Encoding    string `json:"encoding"`
	Size        int    `json:"size"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Content     string `json:"content"`
	SHA         string `json:"sha"`
	URL         string `json:"url"`
	GitURL      string `json:"git_url"`
	HTMLURL     string `json:"html_url"`
	DownloadURL string `json:"download_url"`
	Links       struct {
		Git  string `json:"git"`
		Self string `json:"self"`
		HTML string `json:"html"`
	} `json:"_links"`
}

func (api *ContentsAPI) GetContent(path string) (*Contents, error) {
	url := api.getURL("/repos/:owner/:repo/contents/") + path

	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var contents Contents

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&contents); err != nil {
		return nil, err
	}

	return &contents, nil
}
