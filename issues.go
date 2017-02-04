package ghapi

import (
	"encoding/json"
	"strconv"
	"time"
)

// IssueLabel contains Issue label information. This type is used in IssueResponse.
type IssueLabel struct {
	ID      int    `json:"id"`
	URL     string `json:"url"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Default bool   `json:"default"`
}

// IssueResponse contains Issue information. This value is returned by IssuesAPI for get and edit API calls.
type IssueResponse struct {
	ID            int          `json:"id"`
	URL           string       `json:"url"`
	RepositoryURL string       `json:"repository_url"`
	LabelsURL     string       `json:"labels_url"`
	CommentsURL   string       `json:"comments_url"`
	EventsURL     string       `json:"events_url"`
	HTMLURL       string       `json:"html_url"`
	Number        int          `json:"number"`
	State         string       `json:"state"`
	Title         string       `json:"title"`
	Body          string       `json:"body"`
	User          User         `json:"user"`
	Labels        []IssueLabel `json:"labels"`
	Assignee      *User        `json:"assignee"`
	Milestone     struct {
		URL          string     `json:"url"`
		HTMLURL      string     `json:"html_url"`
		LabelsURL    string     `json:"labels_url"`
		ID           int        `json:"id"`
		Number       int        `json:"number"`
		State        string     `json:"state"`
		Title        string     `json:"title"`
		Description  string     `json:"description"`
		Creator      User       `json:"creator"`
		OpenIssues   int        `json:"open_issues"`
		ClosedIssues int        `json:"closed_issues"`
		CreatedAt    time.Time  `json:"created_at"`
		UpdatedAt    time.Time  `json:"updated_at"`
		ClosedAt     *time.Time `json:"closed_at"`
		DueOn        *time.Time `json:"due_on"`
	} `json:"milestone"`
	Locked      bool `json:"locked"`
	Comments    int  `json:"comments"`
	PullRequest struct {
		URL      string `json:"url"`
		HTMLURL  string `json:"html_url"`
		DiffURL  string `json:"diff_url"`
		PatchURL string `json:"patch_url"`
	} `json:"pull_request"`
	ClosedAt  *time.Time `json:"closed_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	ClosedBy  *User      `json:"closed_by"`
}

// IssueCommentResponse returns information about a specific comment on an issue.
type IssueCommentResponse struct {
	ID        int       `json:"id"`
	URL       string    `json:"url"`
	HTMLURL   string    `json:"html_url"`
	Body      string    `json:"body"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DeleteIssueComment deletes an issue comment by ID.
func (api *IssueAPI) DeleteIssueComment(commentID int) error {
	url := api.getURL("/repos/:owner/:repo/issues/comments/" + strconv.Itoa(commentID))
	return api.DeleteIssueCommentByURL(url)
}

// DeleteIssueCommentByURL deletes an issue comment by URL.
func (api *IssueAPI) DeleteIssueCommentByURL(url string) error {
	resp, err := api.httpDelete(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// GetIssueComment gets an issue comment by ID.
func (api *IssueAPI) GetIssueComment(commentID int) (*IssueCommentResponse, error) {
	url := api.getURL("/repos/:owner/:repo/issues/comments/" + strconv.Itoa(commentID))

	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issueComment IssueCommentResponse

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&issueComment); err != nil {
		return nil, err
	}

	return &issueComment, nil
}

// GetIssue gets an issue by issue number.
func (api *IssueAPI) GetIssue(issueNumber int) (*IssueResponse, error) {
	url := api.getURL("/repos/:owner/:repo/issues/" + strconv.Itoa(issueNumber))
	return api.GetIssueByURL(url)
}

// GetIssueByURL gets an issue by URL.
func (api *IssueAPI) GetIssueByURL(url string) (*IssueResponse, error) {
	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issue IssueResponse

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

// UpdateIssueAssignee updates an issue's assignee by issue number.
func (api *IssueAPI) UpdateIssueAssignee(issueNumber int, assignee string) (*IssueResponse, error) {
	url := api.getURL("/repos/:owner/:repo/issues/" + strconv.Itoa(issueNumber))
	return api.UpdateIssueAssigneeByURL(url, assignee)
}

// UpdateIssueAssigneeByURL updates an issue's assignee by issue URL.
func (api *IssueAPI) UpdateIssueAssigneeByURL(url, assignee string) (*IssueResponse, error) {
	// TODO (judwhite): there can be multiple assignees
	body := struct {
		Assignee string `json:"assignee"`
	}{assignee}

	return api.updateIssueByURL(url, body)
}

// UpdateIssueLabels updates an issue's labels by issue number. The labels passed become the new
// labels. See AddLabel and RemoveLabel to add/remove individual labels.
func (api *IssueAPI) UpdateIssueLabels(issueNumber int, labels []string) (*IssueResponse, error) {
	url := api.getURL("/repos/:owner/:repo/issues/" + strconv.Itoa(issueNumber))
	return api.UpdateIssueLabelsByURL(url, labels)
}

// UpdateIssueLabelsByURL updates an issue's labels by issue URL. The labels passed become the new
// labels. See AddLabel and RemoveLabel to add/remove individual labels.
func (api *IssueAPI) UpdateIssueLabelsByURL(url string, labels []string) (*IssueResponse, error) {
	body := struct {
		Labels []string `json:"labels"`
	}{labels}

	return api.updateIssueByURL(url, body)
}

// AddLabel adds a label to an issue.
func (api *IssueAPI) AddLabel(issueNumber int, labelName string) error {
	issue, err := api.GetIssue(issueNumber)
	if err != nil {
		return err
	}
	var labels []string
	for _, existingLabel := range issue.Labels {
		if existingLabel.Name == labelName {
			return nil
		}
		labels = append(labels, existingLabel.Name)
	}

	labels = append(labels, labelName)
	_, err = api.UpdateIssueLabels(issueNumber, labels)
	return err
}

// RemoveLabel remove a label from an issue.
func (api *IssueAPI) RemoveLabel(issueNumber int, labelName string) error {
	issue, err := api.GetIssue(issueNumber)
	if err != nil {
		return err
	}
	labels := []string{}
	var hasLabel bool
	for _, existingLabel := range issue.Labels {
		if existingLabel.Name == labelName {
			hasLabel = true
		} else {
			labels = append(labels, existingLabel.Name)
		}
	}

	if !hasLabel {
		return nil
	}

	_, err = api.UpdateIssueLabels(issueNumber, labels)
	return err
}

func (api *IssueAPI) updateIssueByURL(url string, body interface{}) (*IssueResponse, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpPatch(url, string(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issue IssueResponse

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&issue); err != nil {
		return nil, err
	}

	return &issue, nil
}
