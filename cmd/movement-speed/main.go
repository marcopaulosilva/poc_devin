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
	
	log.Info("Fetching all League of Legends champions with movement speed data")
	champions, err := championUseCase.GetAllChampions(ctx)
	if err != nil {
		log.Error("Failed to get champions: %v", err)
		os.Exit(1)
	}
	
	sort.Slice(champions, func(i, j int) bool {
		return champions[i].MovementSpeed > champions[j].MovementSpeed
	})
	
	fmt.Printf("\nLeague of Legends Champions Movement Speed (%d):\n\n", len(champions))
	
	fmt.Printf("  %-4s %-15s %-25s %s\n", "Rank", "Name", "Title", "Movement Speed")
	fmt.Printf("  %-4s %-15s %-25s %s\n", "----", "---------------", "-------------------------", "-------------")
	
	for i, champion := range champions {
		fmt.Printf("  %-4d %-15s %-25s %.0f\n", i+1, champion.Name, champion.Title, champion.MovementSpeed)
	}
	
	fmt.Println("\nChampion movement speed data retrieved from Riot Games Data Dragon API")
	log.Success("Movement speed microservice completed successfully")
}
