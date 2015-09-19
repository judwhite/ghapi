package ghapi

import (
	"time"
)

type MilestonePayload struct {
	Url          string      `json:"url"`
	HtmlUrl      string      `json:"html_url"`
	LabelsUrl    string      `json:"labels_url"`
	Id           int         `json:"id"`
	Number       int         `json:"number"`
	State        string      `json:"state"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	Creator      UserPayload `json:"creator"`
	OpenIssues   int         `json:"open_issues"`
	ClosedIssues int         `json:"closed_issues"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	ClosedAt     *time.Time  `json:"closed_at"`
	DueOn        *time.Time  `json:"due_on"`
}
