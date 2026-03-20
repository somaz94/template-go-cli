package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Set via ldflags at build time.
	Version   = "dev"
	GitCommit = "none"
	BuildDate = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("mycli %s (commit: %s, built: %s)\n", Version, GitCommit, BuildDate)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
