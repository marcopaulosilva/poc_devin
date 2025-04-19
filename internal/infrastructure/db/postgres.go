package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/marcopaulosilva/poc_devin/internal/infrastructure/logger"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresConnection(config PostgresConfig, logger logger.Logger) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	logger.Info("Connecting to PostgreSQL database at %s:%d", config.Host, config.Port)
	
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Failed to open database connection: %v", err)
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		logger.Error("Failed to ping database: %v", err)
		return nil, err
	}

	logger.Success("Successfully connected to PostgreSQL database")
	return db, nil
}

func InitializeDatabase(db *sql.DB, logger logger.Logger) error {
	logger.Info("Initializing database schema")

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS champions (
			id SERIAL PRIMARY KEY,
			champion_id VARCHAR(50) UNIQUE NOT NULL,
			name VARCHAR(100) NOT NULL,
			title VARCHAR(200) NOT NULL,
			movement_speed FLOAT NOT NULL,
			rank INT NOT NULL,
			created_at TIMESTAMP NOT NULL
		)
	`)
	if err != nil {
		logger.Error("Failed to create champions table: %v", err)
		return err
	}

	_, err = db.Exec(`
		CREATE INDEX IF NOT EXISTS idx_champions_champion_id ON champions(champion_id)
	`)
	if err != nil {
		logger.Error("Failed to create index on champion_id: %v", err)
		return err
	}

	logger.Success("Database schema initialized successfully")
	return nil
}
