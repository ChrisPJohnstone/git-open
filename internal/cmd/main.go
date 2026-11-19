package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ChrisPJohnstone/git-open/internal/git"
)

func main(cmd *cobra.Command, args []string) error {
	remote, err := git.GitRemoteURL()
	if err != nil {
		return err
	}
	fmt.Printf("Remote: %s\n", remote)
	return nil
}
