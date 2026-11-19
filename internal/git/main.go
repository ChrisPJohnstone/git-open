package git

import (
	"log/slog"

	"github.com/go-git/go-git/v6"
)

type GitConfig struct {
	repo      *git.Repository
	Branch    string
	Hash      string
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
	slog.Debug("Getting git HEAD")
	head, err := cfg.repo.Head()
	if err != nil {
		return nil, err
	}
	cfg.Branch = head.Name().Short()
	cfg.Hash = head.Hash().String()
	cfg.RemoteURL, err = remoteURL(cfg.repo)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
