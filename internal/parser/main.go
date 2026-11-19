package parser

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/ChrisPJohnstone/git-open/internal/git"
)

func url(domain string, extension string, user string, repo string) string {
	return fmt.Sprintf("https://%s.%s/%s/%s", domain, extension, user, repo)
}

func ParseWebURL(gitCfg *git.GitConfig) (string, error) {
	slog.Debug(fmt.Sprintf(
		"Attempting to parse web URL from '%s'",
		gitCfg.RemoteURL,
	))
	if strings.Contains(gitCfg.RemoteURL, "github") {
		return githubURL(gitCfg.RemoteURL, gitCfg.Branch)
	} else if strings.Contains(gitCfg.RemoteURL, "bitbucket") {
		return bitbucketURL(gitCfg.RemoteURL, gitCfg.Hash)
	}
	return "", fmt.Errorf(
		"unable to determine web url for '%s'",
		gitCfg.RemoteURL,
	)
}
