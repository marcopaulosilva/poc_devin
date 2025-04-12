package http

import (
	"context"

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
		champion := entities.Champion{
			ID:    info.ID,
			Key:   info.Key,
			Name:  info.Name,
			Title: info.Title,
		}
		champions = append(champions, champion)
	}

	r.logger.Success("Successfully fetched %d champions", len(champions))
	return champions, nil
}
