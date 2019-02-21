package cmd

import (
	"github.com/MSLacerda/status-hub/internal/app"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs status-hub server",
	Run: func(cmd *cobra.Command, args []string) {
		app.BuildApp()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
