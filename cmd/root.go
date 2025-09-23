package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// Version is set during build
	Version = "0.0.1"
	// logger is the global logger instance
	logger *logrus.Logger
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "testnlearn",
	Short: "A CLI application for testing and learning",
	Long:  `A CLI application built with Cobra that demonstrates version and run commands with configurable sleep duration.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Initialize logger
	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.InfoLevel)
}
