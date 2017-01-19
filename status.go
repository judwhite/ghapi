package ghapi

import (
	"encoding/json"
	"fmt"
	"time"
)

// StatusState represents the state of a commit status (pending, success, error, failure).
type StatusState string

const (
	// Pending commit status state; "pending".
	Pending StatusState = "pending"
	// Success commit status state; "success".
	Success StatusState = "success"
	// Error commit status state; "error".
	Error StatusState = "error"
	// Failure commit status state; "failure".
	Failure StatusState = "failure"
)

// CommitStatus contains a commit status.
type CommitStatus struct {
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

// CommitCombinedStatus contains a combined view of commit statuses for a given ref.
type CommitCombinedStatus struct {
	State      StatusState `json:"state"`
	SHA        string      `json:"sha"`
	TotalCount int         `json:"total_count"`
	Statuses   []struct {
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		State       string    `json:"state"`
		TargetURL   string    `json:"target_url"`
		Description string    `json:"description"`
		ID          int       `json:"id"`
		URL         string    `json:"url"`
		Context     string    `json:"context"`
	} `json:"statuses"`
	Repository struct {
		ID          int    `json:"id"`
		Owner       User   `json:"owner"`
		Name        string `json:"name"`
		FullName    string `json:"full_name"`
		Description string `json:"description"`
		Private     bool   `json:"private"`
		Fork        bool   `json:"fork"`
		URL         string `json:"url"`
		HTMLURL     string `json:"html_url"`
	} `json:"repository"`
	CommitURL string `json:"commit_url"`
	URL       string `json:"url"`
}

// SetStatus creates a commit status for a given ref.
// state is either "pending", "success", "error", or "failure".
// targetURL is the target URL to associate with this status. This URL will be linked from the GitHub UI to allow users
// to easily see the 'source' of the Status.
// description is a short description of the status (often a more detailed description of the state).
// context is a string label to differentiate this status from the status of other systems/steps.
func (api *StatusAPI) SetStatus(sha string, state StatusState, targetURL, description, context string) (*CommitStatus, error) {
	url := api.getURL("/repos/:owner/:repo/statuses/") + sha

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

	var status CommitStatus

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&status); err != nil {
		return nil, err
	}

	return &status, nil
}

// GetList lists statuses for a specific Ref. The Ref can be a SHA, a branch name, or a tag name.
func (api *StatusAPI) GetList(ref string) ([]CommitStatus, error) {
	url := api.getURL(fmt.Sprintf("/repos/:owner/:repo/commits/%s/statuses", ref))

	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	statuses := []CommitStatus{}

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&statuses); err != nil {
		return nil, err
	}

	return statuses, nil
}

// GetCombined returns a combined view of commit statuses for a given ref. The Ref can be a SHA, a branch name, or
// a tag name. The returned state is either "success", "pending", or "failure" ("error" states become "failure").
func (api *StatusAPI) GetCombined(ref string) (*CommitCombinedStatus, error) {
	url := api.getURL(fmt.Sprintf("/repos/:owner/:repo/commits/%s/status", ref))

	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var status CommitCombinedStatus

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&status); err != nil {
		return nil, err
	}

	return &status, nil
}
