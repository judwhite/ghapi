package ghapi

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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

	waitSignal(t, signal)
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

	waitSignal(t, signal)
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

	waitSignal(t, signal)
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

func Test_checkMAC_MatchesExpected(t *testing.T) {
	secret := []byte("key")
	message := []byte("message")
	messageMAC, err := hex.DecodeString("2088df74d5f2146b48146caf4965377e9d0be3a4")
	if err != nil {
		t.Fatal(err)
	}

	sha1hmac := hmac.New(sha1.New, secret)
	ok := checkMAC(sha1hmac, message, messageMAC)

	if !ok {
		t.Fatal("messageMAC doesn't match")
	}
}

type hashError struct {
	writeCalled *bool
	sumCalled   *bool
}

func (h hashError) BlockSize() int {
	return 0
}

func (h hashError) Reset() {
}

func (h hashError) Size() int {
	return 0
}

func (h hashError) Sum(b []byte) []byte {
	*h.sumCalled = true
	return nil
}

func (h hashError) Write(b []byte) (int, error) {
	*h.writeCalled = true
	return 0, errors.New("write error")
}

func Test_checkHMAC_FailsWhenHashReadReturnsErr(t *testing.T) {
	message := []byte("message")
	messageMAC, err := hex.DecodeString("2088df74d5f2146b48146caf4965377e9d0be3a4")
	if err != nil {
		t.Fatal(err)
	}

	// err will never be set for hmac.New(sha1.New, secret)
	// see:
	// - https://golang.org/src/crypto/hmac/hmac.go#L63
	// - https://golang.org/src/crypto/sha1/sha1.go#L61

	hmacErrorOnWrite := hashError{}
	f := false
	hmacErrorOnWrite.writeCalled = &f
	f2 := false
	hmacErrorOnWrite.sumCalled = &f2

	ok := checkMAC(hmacErrorOnWrite, message, messageMAC)

	expect(t, false, ok, "ok")
	expect(t, true, *hmacErrorOnWrite.writeCalled, "hmacErrorOnWrite.writeCalled")
	expect(t, false, *hmacErrorOnWrite.sumCalled, "hmacErrorOnWrite.sumCalled")
}

func TestReadRequest_FailsWhenMACDoesntMatchExpected(t *testing.T) {
	signal := make(chan struct{}, 1)
	secret := []byte("key")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signal <- struct{}{}
		_, _, err := ReadRequest(secret, r)
		expectNotNil(t, err, "err")
		expect(t, ErrSignatureMismatch, err, "err")
	}))
	defer ts.Close()

	body := strings.NewReader("messageJUNKAPPENDED") // the message to be signed
	req, err := http.NewRequest("POST", ts.URL, body)
	req.Header.Add("X-Hub-Signature", "sha1=2088df74d5f2146b48146caf4965377e9d0be3a4")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	waitSignal(t, signal)
}

func TestReadRequest_FailsWhenXGithubEventHeaderNotFound(t *testing.T) {
	signal := make(chan struct{}, 1)
	secret := []byte("key")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signal <- struct{}{}
		_, _, err := ReadRequest(secret, r)
		expectNotNil(t, err, "err")
		expect(t, ErrGitHubEventNotFound, err, "err")
	}))
	defer ts.Close()

	body := strings.NewReader("message") // the message to be signed
	req, err := http.NewRequest("POST", ts.URL, body)
	req.Header.Add("X-Hub-Signature", "sha1=2088df74d5f2146b48146caf4965377e9d0be3a4")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	waitSignal(t, signal)
}

func TestReadRequest_MatchesExpected(t *testing.T) {
	signal := make(chan struct{}, 1)
	secret := []byte("key")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signal <- struct{}{}
		eventType, bodyBytes, err := ReadRequest(secret, r)

		expectNil(t, err, "err")
		expectNotNil(t, eventType, "eventType")
		expectNotNil(t, bodyBytes, "bodyBytes")

		expect(t, PushEventType, eventType, "eventType")
		expect(t, "message", string(bodyBytes), "bodyBytes")
	}))
	defer ts.Close()

	body := strings.NewReader("message") // the message to be signed
	req, err := http.NewRequest("POST", ts.URL, body)
	req.Header.Add("X-Hub-Signature", "sha1=2088df74d5f2146b48146caf4965377e9d0be3a4")
	req.Header.Add("X-Github-Event", "push")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	waitSignal(t, signal)
}
