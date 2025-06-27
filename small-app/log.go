package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func setupLog() {

	// Configure zerolog to include caller information and use a human-friendly timestamp
	zerolog.TimeFieldFormat = time.RFC3339

	// Create a logger that writes to stdout
	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Caller(). // Include caller information (file and line)
		Logger()

	// Set as default logger
	log.Logger = logger
}
