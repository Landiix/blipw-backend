package database

import (
	"blipw/internal/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(cfg *config.Config) (*pgxpool.Pool, error) {
	//postgres://Пользователь:Пароль@Хост:Порт/ИмяБазы?sslmode=disable
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, fmt.Errorf("Unable to creatr connection pool: %w", err)
	}
	err = pool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Unable to ping db: %w", err)
	}

	return pool, nil
}
