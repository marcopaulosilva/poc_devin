package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marcopaulosilva/poc_devin/internal/domain/entities"
	"github.com/marcopaulosilva/poc_devin/internal/infrastructure/client"
	"github.com/marcopaulosilva/poc_devin/internal/infrastructure/logger"
)

type MovementSpeedClient struct {
	httpClient client.HTTPClient
	baseURL    string
	logger     logger.Logger
}

func NewMovementSpeedClient(httpClient client.HTTPClient, baseURL string, logger logger.Logger) *MovementSpeedClient {
	return &MovementSpeedClient{
		httpClient: httpClient,
		baseURL:    baseURL,
		logger:     logger,
	}
}

func (c *MovementSpeedClient) GetChampionsByMovementSpeed(ctx context.Context) ([]entities.ChampionRecord, error) {
	c.logger.Info("Fetching champions by movement speed from API")

	url := fmt.Sprintf("%s/api/champions/movement-speed", c.baseURL)
	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		c.logger.Error("Failed to create request: %v", err)
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("Failed to send request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.logger.Error("API returned non-OK status: %d", resp.StatusCode)
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var apiResponse struct {
		Count     int `json:"count"`
		Champions []struct {
			Rank          int     `json:"rank"`
			ID            string  `json:"id"`
			Name          string  `json:"name"`
			Title         string  `json:"title"`
			MovementSpeed float64 `json:"movementSpeed"`
		} `json:"champions"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		c.logger.Error("Failed to decode API response: %v", err)
		return nil, err
	}

	champions := make([]entities.ChampionRecord, 0, len(apiResponse.Champions))
	for _, c := range apiResponse.Champions {
		champions = append(champions, entities.ChampionRecord{
			ChampionID:    c.ID,
			Name:          c.Name,
			Title:         c.Title,
			MovementSpeed: c.MovementSpeed,
			Rank:          c.Rank,
		})
	}

	c.logger.Success("Successfully fetched %d champions from API", len(champions))
	return champions, nil
}
