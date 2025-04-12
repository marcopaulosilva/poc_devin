package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/marcopaulosilva/poc_devin/internal/domain/usecases"
	httpRepo "github.com/marcopaulosilva/poc_devin/internal/interfaces/http"
	"github.com/marcopaulosilva/poc_devin/internal/infrastructure/client"
	"github.com/marcopaulosilva/poc_devin/internal/infrastructure/logger"
)

func main() {
	log := logger.NewConsoleLogger()
	httpClient := client.NewHTTPClient(10 * time.Second)
	
	baseURL := "https://jsonplaceholder.typicode.com"
	
	userRepo := httpRepo.NewUserRepository(httpClient, log, baseURL)
	
	userUseCase := usecases.NewUserUseCase(userRepo)
	
	ctx := context.Background()
	
	log.Info("Example 1: Fetching a specific user")
	userId := "1" // Example user ID
	user, err := userUseCase.GetUser(ctx, userId)
	if err != nil {
		log.Error("Failed to get user: %v", err)
		os.Exit(1)
	}
	fmt.Printf("\nUser details:\n")
	fmt.Printf("  ID: %s\n", user.ID)
	fmt.Printf("  Name: %s\n\n", user.Name)
	
	log.Info("Example 2: Fetching all users")
	users, err := userUseCase.GetUsers(ctx)
	if err != nil {
		log.Error("Failed to get users: %v", err)
		os.Exit(1)
	}
	
	fmt.Printf("\nAll users (%d):\n", len(users))
	for i, u := range users {
		if i < 5 { // Only print first 5 users to avoid cluttering the console
			fmt.Printf("  %d. %s (ID: %s)\n", i+1, u.Name, u.ID)
		}
	}
	
	if len(users) > 5 {
		fmt.Printf("  ... and %d more users\n", len(users)-5)
	}
	fmt.Println()
	
	log.Success("Application completed successfully")
}
