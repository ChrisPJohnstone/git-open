package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command = &cobra.Command{
	Use:               cmdName,
	Short:             cmdShortHelp,
	PersistentPreRunE: setupLogging,
	RunE:              main,
}

func init() {
	rootCmd.PersistentFlags().BoolP(
		verboseFlag,
		string(verboseFlag[0]),
		false,
		"enable verbose logging",
	)
}

func setupLogging(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()
	verbose, err := flags.GetBool(verboseFlag)
	if err != nil {
		return err
	}
	if verbose {
		opts := &slog.HandlerOptions{Level: slog.LevelDebug}
		handler := slog.NewJSONHandler(os.Stderr, opts)
		logger := slog.New(handler)
		slog.SetDefault(logger)
	}
	slog.Debug("Verbose logging enabled")
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
