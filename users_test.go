package ghapi

import (
	"io/ioutil"
	"net/http"
	"testing"
)

const getUserOctocatResponse string = `{
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
  "site_admin": false,
  "name": "monalisa octocat",
  "company": "GitHub",
  "blog": "https://github.com/blog",
  "location": "San Francisco",
  "email": "octocat@github.com",
  "hireable": false,
  "bio": "There once was...",
  "public_repos": 2,
  "public_gists": 1,
  "followers": 20,
  "following": 0,
  "created_at": "2008-01-14T04:33:35Z",
  "updated_at": "2008-01-14T04:33:35Z"
}`

const getUserOctocatOrganizationsResponse string = `[
  {
    "login": "github",
    "id": 1,
    "url": "https://api.github.com/orgs/github",
    "avatar_url": "https://github.com/images/error/octocat_happy.gif",
    "description": "A great organization"
  }
]`

func TestUserApi_GetUser(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/users/octocat" {
			b, err := ioutil.ReadAll(r.Body)

			expectNil(t, err, "err")

			_, err = w.Write([]byte(getUserOctocatResponse))

			expectNil(t, err, "err")
			expect(t, "GET", r.Method, "r.Method")
			expect(t, "", string(b), "r.Body")
		} else {
			w.WriteHeader(404)
		}
	})
	defer ts.Close()

	u, err := api.User.GetUser("octocat")
	waitSignal(t, signal)

	if err != nil {
		t.Fatal(err)
	}

	expect(t, "octocat", u.Login, "u.Login")
	expect(t, 1, u.ID, "u.Id")
	expect(t, "https://github.com/images/error/octocat_happy.gif", u.AvatarURL, "u.AvatarUrl")
	expect(t, "", u.GravatarID, "u.GravatarId")
	expect(t, "https://api.github.com/users/octocat", u.URL, "u.Url")
	expect(t, "https://github.com/octocat", u.HTMLURL, "u.HtmlUrl")
	expect(t, "https://api.github.com/users/octocat/followers", u.FollowersURL, "u.FollowersUrl")
	expect(t, "https://api.github.com/users/octocat/following{/other_user}", u.FollowingURL, "u.FollowingUrl")
	expect(t, "https://api.github.com/users/octocat/gists{/gist_id}", u.GistsURL, "u.GistsUrl")
	expect(t, "https://api.github.com/users/octocat/starred{/owner}{/repo}", u.StarredURL, "u.StarredUrl")
	expect(t, "https://api.github.com/users/octocat/subscriptions", u.SubscriptionsURL, "u.SubscriptionsUrl")
	expect(t, "https://api.github.com/users/octocat/orgs", u.OrganizationsURL, "u.OrganizationsUrl")
	expect(t, "https://api.github.com/users/octocat/repos", u.ReposURL, "u.ReposUrl")
	expect(t, "https://api.github.com/users/octocat/events{/privacy}", u.EventsURL, "u.EventsUrl")
	expect(t, "https://api.github.com/users/octocat/received_events", u.ReceivedEventsURL, "u.ReceivedEventsUrl")
	expect(t, "User", u.Type, "u.Type")
	expect(t, false, u.SiteAdmin, "u.SiteAdmin")

	expect(t, "monalisa octocat", u.Name, "u.Name")
	expect(t, "GitHub", u.Company, "u.Company")
	expect(t, "https://github.com/blog", u.Blog, "u.Blog")
	expect(t, "San Francisco", u.Location, "u.Location")
	expect(t, "octocat@github.com", u.Email, "u.Email")
	expect(t, false, u.Hireable, "u.Hireable")
	expect(t, "There once was...", u.Bio, "u.Bio")
	expect(t, 2, u.PublicRepos, "u.PublicRepos")
	expect(t, 1, u.PublicGists, "u.PublicGists")
	expect(t, 20, u.Followers, "u.Followers")
	expect(t, 0, u.Following, "u.Following")
	expect(t, date("2008-01-14T04:33:35Z"), u.CreatedAt, "u.CreatedAt")
	expect(t, date("2008-01-14T04:33:35Z"), u.UpdatedAt, "u.UpdatedAt")
}

func TestUserApi_GetUser_ReturnsErrOnHttpErr(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/users/octocat" {
			w.WriteHeader(500)
		} else {
			w.WriteHeader(404)
		}
	})
	defer ts.Close()

	u, err := api.User.GetUser("octocat")
	waitSignal(t, signal)

	expectNil(t, u, "u")
	expectErrHTTPError500(t, err)
}

func TestUserApi_GetUser_ReturnsErrOnJsonDecodeErr(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/users/octocat" {
			_, err := w.Write([]byte("junk" + getUserOctocatOrganizationsResponse))

			expectNil(t, err, "err")
		} else {
			w.WriteHeader(404)
		}
	})
	defer ts.Close()

	u, err := api.User.GetUser("octocat")
	waitSignal(t, signal)

	expectNil(t, u, "u")
	expectJSONSyntaxError(t, err, "invalid character 'j' looking for beginning of value")
}

func TestUserApi_GetOrganizations(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/users/octocat/orgs" {
			b, err := ioutil.ReadAll(r.Body)

			expectNil(t, err, "err")

			_, err = w.Write([]byte(getUserOctocatOrganizationsResponse))

			expectNil(t, err, "err")
			expect(t, "GET", r.Method, "r.Method")
			expect(t, "", string(b), "r.Body")
		} else {
			w.WriteHeader(404)
		}
	})
	defer ts.Close()

	orgs, err := api.User.GetPublicOrganizations("octocat")
	waitSignal(t, signal)

	if err != nil {
		t.Fatal(err)
	}

	expectNotNil(t, orgs, "orgs")
	expect(t, 1, len(orgs), "len(orgs)")
	org := orgs[0]
	expect(t, "github", org.Login, "org.Login")
	expect(t, 1, org.ID, "org.Id")
	expect(t, "https://api.github.com/orgs/github", org.URL, "org.Url")
	expect(t, "https://github.com/images/error/octocat_happy.gif", org.AvatarURL, "org.AvatarUrl")
	expect(t, "A great organization", org.Description, "org.Description")
}

func TestUserApi_GetOrganizations_ReturnsErrOnHttpErr(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/users/octocat/orgs" {
			w.WriteHeader(500)
		} else {
			w.WriteHeader(404)
		}
	})
	defer ts.Close()

	u, err := api.User.GetPublicOrganizations("octocat")
	waitSignal(t, signal)

	expectNil(t, u, "u")
	expectErrHTTPError500(t, err)
}

func TestUserApi_GetOrganizations_ReturnsErrOnJsonDecodeErr(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL != nil && r.URL.Path == "/users/octocat/orgs" {
			_, err := w.Write([]byte("junk" + getUserOctocatOrganizationsResponse))

			expectNil(t, err, "err")
		} else {
			w.WriteHeader(404)
		}
	})
	defer ts.Close()

	u, err := api.User.GetPublicOrganizations("octocat")
	waitSignal(t, signal)

	expectNil(t, u, "u")
	expectJSONSyntaxError(t, err, "invalid character 'j' looking for beginning of value")
}
