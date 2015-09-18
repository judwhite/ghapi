package ghapi

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"errors"
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

	if body, err = ioutil.ReadAll(r.Body); err != nil {
		return "", nil, err
	}

	if ok := checkMAC(body, messageMAC, secret); !ok {
		return "", nil, errors.New("signature mismatch")
	}

	if eventType, err = getEventType(r); err != nil {
		return "", nil, err
	}

	return eventType, body, nil
}

func getEventType(r *http.Request) (GitHubEventType, error) {
	events := r.Header["X-Github-Event"]
	if events == nil || len(events) != 1 {
		return "", errors.New("\"X-Github-Event\" header not found")
	}
	return GitHubEventType(events[0]), nil
}

func getSignature(r *http.Request) (string, error) {
	sigs := r.Header["X-Hub-Signature"]
	if sigs == nil || len(sigs) != 1 {
		return "", errors.New("\"X-Hub-Signature\" header not found")
	}
	sig := sigs[0]
	if sig[:5] != "sha1=" {
		return "", errors.New("\"sha1=\" marker not found")
	}
	return sig[5:], nil
}

func getMAC(r *http.Request) ([]byte, error) {
	if sig, err := getSignature(r); err != nil {
		return nil, err
	} else {
		if messageMAC, err := hex.DecodeString(sig); err != nil {
			return nil, err
		} else {
			return messageMAC, nil
		}
	}
}

func checkMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha1.New, key)
	if _, err := mac.Write(message); err != nil {
		return false
	}
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
