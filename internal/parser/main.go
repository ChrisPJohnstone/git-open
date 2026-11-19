package parser

import (
	"fmt"
	"log/slog"
	"strings"
)

func url(domain string, extension string, user string, repo string) string {
	return fmt.Sprintf("https://%s.%s/%s/%s", domain, extension, user, repo)
}

func ParseWebURL(remote string, branch string) (string, error) {
	slog.Debug(fmt.Sprintf("Attempting to parse web URL from '%s'", remote))
	if strings.Contains(remote, "github") {
		return githubURL(remote, branch)
	} else if strings.Contains(remote, "bitbucket") {
		return bitbucketURL(remote, branch)
	}
	return "", fmt.Errorf("unable to determine web url for '%s'", remote)
}
