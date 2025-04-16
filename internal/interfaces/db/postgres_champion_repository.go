package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/marcopaulosilva/poc_devin/internal/domain/entities"
	"github.com/marcopaulosilva/poc_devin/internal/domain/repositories"
	"github.com/marcopaulosilva/poc_devin/internal/infrastructure/logger"
)

type PostgresChampionRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewPostgresChampionRepository(db *sql.DB, logger logger.Logger) repositories.ChampionRepository {
	return &PostgresChampionRepository{
		db:     db,
		logger: logger,
	}
}

func (r *PostgresChampionRepository) SaveChampions(ctx context.Context, champions []entities.ChampionRecord) error {
	r.logger.Info("Saving %d champions to database", len(champions))

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		r.logger.Error("Failed to begin transaction: %v", err)
		return err
	}

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO champions (champion_id, name, title, movement_speed, rank, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (champion_id) 
		DO UPDATE SET name = $2, title = $3, movement_speed = $4, rank = $5, created_at = $6
	`)
	if err != nil {
		tx.Rollback()
		r.logger.Error("Failed to prepare statement: %v", err)
		return err
	}
	defer stmt.Close()

	for _, champion := range champions {
		_, err := stmt.ExecContext(
			ctx,
			champion.ChampionID,
			champion.Name,
			champion.Title,
			champion.MovementSpeed,
			champion.Rank,
			time.Now(),
		)
		if err != nil {
			tx.Rollback()
			r.logger.Error("Failed to insert champion %s: %v", champion.Name, err)
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		r.logger.Error("Failed to commit transaction: %v", err)
		return err
	}

	r.logger.Success("Successfully saved %d champions to database", len(champions))
	return nil
}

func (r *PostgresChampionRepository) GetChampions(ctx context.Context) ([]entities.ChampionRecord, error) {
	r.logger.Info("Retrieving champions from database")

	rows, err := r.db.QueryContext(ctx, `
		SELECT id, champion_id, name, title, movement_speed, rank, created_at
		FROM champions
		ORDER BY rank ASC
	`)
	if err != nil {
		r.logger.Error("Failed to query champions: %v", err)
		return nil, err
	}
	defer rows.Close()

	var champions []entities.ChampionRecord
	for rows.Next() {
		var champion entities.ChampionRecord
		if err := rows.Scan(
			&champion.ID,
			&champion.ChampionID,
			&champion.Name,
			&champion.Title,
			&champion.MovementSpeed,
			&champion.Rank,
			&champion.CreatedAt,
		); err != nil {
			r.logger.Error("Failed to scan champion row: %v", err)
			return nil, err
		}
		champions = append(champions, champion)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error iterating champion rows: %v", err)
		return nil, err
	}

	r.logger.Success("Successfully retrieved %d champions from database", len(champions))
	return champions, nil
}

func (r *PostgresChampionRepository) GetChampionByID(ctx context.Context, id string) (*entities.ChampionRecord, error) {
	r.logger.Info("Retrieving champion with ID %s from database", id)

	var champion entities.ChampionRecord
	err := r.db.QueryRowContext(ctx, `
		SELECT id, champion_id, name, title, movement_speed, rank, created_at
		FROM champions
		WHERE champion_id = $1
	`, id).Scan(
		&champion.ID,
		&champion.ChampionID,
		&champion.Name,
		&champion.Title,
		&champion.MovementSpeed,
		&champion.Rank,
		&champion.CreatedAt,
	)

	if err == sql.ErrNoRows {
		r.logger.Info("Champion with ID %s not found", id)
		return nil, nil
	} else if err != nil {
		r.logger.Error("Failed to query champion by ID: %v", err)
		return nil, err
	}

	r.logger.Success("Successfully retrieved champion %s from database", champion.Name)
	return &champion, nil
}
