package report

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/codexgenome/codex-genome/internal/analyzer"
)

const (
	panelWidth = 62
	barWidth   = 20
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7C3AED"))

	pathStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#64748B"))

	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#94A3B8"))

	valueStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#A78BFA"))

	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#E2E8F0"))

	barFilledStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7C3AED"))

	barEmptyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#334155"))

	sepStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#1E293B"))

	dimStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#475569"))
)

// Render prints the analysis report to stdout using Lipgloss styles.
func Render(rep *analyzer.Report) {
	sep := sepStyle.Render(strings.Repeat("─", panelWidth))

	fmt.Println()
	fmt.Println(titleStyle.Render(" ◈ Codex Genome") + "  " + pathStyle.Render(rep.RootPath))
	fmt.Println(sep)
	fmt.Println()

	fmt.Println("  " + labelStyle.Width(24).Render("Total Files") +
		valueStyle.Render(fmt.Sprintf("%d", rep.TotalFiles)))
	fmt.Println("  " + labelStyle.Width(24).Render("Total Directories") +
		valueStyle.Render(fmt.Sprintf("%d", rep.TotalDirs)))

	if len(rep.Extensions) > 0 {
		fmt.Println()
		fmt.Println(sep)
		fmt.Println()
		fmt.Println(headerStyle.Render("  File Extensions"))
		fmt.Println()

		for _, ec := range rep.Extensions {
			filled := 0
			if rep.TotalFiles > 0 {
				filled = (ec.Count * barWidth) / rep.TotalFiles
			}
			bar := barFilledStyle.Render(strings.Repeat("█", filled)) +
				barEmptyStyle.Render(strings.Repeat("░", barWidth-filled))

			fmt.Println("  " +
				labelStyle.Width(22).Render(ec.Ext) +
				valueStyle.Render(fmt.Sprintf("%5d", ec.Count)) +
				"  " + bar)
		}
	}

	if len(rep.Languages) > 0 {
		fmt.Println()
		fmt.Println(sep)
		fmt.Println()
		fmt.Println(headerStyle.Render("  Language Profile"))
		fmt.Println()

		primary := rep.PrimaryLanguage
		if primary == "" {
			primary = "-"
		}
		fmt.Println("  " + labelStyle.Width(24).Render("Primary Language") +
			valueStyle.Render(primary))
		fmt.Println()

		for _, ls := range rep.Languages {
			filled := 0
			if rep.TotalFiles > 0 {
				filled = (ls.FileCount * barWidth) / rep.TotalFiles
			}
			bar := barFilledStyle.Render(strings.Repeat("█", filled)) +
				barEmptyStyle.Render(strings.Repeat("░", barWidth-filled))

			fmt.Println("  " +
				labelStyle.Width(16).Render(ls.Language) +
				valueStyle.Render(fmt.Sprintf("%5d", ls.FileCount)) +
				"  " + bar +
				"  " + dimStyle.Render(fmt.Sprintf("%5.1f%%", ls.Percentage)))
		}

		fmt.Println()
		fmt.Println("  " + labelStyle.Width(24).Render("Total Languages") +
			valueStyle.Render(fmt.Sprintf("%d", rep.TotalLanguages)))
	}

	fmt.Println()
	fmt.Println(sep)
	fmt.Println("  " + dimStyle.Width(16).Render("Completed in") +
		valueStyle.Render(rep.Duration.Round(time.Millisecond).String()))
	fmt.Println()
}
