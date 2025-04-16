package entities

import (
	"time"
)

type ChampionRecord struct {
	ID            int       `json:"id"`
	ChampionID    string    `json:"champion_id"`
	Name          string    `json:"name"`
	Title         string    `json:"title"`
	MovementSpeed float64   `json:"movement_speed"`
	Rank          int       `json:"rank"`
	CreatedAt     time.Time `json:"created_at"`
}
