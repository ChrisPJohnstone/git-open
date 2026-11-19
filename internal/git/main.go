package git

import (
	"log/slog"

	"github.com/go-git/go-git/v6"
)

type GitConfig struct {
	repo      *git.Repository
	Branch    string
	RemoteURL string
}

func NewConfig() (*GitConfig, error) {
	slog.Debug("Getting git repo")
	cfg := &GitConfig{}
	var err error
	cfg.repo, err = repo()
	if err != nil {
		return nil, err
	}
	cfg.Branch, err = branch(cfg.repo)
	if err != nil {
		return nil, err
	}
	cfg.RemoteURL, err = remoteURL(cfg.repo)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
