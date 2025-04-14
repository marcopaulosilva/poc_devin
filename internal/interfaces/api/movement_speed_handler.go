package api

import (
	"context"
	"encoding/json"
	"net/http"
	"sort"

	"github.com/marcopaulosilva/poc_devin/internal/domain/usecases"
	"github.com/marcopaulosilva/poc_devin/internal/infrastructure/logger"
)

type MovementSpeedHandler struct {
	championUseCase usecases.ChampionUseCase
	logger          logger.Logger
}

func NewMovementSpeedHandler(championUseCase usecases.ChampionUseCase, logger logger.Logger) *MovementSpeedHandler {
	return &MovementSpeedHandler{
		championUseCase: championUseCase,
		logger:          logger,
	}
}

func (h *MovementSpeedHandler) GetChampionsByMovementSpeed(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("API request received: Get champions by movement speed")

	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()
	champions, err := h.championUseCase.GetAllChampions(ctx)
	if err != nil {
		h.logger.Error("Failed to get champions: %v", err)
		http.Error(w, "Failed to retrieve champion data", http.StatusInternalServerError)
		return
	}

	sort.Slice(champions, func(i, j int) bool {
		return champions[i].MovementSpeed > champions[j].MovementSpeed
	})

	type ChampionResponse struct {
		Rank         int     `json:"rank"`
		ID           string  `json:"id"`
		Name         string  `json:"name"`
		Title        string  `json:"title"`
		MovementSpeed float64 `json:"movementSpeed"`
	}

	response := struct {
		Count     int                `json:"count"`
		Champions []ChampionResponse `json:"champions"`
	}{
		Count:     len(champions),
		Champions: make([]ChampionResponse, 0, len(champions)),
	}

	for i, champion := range champions {
		response.Champions = append(response.Champions, ChampionResponse{
			Rank:         i + 1,
			ID:           champion.ID,
			Name:         champion.Name,
			Title:        champion.Title,
			MovementSpeed: champion.MovementSpeed,
		})
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		h.logger.Error("Failed to marshal response: %v", err)
		http.Error(w, "Failed to generate response", http.StatusInternalServerError)
		return
	}

	h.logger.Success("Successfully returned %d champions sorted by movement speed", len(champions))
	w.Write(jsonResponse)
}
