package ghapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	createRefResponse string = `
	{
		"ref": "refs/heads/featureA",
		"url": "http://127.0.0.1:5285/repos/test_owner/test_repository/git/refs/heads/featureA",
		"object": {
		  "type": "commit",
		  "sha": "aa218f56b14c9653891f9e74264a383fa43fefbd",
		  "url": "http://127.0.0.1:5285/repos/test_owner/test_repository/git/commits/aa218f56b14c9653891f9e74264a383fa43fefbd"
		}
	  }`
)

func TestRefsApi_CreateRefs(t *testing.T) {
	ts, api, signal := makeGitHubAPITestServer(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s", r.URL.Path)
		if r.URL != nil && r.URL.Path == "/repos/test_owner/test_repository/git/refs/" {
			b, err := ioutil.ReadAll(r.Body)

			expectNil(t, err, "err")

			_, err = w.Write([]byte(createRefResponse))

			expectNil(t, err, "err")
			expect(t, "POST", r.Method, "r.Method")
			expect(t, `{"ref":"refs/heads/featureA","sha":"aa218f56b14c9653891f9e74264a383fa43fefbd"}`, string(b), "r.Body")
		} else {
			w.WriteHeader(404)
		}
	})
	defer ts.Close()

	result, err := api.Refs.CreateRef("refs/heads/featureA", "aa218f56b14c9653891f9e74264a383fa43fefbd")
	waitSignal(t, signal)

	if err != nil {
		t.Fatal(err)
	}

	expect(t, "refs/heads/featureA", result.Ref, "result.Ref")
	expect(t, "http://127.0.0.1:5285/repos/test_owner/test_repository/git/refs/heads/featureA", result.URL, "result.URL")
	expect(t, "aa218f56b14c9653891f9e74264a383fa43fefbd", result.Object.SHA, "result.Object.SHA")
	expect(t, "commit", result.Object.Type, "result.Object.Type")
	expect(t, "http://127.0.0.1:5285/repos/test_owner/test_repository/git/commits/aa218f56b14c9653891f9e74264a383fa43fefbd", result.Object.URL, "result.Object.URL")
}
