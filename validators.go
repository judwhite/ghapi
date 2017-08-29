package ghapi

import (
	"errors"
	"fmt"
	"strings"
)

// ValidateOwnerName returns an error of the owner name is not valid according to GitHub rules.
func ValidateOwnerName(owner string) error {
	if len(owner) == 0 {
		return errors.New("owner is empty")
	}
	if len(owner) > 39 {
		return errors.New("owner is too long (maximum is 39 characters)")
	}

	badMatch := "owner may only contain alphanumeric characters or single hyphens, and cannot begin or end with a hyphen"
	if strings.HasPrefix(owner, "-") || strings.HasSuffix(owner, "-") || strings.Contains(owner, "--") {
		return errors.New(badMatch)
	}
	lowerOwner := strings.ToLower(owner)
	for _, r := range lowerOwner {
		if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
			return fmt.Errorf("owner name '%s' contains invalid character '%c'", owner, r)
		}
	}
	return nil
}

// ValidateRepoName returns an error of the repo name is not valid according to GitHub rules.
func ValidateRepoName(repo string) error {
	if len(repo) == 0 {
		return errors.New("repo is empty")
	}
	if len(repo) > 100 {
		return errors.New("repo is too long (maximum is 100 characters)")
	}
	lowerRepo := strings.ToLower(repo)
	if repo == "." || repo == ".." || strings.HasSuffix(lowerRepo, ".git") || strings.HasSuffix(lowerRepo, ".wiki") {
		return fmt.Errorf("repo name '%s' is reserved", repo)
	}
	for _, r := range lowerRepo {
		if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '.' && r != '_' && r != '-' {
			return fmt.Errorf("repo name '%s' contains invalid character '%c'", repo, r)
		}
	}
	return nil
}
