package cmd

import (
	"testnlearn/core"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the application with configurable sleep duration",
	Long:  `Run the application which sleeps for a configurable duration and then returns success.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Starting run command")

		// Call the core layer
		err := core.Run(logger)
		if err != nil {
			logger.WithError(err).Error("Run command failed")
			return
		}

		logger.Info("Run command completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
