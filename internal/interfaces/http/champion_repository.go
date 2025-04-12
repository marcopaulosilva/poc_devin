package http

import (
	"context"
	"fmt"

	"github.com/marcopaulosilva/poc_devin/internal/domain/entities"
	"github.com/marcopaulosilva/poc_devin/internal/infrastructure/client"
	"github.com/marcopaulosilva/poc_devin/internal/infrastructure/logger"
)

type ChampionRepository struct {
	httpClient client.HTTPClient
	logger     logger.Logger
	baseURL    string
	apiKey     string
}

func NewChampionRepository(httpClient client.HTTPClient, logger logger.Logger, baseURL string, apiKey string) *ChampionRepository {
	return &ChampionRepository{
		httpClient: httpClient,
		logger:     logger,
		baseURL:    baseURL,
		apiKey:     apiKey,
	}
}

func (r *ChampionRepository) GetAllChampions(ctx context.Context) ([]entities.Champion, error) {
	url := "https://ddragon.leagueoflegends.com/cdn/15.7.1/data/en_US/champion.json"
	r.logger.Info("Fetching all champions from Data Dragon")

	data, err := r.httpClient.Get(ctx, url)
	if err != nil {
		r.logger.Error("Failed to fetch champions: %v", err)
		return nil, err
	}

	var championData entities.ChampionData
	if err := client.ParseJSON(data, &championData); err != nil {
		r.logger.Error("Failed to parse champion data: %v", err)
		return nil, err
	}

	champions := make([]entities.Champion, 0, len(championData.Data))
	for _, info := range championData.Data {
		detailURL := fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/15.7.1/data/en_US/champion/%s.json", info.ID)
		r.logger.Info("Fetching detailed data for champion: %s", info.Name)
		
		detailData, err := r.httpClient.Get(ctx, detailURL)
		if err != nil {
			r.logger.Error("Failed to fetch detailed data for champion %s: %v", info.Name, err)
			continue
		}
		
		var detailChampionData struct {
			Data map[string]struct {
				Stats struct {
					Movespeed float64 `json:"movespeed"`
				} `json:"stats"`
			} `json:"data"`
		}
		
		if err := client.ParseJSON(detailData, &detailChampionData); err != nil {
			r.logger.Error("Failed to parse detailed champion data for %s: %v", info.Name, err)
			continue
		}
		
		movespeed := 0.0
		if champData, ok := detailChampionData.Data[info.ID]; ok {
			movespeed = champData.Stats.Movespeed
		}
		
		champion := entities.Champion{
			ID:           info.ID,
			Key:          info.Key,
			Name:         info.Name,
			Title:        info.Title,
			MovementSpeed: movespeed,
		}
		champions = append(champions, champion)
	}

	r.logger.Success("Successfully fetched %d champions with movement speed data", len(champions))
	return champions, nil
}
