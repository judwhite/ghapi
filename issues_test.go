package ghapi

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

const get_issue_1_response string = `{
  "id": 1,
  "url": "https://api.github.com/repos/octocat/Hello-World/issues/1347",
  "labels_url": "https://api.github.com/repos/octocat/Hello-World/issues/1347/labels{/name}",
  "comments_url": "https://api.github.com/repos/octocat/Hello-World/issues/1347/comments",
  "events_url": "https://api.github.com/repos/octocat/Hello-World/issues/1347/events",
  "html_url": "https://github.com/octocat/Hello-World/issues/1347",
  "number": 1347,
  "state": "open",
  "title": "Found a bug",
  "body": "I'm having a problem with this.",
  "user": {
    "login": "octocat",
    "id": 1,
    "avatar_url": "https://github.com/images/error/octocat_happy.gif",
    "gravatar_id": "",
    "url": "https://api.github.com/users/octocat",
    "html_url": "https://github.com/octocat",
    "followers_url": "https://api.github.com/users/octocat/followers",
    "following_url": "https://api.github.com/users/octocat/following{/other_user}",
    "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
    "organizations_url": "https://api.github.com/users/octocat/orgs",
    "repos_url": "https://api.github.com/users/octocat/repos",
    "events_url": "https://api.github.com/users/octocat/events{/privacy}",
    "received_events_url": "https://api.github.com/users/octocat/received_events",
    "type": "User",
    "site_admin": false
  },
  "labels": [
    {
      "url": "https://api.github.com/repos/octocat/Hello-World/labels/bug",
      "name": "bug",
      "color": "f29513"
    }
  ],
  "assignee": {
    "login": "octocat",
    "id": 1,
    "avatar_url": "https://github.com/images/error/octocat_happy.gif",
    "gravatar_id": "",
    "url": "https://api.github.com/users/octocat",
    "html_url": "https://github.com/octocat",
    "followers_url": "https://api.github.com/users/octocat/followers",
    "following_url": "https://api.github.com/users/octocat/following{/other_user}",
    "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
    "organizations_url": "https://api.github.com/users/octocat/orgs",
    "repos_url": "https://api.github.com/users/octocat/repos",
    "events_url": "https://api.github.com/users/octocat/events{/privacy}",
    "received_events_url": "https://api.github.com/users/octocat/received_events",
    "type": "User",
    "site_admin": false
  },
  "milestone": {
    "url": "https://api.github.com/repos/octocat/Hello-World/milestones/1",
    "html_url": "https://github.com/octocat/Hello-World/milestones/v1.0",
    "labels_url": "https://api.github.com/repos/octocat/Hello-World/milestones/1/labels",
    "id": 1002604,
    "number": 1,
    "state": "open",
    "title": "v1.0",
    "description": "Tracking milestone for version 1.0",
    "creator": {
      "login": "octocat",
      "id": 1,
      "avatar_url": "https://github.com/images/error/octocat_happy.gif",
      "gravatar_id": "",
      "url": "https://api.github.com/users/octocat",
      "html_url": "https://github.com/octocat",
      "followers_url": "https://api.github.com/users/octocat/followers",
      "following_url": "https://api.github.com/users/octocat/following{/other_user}",
      "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
      "organizations_url": "https://api.github.com/users/octocat/orgs",
      "repos_url": "https://api.github.com/users/octocat/repos",
      "events_url": "https://api.github.com/users/octocat/events{/privacy}",
      "received_events_url": "https://api.github.com/users/octocat/received_events",
      "type": "User",
      "site_admin": false
    },
    "open_issues": 4,
    "closed_issues": 8,
    "created_at": "2011-04-10T20:09:31Z",
    "updated_at": "2014-03-03T18:58:10Z",
    "closed_at": "2013-02-12T13:22:01Z",
    "due_on": "2012-10-09T23:39:01Z"
  },
  "locked": false,
  "comments": 0,
  "pull_request": {
    "url": "https://api.github.com/repos/octocat/Hello-World/pulls/1347",
    "html_url": "https://github.com/octocat/Hello-World/pull/1347",
    "diff_url": "https://github.com/octocat/Hello-World/pull/1347.diff",
    "patch_url": "https://github.com/octocat/Hello-World/pull/1347.patch"
  },
  "closed_at": null,
  "created_at": "2011-04-22T13:33:48Z",
  "updated_at": "2011-04-22T13:33:48Z",
  "closed_by": {
    "login": "octocat",
    "id": 1,
    "avatar_url": "https://github.com/images/error/octocat_happy.gif",
    "gravatar_id": "",
    "url": "https://api.github.com/users/octocat",
    "html_url": "https://github.com/octocat",
    "followers_url": "https://api.github.com/users/octocat/followers",
    "following_url": "https://api.github.com/users/octocat/following{/other_user}",
    "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
    "organizations_url": "https://api.github.com/users/octocat/orgs",
    "repos_url": "https://api.github.com/users/octocat/repos",
    "events_url": "https://api.github.com/users/octocat/events{/privacy}",
    "received_events_url": "https://api.github.com/users/octocat/received_events",
    "type": "User",
    "site_admin": false
  }
}`

