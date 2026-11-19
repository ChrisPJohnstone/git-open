package cmd

import (
	"fmt"
	"regexp"
	"strings"
)

func parseGithubURL(remote string) (string, error) {
	patternSSH := regexp.MustCompile(`^git@.*github.com:([^/]*)/([^\.]*).*$`)
	parts := patternSSH.FindStringSubmatch(remote)
	if len(parts) == 3 {
		return fmt.Sprintf("https://github.com/%s/%s", parts[1], parts[2]), nil
	}
	return "", fmt.Errorf("Unable to parse github URL for '%s'", remote)
}

func parseBitbucketURL(remote string) (string, error) {
	return "", fmt.Errorf("Unable to parse bitbucket URL for '%s'", remote)
}

func parseWebURL(remote string) (string, error) {
	if strings.Contains(remote, "github") {
		return parseGithubURL(remote)
	} else if strings.Contains(remote, "bitbucket") {
		return parseBitbucketURL(remote)
	}
	return "", fmt.Errorf("Unable to determine web URL for '%s'", remote)
}
