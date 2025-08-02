package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sahasajib/students-api/internal/config"
	"github.com/sahasajib/students-api/internal/config/http/handlers/student"
)



func main() {
	// Load the configuration
	cfg := config.MustLoad()

	//setup router and server
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New())

	server := http.Server{
		Addr: cfg.HTTPServer.Address,
		Handler: router,
	}
	slog.Info("Starting server", slog.String("address", cfg.HTTPServer.Address))
	fmt.Println("Server is running on", cfg.HTTPServer.Address)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server: ")
		}
	}()

	<-done
	
	slog.Info("Shutting down server gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 6 * time.Second)
	defer cancel()

	err := server.Shutdown(ctx) 
		if err != nil {
			slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
		}
	slog.Info("Server stopped gracefully") 
}