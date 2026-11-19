package cmd

import (
	"fmt"
	"io"
	"log/slog"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ChrisPJohnstone/git-open/internal/git"
	"github.com/ChrisPJohnstone/git-open/internal/parser"
)

func openURL(url string) error {
	slog.Debug(fmt.Sprintf("Attempting to open '%s'", url))
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		slog.Debug("OS identified as MacOS")
		cmd = exec.Command("open", url)
	default:
		slog.Debug("OS unidentified")
		cmd = exec.Command("xdg-open", url)
	}
	slog.Debug(fmt.Sprintf(
		"Attempting to run '%s'",
		strings.Join(cmd.Args, " "),
	))
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	return cmd.Start()
}

func main(cmd *cobra.Command, args []string) error {
	remoteURL, err := git.GitRemoteURL()
	if err != nil {
		return err
	}
	webURL, err := parser.ParseWebURL(remoteURL)
	if err != nil {
		return err
	}
	return openURL(webURL)
}
