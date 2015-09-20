package ghapi

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*
func TestReadRequest(t *testing.T) {
	signal := make(chan struct{}, 1)
	secret := []byte("webhook_secret")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signal <- struct{}{}
		_, _, err := ReadRequest(secret, r)
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	body := strings.NewReader("{ method: \"test\" }")
	req, err := http.NewRequest("POST", ts.URL, body)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	<-signal
}
*/

func TestReadRequest_FailsWhenXHubSignatureNotPresent(t *testing.T) {
	signal := make(chan struct{}, 1)
	secret := []byte("webhook_secret")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signal <- struct{}{}
		_, _, err := ReadRequest(secret, r)
		if err == nil {
			t.Fatal("expected X-Hub-Signature missing error; no error returned")
		} else if err != ErrSignatureNotFound {
			t.Fatalf("expected X-Hub-Signature missing error; got %v", err)
		}
	}))
	defer ts.Close()

	body := strings.NewReader("{ method: \"test\" }")
	req, err := http.NewRequest("POST", ts.URL, body)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	<-signal
}

func TestReadRequest_FailsWhenXHubSignatureMarkerNotPresent(t *testing.T) {
	signal := make(chan struct{}, 1)
	secret := []byte("webhook_secret")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signal <- struct{}{}
		_, _, err := ReadRequest(secret, r)
		if err == nil {
			t.Fatal("expected ErrSignatureMarkerNotFound error; no error returned")
		} else if err != ErrSignatureMarkerNotFound {
			t.Fatalf("expected ErrSignatureMarkerNotFound; got %v", err)
		}
	}))
	defer ts.Close()

	body := strings.NewReader("{ method: \"test\" }")
	req, err := http.NewRequest("POST", ts.URL, body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("X-Hub-Signature", "webhook_signature")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	<-signal
}

func TestReadRequest_FailsWhenXGithubEventNotPresent(t *testing.T) {
	signal := make(chan struct{}, 1)
	secret := []byte("webhook_secret")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signal <- struct{}{}
		_, _, err := ReadRequest(secret, r)
		if err == nil {
			t.Fatal("expected hex encoding error; no error returned")
		} else if err.Error() != "encoding/hex: invalid byte: U+0077 'w'" {
			t.Fatalf("expected hex encoding error; got %v", err)
		}
	}))
	defer ts.Close()

	body := strings.NewReader("{ method: \"test\" }")
	req, err := http.NewRequest("POST", ts.URL, body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("X-Hub-Signature", "sha1=webhook_secret")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	<-signal
}

func TestReadRequest_FailsWhenRequestBodyIsNil(t *testing.T) {
	secret := []byte("webhook_secret")
	req, err := http.NewRequest("POST", "http://asdf", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("X-Hub-Signature", "sha1=deadbeef")

	_, _, err = ReadRequest(secret, req)
	if err == nil {
		t.Fatal("expected ErrHttpRequestBodyNil; no error returned")
	} else if err != ErrHttpRequestBodyNil {
		t.Fatalf("expected ErrHttpRequestBodyNil; got %v", err)
	}
}

type errTooLargerReader struct {
}

func (r *errTooLargerReader) Read(p []byte) (int, error) {
	return 0, bytes.ErrTooLarge
}

func TestReadRequest_FailsWhenRequestBodyCannotBeRead(t *testing.T) {
	secret := []byte("webhook_secret")
	req, err := http.NewRequest("POST", "http://asdf", &errTooLargerReader{})
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("X-Hub-Signature", "sha1=deadbeef")

	_, _, err = ReadRequest(secret, req)
	if err == nil {
		t.Fatal("expected bytes.Buffer: too large; no error returned")
	} else if err != bytes.ErrTooLarge {
		t.Fatalf("expected bytes.Buffer: too large; got %v", err)
	}
}
