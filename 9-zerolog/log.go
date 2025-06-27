package main

import (
	"errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	configureLogger()
	//moreLoggingExamples()
	contextLogger()
}

func configureLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// setting the config as global, anywhere we use zerolog logger, below config would be used
	log.Logger = zerolog.New(os.Stdout).
		With().
		Timestamp().
		Caller().                     // file path and line number
		Str("service", "my-service"). // Adding service name to all logs helps with filtering
		Str("version", "1.0.0").
		Logger()

	log.Info().Msg("logger configured with JSON output")
	log.Info().Str("name", "john").Send() // we must use msg or send to output the logs

}

type User struct {
	ID       string
	Username string
	Email    string
	Role     string
}

func moreLoggingExamples() {

	log.Warn().Msg("Resource usage is high")

	log.Error().
		Str("component", "database").
		Str("operation", "query").
		Err(errors.New("connection timeout")).
		Msg("Database operation failed")

	// Create a user for our examples
	user := User{
		ID:       "usr_123456",
		Username: "johndoe",
		Email:    "john.doe@example.com",
		Role:     "admin",
	}

	// Log a user login event with structured data
	log.Info().
		Str("event", "user_login").
		Any("user struct", user).
		Str("username", user.Username).
		Str("ip_address", "192.168.1.1").
		Msg("User logged in successfully")

	// Log with sub-objects using Dict
	log.Info().
		Str("event", "user_updated").
		Dict("user", zerolog.Dict().
			Str("id", user.ID).
			Str("username", user.Username).
			Str("email", user.Email).
			Str("role", user.Role)).
		Dict("changes", zerolog.Dict().
			Str("old_email", "old.email@example.com").
			Str("new_email", user.Email)).
		Msg("User profile updated")

}

func contextLogger() {
	// Create a user for our examples
	user := User{
		ID:       "usr_123456",
		Username: "johndoe",
		Email:    "john.doe@example.com",
		Role:     "admin",
	}
	// Create a context logger with common fields
	contextLogger := log.With().
		Str("request_id", "req_567890").
		Str("user_id", user.ID).
		Logger()

	// Use the context logger for related operations
	contextLogger.Info().Msg("Request processing started")

}
