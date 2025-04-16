package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/marcopaulosilva/poc_devin/internal/domain/repositories"
	"github.com/marcopaulosilva/poc_devin/internal/infrastructure/client"
	dbInfra "github.com/marcopaulosilva/poc_devin/internal/infrastructure/db"
	"github.com/marcopaulosilva/poc_devin/internal/infrastructure/logger"
	"github.com/marcopaulosilva/poc_devin/internal/interfaces/api"
	"github.com/marcopaulosilva/poc_devin/internal/interfaces/db"
)

func main() {
	log := logger.NewConsoleLogger()
	log.Info("Starting champion movement speed consumer application")

	apiBaseURL := getEnvOrDefault("API_BASE_URL", "http://movement-speed-api")
	dbHost := getEnvOrDefault("DB_HOST", "localhost")
	dbPort := getEnvOrDefault("DB_PORT", "5432")
	dbUser := getEnvOrDefault("DB_USER", "postgres")
	dbPassword := getEnvOrDefault("DB_PASSWORD", "postgres")
	dbName := getEnvOrDefault("DB_NAME", "champions")
	dbSSLMode := getEnvOrDefault("DB_SSL_MODE", "disable")
	syncInterval := getEnvOrDefault("SYNC_INTERVAL", "60") // seconds

	interval, err := time.ParseDuration(syncInterval + "s")
	if err != nil {
		log.Error("Invalid sync interval: %v", err)
		os.Exit(1)
	}

	dbConfig := dbInfra.PostgresConfig{
		Host:     dbHost,
		Port:     parsePort(dbPort),
		User:     dbUser,
		Password: dbPassword,
		DBName:   dbName,
		SSLMode:  dbSSLMode,
	}

	dbConn, err := dbInfra.NewPostgresConnection(dbConfig, log)
	if err != nil {
		log.Error("Failed to connect to database: %v", err)
		os.Exit(1)
	}
	defer dbConn.Close()

	if err := dbInfra.InitializeDatabase(dbConn, log); err != nil {
		log.Error("Failed to initialize database: %v", err)
		os.Exit(1)
	}

	httpClient := client.NewHTTPClient(10 * time.Second)

	apiClient := api.NewMovementSpeedClient(httpClient, apiBaseURL, log)

	championRepo := db.NewPostgresChampionRepository(dbConn, log)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go syncChampionsData(ctx, apiClient, championRepo, interval, log)

	<-stop
	log.Info("Shutting down...")
}

func syncChampionsData(
	ctx context.Context,
	apiClient *api.MovementSpeedClient,
	repo repositories.ChampionRepository,
	interval time.Duration,
	log logger.Logger,
) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	if err := performSync(ctx, apiClient, repo, log); err != nil {
		log.Error("Initial sync failed: %v", err)
	}

	for {
		select {
		case <-ticker.C:
			if err := performSync(ctx, apiClient, repo, log); err != nil {
				log.Error("Sync failed: %v", err)
			}
		case <-ctx.Done():
			log.Info("Sync process terminated")
			return
		}
	}
}

func performSync(
	ctx context.Context,
	apiClient *api.MovementSpeedClient,
	repo repositories.ChampionRepository,
	log logger.Logger,
) error {
	log.Info("Syncing champion data from API to database")

	champions, err := apiClient.GetChampionsByMovementSpeed(ctx)
	if err != nil {
		return err
	}

	if err := repo.SaveChampions(ctx, champions); err != nil {
		return err
	}

	log.Success("Successfully synced %d champions to database", len(champions))
	return nil
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func parsePort(port string) int {
	var p int
	_, err := fmt.Sscanf(port, "%d", &p)
	if err != nil || p <= 0 || p > 65535 {
		return 5432 // Default PostgreSQL port
	}
	return p
}
