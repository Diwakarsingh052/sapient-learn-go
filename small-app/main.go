package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"os/signal"
	"small-app/internal/handlers"
	"small-app/internal/models"
	"time"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: Error loading .env file:", err)
	}
	err := startApp()
	if err != nil {
		panic(err)
	}

}

func startApp() error {
	apiPort := os.Getenv("API_PORT")
	//m := http.NewServeMux()

	// Initializing Cache Support
	c := models.NewConn()

	h := handlers.API(c)
	api := http.Server{
		Addr:         ":" + apiPort,
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		IdleTimeout:  800 * time.Second,
		Handler:      h,
	}

	serverError := make(chan error)
	go func() {
		err := api.ListenAndServe()
		serverError <- err
	}()
	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, os.Kill)

	select {
	case err := <-serverError:
		return err
	case <-shutdown:
		fmt.Println("Graceful Shut down...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		//Shutdown gracefully shuts down the server without interrupting any active connections.
		//Shutdown works by first closing all open listeners, then closing all idle connections,
		err := api.Shutdown(ctx)
		if err != nil {
			// forceful shutdown
			err := api.Close()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
