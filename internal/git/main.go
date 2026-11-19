package git

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/go-git/go-git/v6"
)

func gitRepo() (*git.Repository, error) {
	repo, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{
		DetectDotGit: true,
	})
	if err != nil {
		return nil, err
	}
	slog.Debug("Repo opened succesfully")
	return repo, nil
}

func GitRemoteURL() (string, error) {
	slog.Debug("Getting git repo")
	repo, err := gitRepo()
	if err != nil {
		return "", err
	}
	slog.Debug("Getting git remote")
	remotes, err := repo.Remotes()
	if err != nil {
		return "", err
	}
	switch len(remotes) {
	case 1:
	case 0:
		return "", errors.New("There are no remotes for this repo")
	default:
		return "", errors.New("Repo has multiple remotes")
	}
	slog.Debug("Getting git remote URL")
	remoteConfig := remotes[0].Config()
	switch len(remoteConfig.URLs) {
	case 1:
	case 0:
		return "", fmt.Errorf("Remote '%s' has no URLs", remoteConfig.Name)
	default:
		return "", fmt.Errorf(
			"Remote '%s' has multiple URLs",
			remoteConfig.Name,
		)
	}
	return remoteConfig.URLs[0], nil
}
