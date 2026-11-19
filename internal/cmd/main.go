package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ChrisPJohnstone/git-open/internal/git"
)

func main(cmd *cobra.Command, args []string) error {
	remoteURL, err := git.GitRemoteURL()
	if err != nil {
		return err
	}
	webURL, err := parseWebURL(remoteURL)
	if err != nil {
		return err
	}
	fmt.Printf("URL: %s\n", webURL)
	return nil
}
