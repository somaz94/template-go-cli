package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	output  string
	noColor bool
)

var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "A brief description of your CLI tool",
	Long:  "mycli — A longer description explaining what your CLI tool does and its main features.",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "color", "output format: color, plain, json")
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "disable color output")
}

// Execute runs the root command.
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	return nil
}
