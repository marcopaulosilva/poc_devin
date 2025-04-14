package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/marcopaulosilva/poc_devin/internal/domain/usecases"
	"github.com/marcopaulosilva/poc_devin/internal/infrastructure/client"
	"github.com/marcopaulosilva/poc_devin/internal/infrastructure/logger"
	"github.com/marcopaulosilva/poc_devin/internal/interfaces/api"
	httpRepo "github.com/marcopaulosilva/poc_devin/internal/interfaces/http"
)

func main() {
	log := logger.NewConsoleLogger()
	httpClient := client.NewHTTPClient(10 * time.Second)
	
	apiKey := os.Getenv("RIOT_API_KEY")
	if apiKey == "" {
		log.Error("RIOT_API_KEY environment variable not set")
		os.Exit(1)
	}
	
	baseURL := "https://na1.api.riotgames.com"
	
	championRepo := httpRepo.NewChampionRepository(httpClient, log, baseURL, apiKey)
	championUseCase := usecases.NewChampionUseCase(championRepo)
	
	movementSpeedHandler := api.NewMovementSpeedHandler(championUseCase, log)
	
	mux := http.NewServeMux()
	mux.HandleFunc("/api/champions/movement-speed", movementSpeedHandler.GetChampionsByMovementSpeed)
	
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	
	port := 8080
	server := api.NewServer(port, mux, log)
	
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	
	go func() {
		log.Info("Starting REST API server on port %d", port)
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Error("Failed to start server: %v", err)
			os.Exit(1)
		}
	}()
	
	<-stop
	
	log.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := server.Shutdown(ctx); err != nil {
		log.Error("Server shutdown failed: %v", err)
	}
	
	log.Success("Server gracefully stopped")
}
