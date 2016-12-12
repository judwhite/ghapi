package ghapi

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"hash"
	"io/ioutil"
	"net/http"
)

func ReadRequest(secret []byte, r *http.Request) (GitHubEventType, []byte, error) {
	var body []byte
	var err error
	var messageMAC []byte
	var eventType GitHubEventType

	if messageMAC, err = getMAC(r); err != nil {
		return "", nil, err
	}

	if r.Body == nil {
		return "", nil, ErrHTTPRequestBodyNil
	}

	if body, err = ioutil.ReadAll(r.Body); err != nil {
		return "", nil, err
	}

	sha1hmac := hmac.New(sha1.New, secret)
	if ok := checkMAC(sha1hmac, body, messageMAC); !ok {
		return "", nil, ErrSignatureMismatch
	}

	if eventType, err = getEventType(r); err != nil {
		return "", nil, err
	}

	return eventType, body, nil
}

func getEventType(r *http.Request) (GitHubEventType, error) {
	events := r.Header["X-Github-Event"]
	if events == nil || len(events) != 1 {
		return "", ErrGitHubEventNotFound
	}
	return GitHubEventType(events[0]), nil
}

func getSignature(r *http.Request) (string, error) {
	sigs := r.Header["X-Hub-Signature"]
	if sigs == nil || len(sigs) != 1 {
		return "", ErrSignatureNotFound
	}
	sig := sigs[0]
	if sig[:5] != "sha1=" {
		return "", ErrSignatureMarkerNotFound
	}
	return sig[5:], nil
}

func getMAC(r *http.Request) ([]byte, error) {
	sig, err := getSignature(r)
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
