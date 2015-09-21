package ghapi

import (
	"encoding/json"
	"strconv"
	"time"
)

type IssuePayload struct {
	Url         string                   `json:"url"`
	LabelsUrl   string                   `json:"labels_url"`
	CommentsUrl string                   `json:"comments_url"`
	EventsUrl   string                   `json:"events_url"`
	HtmlUrl     string                   `json:"html_url"`
	Id          int                      `json:"id"`
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
	Url      string `json:"url"`
	HtmlUrl  string `json:"html_url"`
	DiffUrl  string `json:"diff_url"`
	PatchUrl string `json:"patch_url"`
}

type IssueCommentPayload struct {
	Url       string    `json:"url"`
	HtmlUrl   string    `json:"html_url"`
	IssueUrl  string    `json:"issue_url"`
	Id        int       `json:"id"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
}

func (api *IssueApi) DeleteIssueComment(commentId int) error {
	url := api.getUrl("/repos/:owner/:repo/issues/comments/" + strconv.Itoa(commentId))
	return api.DeleteIssueCommentByUrl(url)
}

func (api *IssueApi) DeleteIssueCommentByUrl(url string) error {
	resp, err := api.httpDelete(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (api *IssueApi) GetIssueComment(commentId int) (*IssueCommentPayload, error) {
	url := api.getUrl("/repos/:owner/:repo/issues/comments/" + strconv.Itoa(commentId))

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

func (api *IssueApi) GetIssue(issueNumber int) (*IssuePayload, error) {
	url := api.getUrl("/repos/:owner/:repo/issues/" + strconv.Itoa(issueNumber))
	return api.GetIssueByUrl(url)
}

func (api *IssueApi) GetIssueByUrl(url string) (*IssuePayload, error) {
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

func (api *IssueApi) UpdateIssueAssignee(issueNumber int, assignee string) (*IssuePayload, error) {
	url := api.getUrl("/repos/:owner/:repo/issues/" + strconv.Itoa(issueNumber))
	return api.UpdateIssueAssigneeByUrl(url, assignee)
}

func (api *IssueApi) UpdateIssueAssigneeByUrl(url, assignee string) (*IssuePayload, error) {
	body := struct {
		Assignee string `json:"assignee"`
	}{assignee}

	return api.updateIssueByUrl(url, body)
}

func (api *IssueApi) updateIssueByUrl(url string, body interface{}) (*IssuePayload, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpPatch(url, string(b))
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
