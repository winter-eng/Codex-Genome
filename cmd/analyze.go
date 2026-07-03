package cmd

import (
	"time"

	"github.com/codexgenome/codex-genome/internal/analyzer"
	"github.com/codexgenome/codex-genome/internal/report"
	"github.com/codexgenome/codex-genome/internal/scanner"
	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze <path>",
	Short: "Analyze a project directory",
	Args:  cobra.ExactArgs(1),
	RunE:  runAnalyze,
}

func runAnalyze(cmd *cobra.Command, args []string) error {
	start := time.Now()

	project, err := scanner.Scan(args[0])
	if err != nil {
		return err
	}

	rep := analyzer.Analyze(project)
	rep.Duration = time.Since(start)

	report.Render(rep)
	return nil
}
