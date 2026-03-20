package cli

import (
	"fmt"

	"github.com/YOUR_USERNAME/YOUR_PROJECT/internal/example"
	"github.com/spf13/cobra"
)

var greetCmd = &cobra.Command{
	Use:   "greet [name]",
	Short: "Print a greeting message",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := ""
		if len(args) > 0 {
			name = args[0]
		}
		fmt.Println(example.Greet(name))
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)
}
