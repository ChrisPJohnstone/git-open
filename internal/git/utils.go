package git

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/go-git/go-git/v6"
)

func repo() (*git.Repository, error) {
	repo, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{
		DetectDotGit: true,
	})
	if err != nil {
		return nil, err
	}
	slog.Debug("Repo opened succesfully")
	return repo, nil
}

func branch(repo *git.Repository) (string, error) {
	slog.Debug("Getting git HEAD")
	head, err := repo.Head()
	if err != nil {
		return "", err
	}
	return head.Name().Short(), nil
}

func remoteURL(repo *git.Repository) (string, error) {
	slog.Debug("Getting git remote")
	remotes, err := repo.Remotes()
	if err != nil {
		return "", err
	}
	switch len(remotes) {
	case 1:
	case 0:
		return "", errors.New("there are no remotes for this repo")
	default:
		return "", errors.New("repo has multiple remotes")
	}
	slog.Debug("Getting git remote URL")
	remoteConfig := remotes[0].Config()
	switch len(remoteConfig.URLs) {
	case 1:
	case 0:
		return "", fmt.Errorf("remote '%s' has no urls", remoteConfig.Name)
	default:
		return "", fmt.Errorf(
			"remote '%s' has multiple urls",
			remoteConfig.Name,
		)
	}
	return remoteConfig.URLs[0], nil
}
