package ghapi

import "encoding/json"

//CreateRefResponse is returned by RefsAPI.CreateRef.
type CreateRefResponse struct {
	Ref    string `json:"ref"`
	URL    string `json:"url"`
	Object struct {
		Type string `json:"type"`
		SHA  string `json:"sha"`
		URL  string `json:"url"`
	} `json:"object"`
}

//CreateRef creates a reference in a repository
func (api *RefsAPI) CreateRef(ref, sha string) (*CreateRefResponse, error) {
	url := api.getURL("/repos/:owner/:repo/git/refs/")

	body := struct {
		Ref string `json:"ref"`
		SHA string `json:"sha"`
	}{
		Ref: ref,
		SHA: sha,
	}

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpPost(url, string(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response CreateRefResponse

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
