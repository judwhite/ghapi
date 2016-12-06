package ghapi

import (
	"encoding/json"
	"fmt"
	"time"
)

type StatusState string

const (
	Pending StatusState = "pending"
	Success StatusState = "success"
	Error   StatusState = "error"
	Failure StatusState = "failure"
)

type StatusPayload struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	State       string    `json:"state"`
	TargetURL   string    `json:"target_url"`
	Description string    `json:"description"`
	ID          int       `json:"id"`
	URL         string    `json:"url"`
	Context     string    `json:"context"`
	Creator     User      `json:"creator"`
}

type CombinedStatusPayload struct {
	State      StatusState       `json:"state"`
	SHA        string            `json:"sha"`
	TotalCount int               `json:"total_count"`
	Repository RepositoryPayload `json:"repository"`
	Statuses   []StatusPayload   `json:"statuses"`
	CommitURL  string            `json:"commit_url"`
	URL        string            `json:"url"`
}

func (api *StatusApi) SetStatus(sha string, state StatusState, targetURL, description, context string) (*StatusPayload, error) {
	url := api.getUrl("/repos/:owner/:repo/statuses/") + sha

	body := struct {
		State       string `json:"state"`
		TargetURL   string `json:"target_url"`
		Description string `json:"description"`
		Context     string `json:"context"`
	}{
		State:       string(state),
		TargetURL:   targetURL,
		Description: description,
		Context:     context,
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

	var status StatusPayload

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&status); err != nil {
		return nil, err
	}

	return &status, nil
}

func (api *StatusApi) GetList(ref string) ([]StatusPayload, error) {
	url := api.getUrl(fmt.Sprintf("/repos/:owner/:repo/commits/%s/statuses", ref))

	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var statuses []StatusPayload

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&statuses); err != nil {
		return nil, err
	}

	return statuses, nil
}

func (api *StatusApi) GetCombined(ref string) (*CombinedStatusPayload, error) {
	url := api.getUrl(fmt.Sprintf("/repos/:owner/:repo/commits/%s/status", ref))

	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var status CombinedStatusPayload

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&status); err != nil {
		return nil, err
	}

	return &status, nil
}
