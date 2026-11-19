package parser

import (
	"fmt"
	"log/slog"
	"regexp"
)

func githubBaseURL(remote string) (string, error) {
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

func githubURL(remote string, branch string, path string) (string, error) {
	url, err := githubBaseURL(remote)
	if err != nil {
		return "", err
	}
	url = fmt.Sprintf("%s/tree/%s", url, branch)
	if path != "" {
		url = fmt.Sprintf("%s/%s", url, path)
	}
	return url, nil
}
