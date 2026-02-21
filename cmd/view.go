package cmd

import (
	"fmt"
	"time"

	"github.com/naru-bot/upp/internal/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "view <name|url|id>",
		Short: "Show full configuration for a target",
		Long: `Show the full configuration and latest check for a target.

Examples:
  upp view "My Site"
  upp view https://example.com
  upp view 1`,
		Args: requireArgs(1),
		Run:  runView,
	})
}

type viewOutput struct {
	Target    db.Target       `json:"target"`
	LastCheck *db.CheckResult `json:"last_check,omitempty"`
}

func runView(cmd *cobra.Command, args []string) {
	t, err := db.GetTarget(args[0])
	if err != nil {
		exitError(err.Error())
	}

	var lastCheck *db.CheckResult
	if checks, err := db.GetCheckHistory(t.ID, 1); err == nil && len(checks) > 0 {
		lastCheck = &checks[0]
	}

	if jsonOutput {
		printJSON(viewOutput{Target: *t, LastCheck: lastCheck})
		return
	}

	fmt.Printf("Target: %s (id %d)\n", t.Name, t.ID)
	fmt.Printf("URL: %s\n", t.URL)
	fmt.Printf("Type: %s\n", t.Type)
	fmt.Printf("Interval: %ds\n", t.Interval)
	fmt.Printf("Timeout: %ds\n", t.Timeout)
	fmt.Printf("Retries: %d\n", t.Retries)
	fmt.Printf("Paused: %v\n", t.Paused)
	fmt.Printf("Created: %s\n", t.CreatedAt.Format(time.RFC3339))

	if t.Selector != "" {
		fmt.Printf("Selector: %s\n", t.Selector)
	}
	if t.Headers != "" {
		fmt.Printf("Headers: %s\n", t.Headers)
	}
	if t.Expect != "" {
		fmt.Printf("Expect: %s\n", t.Expect)
	}
	if t.Threshold > 0 {
		fmt.Printf("Threshold: %.1f%%\n", t.Threshold)
	}

	if lastCheck == nil {
		fmt.Println("Last check: none (run 'upp check')")
		return
	}

	fmt.Printf("Last check: %s\n", lastCheck.CheckedAt.Format(time.RFC3339))
	fmt.Printf("Status: %s\n", lastCheck.Status)
	if lastCheck.StatusCode != 0 {
		fmt.Printf("Status code: %d\n", lastCheck.StatusCode)
	}
	if lastCheck.ResponseTime != 0 {
		fmt.Printf("Response time: %dms\n", lastCheck.ResponseTime)
	}
	if lastCheck.Error != "" {
		fmt.Printf("Error: %s\n", lastCheck.Error)
	}
}
