package http

import (
	"context"
	"fmt"

	"github.com/marcopaulosilva/poc_devin/internal/domain/entities"
	"github.com/marcopaulosilva/poc_devin/internal/infrastructure/client"
	"github.com/marcopaulosilva/poc_devin/internal/infrastructure/logger"
)

type UserRepository struct {
	httpClient client.HTTPClient
	logger     logger.Logger
	baseURL    string
}

func NewUserRepository(httpClient client.HTTPClient, logger logger.Logger, baseURL string) *UserRepository {
	return &UserRepository{
		httpClient: httpClient,
		logger:     logger,
		baseURL:    baseURL,
	}
}

func (r *UserRepository) GetUser(ctx context.Context, id string) (*entities.User, error) {
	url := fmt.Sprintf("%s/users/%s", r.baseURL, id)
	r.logger.Info("Fetching user with ID: %s", id)

	data, err := r.httpClient.Get(ctx, url)
	if err != nil {
		r.logger.Error("Failed to fetch user: %v", err)
		return nil, err
	}

	var user entities.User
	if err := client.ParseJSON(data, &user); err != nil {
		r.logger.Error("Failed to parse user data: %v", err)
		return nil, err
	}

	r.logger.Success("Successfully fetched user: %s", user.Name)
	return &user, nil
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]entities.User, error) {
	url := fmt.Sprintf("%s/users", r.baseURL)
	r.logger.Info("Fetching all users")

	data, err := r.httpClient.Get(ctx, url)
	if err != nil {
		r.logger.Error("Failed to fetch users: %v", err)
		return nil, err
	}

	var users []entities.User
	if err := client.ParseJSON(data, &users); err != nil {
		r.logger.Error("Failed to parse users data: %v", err)
		return nil, err
	}

	r.logger.Success("Successfully fetched %d users", len(users))
	return users, nil
}
