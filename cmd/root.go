package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "codex-genome",
	Short:        "Analyze software project structure and composition",
	Long:         `Codex Genome inspects your project's directory structure, file composition, and layout.`,
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}
