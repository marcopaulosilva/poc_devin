package repositories

import (
	"context"

	"github.com/marcopaulosilva/poc_devin/internal/domain/entities"
)

type ChampionRepository interface {
	SaveChampions(ctx context.Context, champions []entities.ChampionRecord) error
	
	GetChampions(ctx context.Context) ([]entities.ChampionRecord, error)
	
	GetChampionByID(ctx context.Context, id string) (*entities.ChampionRecord, error)
}
