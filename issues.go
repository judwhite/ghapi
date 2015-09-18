package ghapi

import (
	"encoding/json"
	"strconv"
	"time"
)

type IssuePayload struct {
	Url         string         `json:"url"`
	LabelsUrl   string         `json:"labels_url"`
	CommentsUrl string         `json:"comments_url"`
	EventsUrl   string         `json:"events_url"`
	HtmlUrl     string         `json:"html_url"`
	Id          int            `json:"id"`
	Number      int            `json:"number"`
	Title       string         `json:"title"`
	User        UserPayload    `json:"user"`
	Labels      []LabelPayload `json:"labels"`
	State       string         `json:"state"`
	Locked      bool           `json:"locked"`
	Assignee    *UserPayload   `json:"assignee"`
	//Milestone interface{} `json:"milestone"`
	Comments  int        `json:"comments"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	ClosedAt  *time.Time `json:"closed_at"`
	Body      string     `json:"body"`
}

type IssueCommentPayload struct {
	Url       string      `json:"url"`
	HtmlUrl   string      `json:"html_url"`
	IssueUrl  string      `json:"issue_url"`
	Id        int         `json:"id"`
	User      UserPayload `json:"user"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Body      string      `json:"body"`
}

func (api *IssueApi) DeleteIssueComment(commentId int) error {
	// TODO
	url := api.getUrl("/repos/:owner/:repo/issues/comments/" + strconv.Itoa(commentId))
	_, err := api.httpDelete(url)
	return err
}

func (api *IssueApi) GetIssueComment(commentId int) (*IssueCommentPayload, error) {
	url := api.getUrl("/repos/:owner/:repo/issues/comments/" + strconv.Itoa(commentId))

	resp, err := api.httpGet(url)
	if err != nil {
		return nil, err
	}

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

	issue := &IssuePayload{}

	j := json.NewDecoder(resp.Body)
	if err = j.Decode(&issue); err != nil {
		return nil, err
	}

	return issue, nil
}
