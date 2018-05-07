package ghapi

import "encoding/json"

// CreateRefResponse is returned by RefsAPI.Create.
type CreateRefResponse struct {
	Ref    string `json:"ref"`
	URL    string `json:"url"`
	Object struct {
		Type string `json:"type"`
		SHA  string `json:"sha"`
		URL  string `json:"url"`
	} `json:"object"`
}

// Create creates a reference in a repository from the specified SHA. 'ref' is the name of the fully qualified reference
// (ie: refs/heads/master). If it doesn't start with 'refs' and have at least two slashes, it will be rejected.
func (api *RefsAPI) Create(ref, sha string) (*CreateRefResponse, error) {
	url := api.getURL("/repos/:owner/:repo/git/refs")

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

// Get information about a ref. `ref` must be formatted as heads/master, not just master
// https://developer.github.com/v3/git/refs/#get-a-reference
func (api *RefsAPI) Get(ref string) (*CreateRefResponse, error) {
	url := api.getURL("/repos/:owner/:repo/git/refs/")

	resp, err := api.httpGet(url + ref)
	if err != nil {
		return nil, err
	}

	var refInfo CreateRefResponse

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&refInfo); err != nil {
		return nil, err
	}

	return &refInfo, nil
}