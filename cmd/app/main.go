package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/marcopaulosilva/poc_devin/internal/domain/usecases"
	httpRepo "github.com/marcopaulosilva/poc_devin/internal/interfaces/http"
	"github.com/marcopaulosilva/poc_devin/internal/infrastructure/client"
	"github.com/marcopaulosilva/poc_devin/internal/infrastructure/logger"
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
	
	ctx := context.Background()
	
	log.Info("Fetching all League of Legends champions")
	champions, err := championUseCase.GetAllChampions(ctx)
	if err != nil {
		log.Error("Failed to get champions: %v", err)
		os.Exit(1)
	}
	
	sort.Slice(champions, func(i, j int) bool {
		return champions[i].Name < champions[j].Name
	})
	
	fmt.Printf("\nAll League of Legends Champions (%d):\n\n", len(champions))
	
	for i, champion := range champions {
		fmt.Printf("  %3d. %-15s - %s\n", i+1, champion.Name, champion.Title)
	}
	
	fmt.Println("\nChampion data retrieved from Riot Games Data Dragon API")
	log.Success("Application completed successfully")
}
