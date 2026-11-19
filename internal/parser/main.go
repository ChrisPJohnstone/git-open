package parser

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"
)

func url(domain string, extension string, user string, repo string) string {
	return fmt.Sprintf("https://%s.%s/%s/%s", domain, extension, user, repo)
}

func parseGithubURL(remote string) (string, error) {
	domain := "github"
	slog.Debug(fmt.Sprintf(
		"Attempting to parse %s URL from '%s'",
		domain,
		remote,
	))
	patternSSH := regexp.MustCompile(githubPatternSSH)
	parts := patternSSH.FindStringSubmatch(remote)
	if len(parts) == 4 {
		return url(domain, parts[1], parts[2], parts[3]), nil
	}
	return "", fmt.Errorf("unable to parse %s url for '%s'", domain, remote)
}

func parseBitbucketURL(remote string) (string, error) {
	domain := "bitbucket"
	slog.Debug(fmt.Sprintf(
		"Attempting to parse %s URL from '%s'",
		domain,
		remote,
	))
	patternSSH := regexp.MustCompile(bitbucketPatternSSH)
	parts := patternSSH.FindStringSubmatch(remote)
	if len(parts) == 4 {
		return url(domain, parts[1], parts[2], parts[3]), nil
	}
	return "", fmt.Errorf("unable to parse %s url for '%s'", domain, remote)
}

func ParseWebURL(remote string) (string, error) {
	slog.Debug(fmt.Sprintf("Attempting to parse web URL from '%s'", remote))
	if strings.Contains(remote, "github") {
		return parseGithubURL(remote)
	} else if strings.Contains(remote, "bitbucket") {
		return parseBitbucketURL(remote)
	}
	return "", fmt.Errorf("unable to determine web url for '%s'", remote)
}
