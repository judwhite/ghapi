package ghapi

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"hash"
	"io/ioutil"
	"net/http"
)

// ReadRequest takes an expected GitHub webhook secret and an *http.Request. If signature validation succeeds the
// GitHubEventType and request body as a byte slice are returned.
func ReadRequest(secret []byte, r *http.Request) (GitHubEventType, []byte, error) {
	var body []byte
	var err error
	var eventType GitHubEventType

	if r.Body == nil {
		return "", nil, ErrHTTPRequestBodyNil
	}

	if body, err = ioutil.ReadAll(r.Body); err != nil {
		return "", nil, err
	}

	if err = ValidateEvent(secret, body, r.Header); err != nil {
		return "", nil, err
	}

	if eventType, err = GetEventType(r.Header); err != nil {
		return "", nil, err
	}

	return eventType, body, nil
}

// ValidateEvent takes an expected GitHub webhook secret, a body, and an http.Header. If signature validation succeeds
// nil is returned.
//
// This function may be used when it's necessary to inspect the body of an event before validating its signature; for
// example, if the webhook stores separate secrets per organization but has a single endpoint to receive GitHub events
// it can extract the organization to retrieve the appropriate secret.
func ValidateEvent(secret []byte, body []byte, header http.Header) error {
	var err error

	if body == nil {
		return ErrHTTPRequestBodyNil
	}

	var messageMAC []byte
	if messageMAC, err = getMAC(header); err != nil {
		return err
	}

	sha1hmac := hmac.New(sha1.New, secret)
	if ok := checkMAC(sha1hmac, body, messageMAC); !ok {
		return ErrSignatureMismatch
	}

	return nil
}

// GetEventType returns the GitHubEventType based on the "X-Github-Event" header value from a received webhook request.
func GetEventType(header http.Header) (GitHubEventType, error) {
	events := header["X-Github-Event"]
	if events == nil || len(events) != 1 {
		return "", ErrGitHubEventNotFound
	}
	return GitHubEventType(events[0]), nil
}

func getSignature(header http.Header) (string, error) {
	sigs := header["X-Hub-Signature"]
	if sigs == nil || len(sigs) != 1 {
		return "", ErrSignatureNotFound
	}
	sig := sigs[0]
	if sig[:5] != "sha1=" {
		return "", ErrSignatureMarkerNotFound
	}
	return sig[5:], nil
}

func getMAC(header http.Header) ([]byte, error) {
	sig, err := getSignature(header)
	if err != nil {
		return nil, err
	}

	messageMAC, err := hex.DecodeString(sig)
	if err != nil {
		return nil, err
	}

	return messageMAC, nil
}

func checkMAC(mac hash.Hash, message, messageMAC []byte) bool {
	if _, err := mac.Write(message); err != nil {
		return false
	}
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
