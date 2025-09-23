# Testnlearn CLI Application

A simple CLI application built with Go and Cobra that demonstrates version and run commands with configurable sleep duration.

## Features

- **Version Command**: Displays the application version
- **Run Command**: Sleeps for a configurable duration and returns success
- **Environment Configuration**: Sleep duration is configurable via `.env` file
- **Logging**: Uses structured logging with logrus
- **Two-Layer Architecture**: 
  - `cmd` layer: CLI command handling
  - `core` layer: Business logic

## Project Structure

```
testnlearn/
├── cmd/           # CLI command layer
│   ├── root.go    # Root command and logger setup
│   ├── version.go # Version command
│   └── run.go     # Run command
├── core/          # Business logic layer
│   └── runner.go  # Core run functionality
├── main.go        # Application entry point
├── .env           # Environment configuration
└── go.mod         # Go module dependencies
```

## Usage

### Build the application
```bash
go mod tidy
go build -o testnlearn .
```

### Run commands
```bash
# Show version
./testnlearn version

# Run with default sleep duration (5 seconds)
./testnlearn run

# Show help
./testnlearn --help
```

### Configuration

The sleep duration can be configured in the `.env` file:
```
SLEEP_DURATION=3
```

If no `.env` file exists or `SLEEP_DURATION` is not set, it defaults to 5 seconds.

## Dependencies

- [Cobra](https://github.com/spf13/cobra): CLI framework
- [Logrus](https://github.com/sirupsen/logrus): Structured logging
- [Godotenv](https://github.com/joho/godotenv): Environment variable loading
