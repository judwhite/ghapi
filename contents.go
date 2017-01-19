package ghapi

import "encoding/json"

// Contents is returned by ContentsAPI.GetContent. The Content field is Base64 encoded.
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

// GetContent gets the content for the specified path from the default branch. The Contents.Content field
// will be Base64 encoded. This API supports files up to 1MB in size.
// TODO (judwhite): API can get contents from others refs; see https://developer.github.com/v3/repos/contents/#get-contents
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
