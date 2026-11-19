package cmd

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"
)

func parseGithubURL(remote string) (string, error) {
	slog.Debug(fmt.Sprintf("Attempting to parse Github URL from '%s'", remote))
	patternSSH := regexp.MustCompile(`^git@.*github.com:([^/]*)/([^\.]*).*$`)
	parts := patternSSH.FindStringSubmatch(remote)
	if len(parts) == 3 {
		return fmt.Sprintf("https://github.com/%s/%s", parts[1], parts[2]), nil
	}
	return "", fmt.Errorf("Unable to parse github URL for '%s'", remote)
}

func parseBitbucketURL(remote string) (string, error) {
	slog.Debug(fmt.Sprintf(
		"Attempting to parse Bitbucket URL from '%s'",
		remote,
	))
	patternSSH := regexp.MustCompile(`^git@.*bitbucket.org:([^/]*)/([^\.]*).*$`)
	parts := patternSSH.FindStringSubmatch(remote)
	if len(parts) == 3 {
		return fmt.Sprintf("https://bitbucket.org/%s/%s", parts[1], parts[2]), nil
	}
	return "", fmt.Errorf("Unable to parse bitbucket URL for '%s'", remote)
}

func parseWebURL(remote string) (string, error) {
	slog.Debug(fmt.Sprintf("Attempting to parse web URL from '%s'", remote))
	if strings.Contains(remote, "github") {
		return parseGithubURL(remote)
	} else if strings.Contains(remote, "bitbucket") {
		return parseBitbucketURL(remote)
	}
	return "", fmt.Errorf("Unable to determine web URL for '%s'", remote)
}
