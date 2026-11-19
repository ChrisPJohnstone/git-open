package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main(cmd *cobra.Command, args []string) error {
	fmt.Printf("Hello World!\n")
	return nil
}
