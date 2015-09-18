package ghapi

type PullRequestAction string

const (
	Assigned    PullRequestAction = "assigned"
	Unassigned                    = "unassigned"
	Labeled                       = "labeled"
	Unlabeled                     = "unlabeled"
	Opened                        = "opened"
	Closed                        = "closed"
	Reopened                      = "reopened"
	Synchronize                   = "synchronize"
)

type PullRequestPayload struct {
}

func (p *PullRequestsApi) ListPullRequests(owner, repo string) ([]PullRequestPayload, error) {
	// TODO
	_ = p.getUrl("/repos/:owner/:repo/pulls")
	return nil, nil
}
