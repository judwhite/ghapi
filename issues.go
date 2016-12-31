package ghapi

import (
	"encoding/json"
	"strconv"
	"time"
)

type IssuePayload struct {
	URL         string                   `json:"url"`
	LabelsURL   string                   `json:"labels_url"`
	CommentsURL string                   `json:"comments_url"`
	EventsURL   string                   `json:"events_url"`
	HTMLURL     string                   `json:"html_url"`
	ID          int                      `json:"id"`
	Number      int                      `json:"number"`
	Title       string                   `json:"title"`
	User        User                     `json:"user"`
	Labels      []LabelPayload           `json:"labels"`
	State       string                   `json:"state"`
	Locked      bool                     `json:"locked"`
	Assignee    *User                    `json:"assignee"`
	Milestone   *MilestonePayload        `json:"milestone"`
	Comments    int                      `json:"comments"`
	PullRequest *IssuePullRequestPayload `json:"pull_request"`
	CreatedAt   time.Time                `json:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at"`
	ClosedAt    *time.Time               `json:"closed_at"`
	Body        string                   `json:"body"`
	ClosedBy    *User                    `json:"closed_by"`
}

type IssuePullRequestPayload struct {
	URL      string `json:"url"`
	HTMLURL  string `json:"html_url"`
	DiffURL  string `json:"diff_url"`
	PatchURL string `json:"patch_url"`
}

type IssueCommentPayload struct {
	URL       string    `json:"url"`
	HTMLURL   string    `json:"html_url"`
	IssueURL  string    `json:"issue_url"`
	ID        int       `json:"id"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
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
func (api *IssueAPI) GetIssueComment(commentID int) (*IssueCommentPayload, error) {
	url := api.getURL("/repos/:owner/:repo/issues/comments/" + strconv.Itoa(commentID))

	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	issueComment := &IssueCommentPayload{}

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&issueComment); err != nil {
		return nil, err
	}

	return issueComment, nil
}

// GetIssue gets an issue by issue number.
func (api *IssueAPI) GetIssue(issueNumber int) (*IssuePayload, error) {
	url := api.getURL("/repos/:owner/:repo/issues/" + strconv.Itoa(issueNumber))
	return api.GetIssueByURL(url)
}

// GetIssueByURL gets an issue by URL.
func (api *IssueAPI) GetIssueByURL(url string) (*IssuePayload, error) {
	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	issue := &IssuePayload{}

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&issue); err != nil {
		return nil, err
	}

	return issue, nil
}

// UpdateIssueAssignee updates an issue's assignee by issue number.
func (api *IssueAPI) UpdateIssueAssignee(issueNumber int, assignee string) (*IssuePayload, error) {
	url := api.getURL("/repos/:owner/:repo/issues/" + strconv.Itoa(issueNumber))
	return api.UpdateIssueAssigneeByURL(url, assignee)
}

// UpdateIssueAssigneeByURL updates an issue's assignee by issue URL.
func (api *IssueAPI) UpdateIssueAssigneeByURL(url, assignee string) (*IssuePayload, error) {
	// TODO (judwhite): there can be multiple assignees
	body := struct {
		Assignee string `json:"assignee"`
	}{assignee}

	return api.updateIssueByURL(url, body)
}

// UpdateIssueLabels updates an issue's labels by issue number. The labels passed become the new
// labels. See AddLabel and RemoveLabel to add/remove individual labels.
func (api *IssueAPI) UpdateIssueLabels(issueNumber int, labels []string) (*IssuePayload, error) {
	url := api.getURL("/repos/:owner/:repo/issues/" + strconv.Itoa(issueNumber))
	return api.UpdateIssueLabelsByURL(url, labels)
}

// UpdateIssueLabelsByURL updates an issue's labels by issue URL. The labels passed become the new
// labels. See AddLabel and RemoveLabel to add/remove individual labels.
func (api *IssueAPI) UpdateIssueLabelsByURL(url string, labels []string) (*IssuePayload, error) {
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

func (api *IssueAPI) updateIssueByURL(url string, body interface{}) (*IssuePayload, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpPatch(url, string(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issue IssuePayload

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&issue); err != nil {
		return nil, err
	}

	return &issue, nil
}
