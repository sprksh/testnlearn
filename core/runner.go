package core

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Run executes the main application logic
// It reads the sleep duration from environment variables and sleeps for that duration
func Run(logger *logrus.Logger) error {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		logger.WithError(err).Warn("Could not load .env file, using default values")
	}

	// Get sleep duration from environment variable
	sleepDurationStr := os.Getenv("SLEEP_DURATION")
	if sleepDurationStr == "" {
		sleepDurationStr = "5" // Default to 5 seconds
	}

	sleepDuration, err := strconv.Atoi(sleepDurationStr)
	if err != nil {
		logger.WithError(err).Error("Invalid SLEEP_DURATION value, using default")
		sleepDuration = 5
	}

	logger.WithField("duration", sleepDuration).Info("Starting sleep")

	// Sleep for the configured duration
	time.Sleep(time.Duration(sleepDuration) * time.Second)

	logger.Info("Sleep completed successfully")
	return nil
}