func TestIssueApi_DeleteIssueComment(t *testing.T) {
	ts, api := makeGitHubApiTestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/repos/test_owner/test_repository/issues/comments/1" {

			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			body := string(b)

			expect(t, "DELETE", r.Method, "r.Method")
			expect(t, "", body, "r.Body")
			w.WriteHeader(200)
		} else {
			w.WriteHeader(404)
		}
	}))
	defer ts.Close()

	err := api.Issue.DeleteIssueComment(1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIssueApi_DeleteIssueComment_404(t *testing.T) {
	ts, api := makeGitHubApiTestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/repos/test_owner/test_repository/issues/comments/1" {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(404)
		}
	}))
	defer ts.Close()

	err := api.Issue.DeleteIssueComment(2)
	if err != nil {
		if e, ok := err.(ErrHttpError); !ok {
			t.Fatal("err is not of type ErrHttpError")
		} else {
			expect(t, 404, e.StatusCode, "e.StatusCode")
		}
	} else {
		t.Fatal("expected error")
	}
}

func TestIssueApi_GetIssue(t *testing.T) {
	ts, api := makeGitHubApiTestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/repos/test_owner/test_repository/issues/1347" {
			_, err := w.Write([]byte(get_issue_1_response))
			if err != nil {
				t.Fatal(err)
			}
		} else {
			w.WriteHeader(404)
		}
	}))
	defer ts.Close()

	issue, err := api.Issue.GetIssue(1347)
	if err != nil {
		t.Fatal(err)
	}
	//expect(t, "", issue, "")

	expect(t, 1, issue.Id, "issue.Id")
	expect(t, "https://api.github.com/repos/octocat/Hello-World/issues/1347", issue.Url, "issue.Url")
	expect(t, "https://api.github.com/repos/octocat/Hello-World/issues/1347/labels{/name}", issue.LabelsUrl, "issue.LabelsUrl")
	expect(t, "https://api.github.com/repos/octocat/Hello-World/issues/1347/comments", issue.CommentsUrl, "issue.CommentsUrl")
	expect(t, "https://api.github.com/repos/octocat/Hello-World/issues/1347/events", issue.EventsUrl, "issue.EventsUrl")
	expect(t, "https://github.com/octocat/Hello-World/issues/1347", issue.HtmlUrl, "issue.HtmlUrl")
	expect(t, 1347, issue.Number, "issue.Number")
	expect(t, "open", issue.State, "issue.State")
	expect(t, "Found a bug", issue.Title, "issue.Title")
	expect(t, "I'm having a problem with this.", issue.Body, "issue.Body")

	expect(t, "octocat", issue.User.Login, "issue.User.Login")
	expect(t, 1, issue.User.Id, "issue.User.Id")
	expect(t, "https://github.com/images/error/octocat_happy.gif", issue.User.AvatarUrl, "issue.User.AvatarUrl")
	expect(t, "", issue.User.GravatarId, "issue.User.GravatarId")
	expect(t, "https://api.github.com/users/octocat", issue.User.Url, "issue.User.Url")
	expect(t, "https://github.com/octocat", issue.User.HtmlUrl, "issue.User.HtmlUrl")
	expect(t, "https://api.github.com/users/octocat/followers", issue.User.FollowersUrl, "issue.User.FollowersUrl")
	expect(t, "https://api.github.com/users/octocat/following{/other_user}", issue.User.FollowingUrl, "issue.User.FollowingUrl")
	expect(t, "https://api.github.com/users/octocat/gists{/gist_id}", issue.User.GistsUrl, "issue.User.GistsUrl")
	expect(t, "https://api.github.com/users/octocat/starred{/owner}{/repo}", issue.User.StarredUrl, "issue.User.StarredUrl")
	expect(t, "https://api.github.com/users/octocat/subscriptions", issue.User.SubscriptionsUrl, "issue.User.SubscriptionsUrl")
	expect(t, "https://api.github.com/users/octocat/orgs", issue.User.OrganizationsUrl, "issue.User.OrganizationsUrl")
	expect(t, "https://api.github.com/users/octocat/repos", issue.User.ReposUrl, "issue.User.ReposUrl")
	expect(t, "https://api.github.com/users/octocat/events{/privacy}", issue.User.EventsUrl, "issue.User.EventsUrl")
	expect(t, "https://api.github.com/users/octocat/received_events", issue.User.ReceivedEventsUrl, "issue.User.ReceivedEventsUrl")
	expect(t, "User", issue.User.Type, "issue.User.Type")
	expect(t, false, issue.User.SiteAdmin, "issue.User.SiteAdmin")

	expect(t, 1, len(issue.Labels), "len(issue.Labels)")
	expect(t, "https://api.github.com/repos/octocat/Hello-World/labels/bug", issue.Labels[0].Url, "issue.Labels[0].Url")
	expect(t, "bug", issue.Labels[0].Name, "issue.Labels[0].Name")
	expect(t, "f29513", issue.Labels[0].Color, "issue.Labels[0].Color")

	expectNotNil(t, issue.Assignee, "issue.Assignee")

	expect(t, "octocat", issue.Assignee.Login, "issue.Assignee.Login")
	expect(t, 1, issue.Assignee.Id, "issue.Assignee.Id")
	expect(t, "https://github.com/images/error/octocat_happy.gif", issue.Assignee.AvatarUrl, "issue.Assignee.AvatarUrl")
	expect(t, "", issue.Assignee.GravatarId, "issue.Assignee.GravatarId")
	expect(t, "https://api.github.com/users/octocat", issue.Assignee.Url, "issue.Assignee.Url")
	expect(t, "https://github.com/octocat", issue.Assignee.HtmlUrl, "issue.Assignee.HtmlUrl")
	expect(t, "https://api.github.com/users/octocat/followers", issue.Assignee.FollowersUrl, "issue.Assignee.FollowersUrl")
	expect(t, "https://api.github.com/users/octocat/following{/other_user}", issue.Assignee.FollowingUrl, "issue.Assignee.FollowingUrl")
	expect(t, "https://api.github.com/users/octocat/gists{/gist_id}", issue.Assignee.GistsUrl, "issue.Assignee.GistsUrl")
	expect(t, "https://api.github.com/users/octocat/starred{/owner}{/repo}", issue.Assignee.StarredUrl, "issue.Assignee.StarredUrl")
	expect(t, "https://api.github.com/users/octocat/subscriptions", issue.Assignee.SubscriptionsUrl, "issue.Assignee.SubscriptionsUrl")
	expect(t, "https://api.github.com/users/octocat/orgs", issue.Assignee.OrganizationsUrl, "issue.Assignee.OrganizationsUrl")
	expect(t, "https://api.github.com/users/octocat/repos", issue.Assignee.ReposUrl, "issue.Assignee.ReposUrl")
	expect(t, "https://api.github.com/users/octocat/events{/privacy}", issue.Assignee.EventsUrl, "issue.Assignee.EventsUrl")
	expect(t, "https://api.github.com/users/octocat/received_events", issue.Assignee.ReceivedEventsUrl, "issue.Assignee.ReceivedEventsUrl")
	expect(t, "User", issue.Assignee.Type, "issue.Assignee.Type")
	expect(t, false, issue.Assignee.SiteAdmin, "issue.Assignee.SiteAdmin")

	// https://developer.github.com/v3/issues/#get-a-single-issue
	// TODO: milestone

	expect(t, false, issue.Locked, "issue.Locked")
	expect(t, 0, issue.Comments, "issue.Comments")

	// TODO: PullRequest

	expectedCreatedAt, err := time.Parse(time.RFC3339, "2011-04-22T13:33:48Z")
	if err != nil {
		t.Fatal(err)
	}
	expectedUpdatedAt, err := time.Parse(time.RFC3339, "2011-04-22T13:33:48Z")
	if err != nil {
		t.Fatal(err)
	}

	expectNil(t, issue.ClosedAt, "issue.ClosedAt")
	expect(t, expectedCreatedAt, issue.CreatedAt, "issue.CreatedAt")
	expect(t, expectedUpdatedAt, issue.UpdatedAt, "issue.UpdatedAt")

	// TODO Closed By
	//issue.ClosedBy
}
