package usecases

import (
	"context"

	"github.com/marcopaulosilva/poc_devin/internal/domain/entities"
)

type ChampionUseCase interface {
	GetAllChampions(ctx context.Context) ([]entities.Champion, error)
}

type ChampionUseCaseImpl struct {
	championRepo ChampionRepository
}

type ChampionRepository interface {
	GetAllChampions(ctx context.Context) ([]entities.Champion, error)
}

func NewChampionUseCase(championRepo ChampionRepository) ChampionUseCase {
	return &ChampionUseCaseImpl{
		championRepo: championRepo,
	}
}

func (uc *ChampionUseCaseImpl) GetAllChampions(ctx context.Context) ([]entities.Champion, error) {
	return uc.championRepo.GetAllChampions(ctx)
}
