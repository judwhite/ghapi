package ghapi

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/url"
)

// Branch represents a branch in a repository. This value is returned by BranchesAPI.GetBranch.
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
				SHA string `json:"sha"`
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
			SHA string `json:"sha"`
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

// GetBranch gets a branch by name.
func (api *BranchesAPI) GetBranch(branch string) (*Branch, error) {
	url := api.getURL("/repos/:owner/:repo/branches/" + url.PathEscape(branch))

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

// BranchProtection specifies branch protection options.
type BranchProtection struct {
	// RequiredStatusChecks specifies which status checks must pass before branches can be merged.
	RequiredStatusChecks *RequiredStatusChecks `json:"required_status_checks"`
	// RequiredPullRequestReviews specifies how pull request reviews are handled. When enabled, all commits must be
	// made to a non-protected branch and submitted via a pull request with at least one approved review and no changes
	// requested before it can be merged.
	RequiredPullRequestReviews *RequiredPullRequestReviews `json:"required_pull_request_reviews"`
	// Restrictions specifies which users and teams can push to this branch.
	Restrictions *Restrictions `json:"restrictions"`
	// EnforceAdmins enforces all configured restrictions for administrators.
	EnforceAdmins bool `json:"enforce_admins"`
}

// RequiredStatusChecks specifies which status checks must pass before branches can be merged.
type RequiredStatusChecks struct {
	// Strict ensures the branch has been tested with the latest code on the target ref.
	Strict   bool     `json:"strict"`
	Contexts []string `json:"contexts"`
}

// RequiredPullRequestReviews specifies how pull request reviews are handled. When enabled, all commits must be
// made to a non-protected branch and submitted via a pull request with at least one approved review and no changes
// requested before it can be merged.
type RequiredPullRequestReviews struct {
	// Specify which users and teams can dismiss pull request reviews.
	DismissalRestrictions Restrictions `json:"dismissal_restrictions"`
	// Dismiss approved reviews automatically when a new commit is pushed.
	DismissStaleReviews bool `json:"dismiss_stale_reviews"`
	// Blocks merge until code owners have reviewed.
	RequireCodeOwnerReviews bool `json:"require_code_owner_reviews"`
}

// Restrictions contains a list of users and teams able to perform a specified action.
type Restrictions struct {
	// The list of user logins
	Users []string `json:"users"`
	// The list of team slugs
	Teams []string `json:"teams"`
}

// Protect enables branch protection for the specified branch.
func (api *BranchesAPI) Protect(branch string, opts BranchProtection) error {
	b, err := json.Marshal(opts)
	if err != nil {
		return err
	}

	apiURL := api.getURL("/repos/:owner/:repo/branches/" + url.PathEscape(branch) + "/protection")

	body := string(b)
	resp, err := api.doHTTPRequest("PUT", apiURL, &body, "application/vnd.github.loki-preview+json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(ioutil.Discard, resp.Body)
	return err
}
