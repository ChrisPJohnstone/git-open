package parser

import (
	"fmt"
	"log/slog"
	"regexp"
)

func bitbucketBaseURL(remote string) (string, error) {
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

func bitbucketURL(remote string, hash string) (string, error) {
	url, err := bitbucketBaseURL(remote)
	if err != nil {
		return "", err
	}
	url = fmt.Sprintf("%s/src/%s", url, hash)
	return url, nil
}
