package repository

import (
	"context"

	"blipw/internal/models"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (models.User, error) {
	var u models.User

	query := "SELECT id, username,password,created_at FROM users WHERE username = $1"
	row := r.pool.QueryRow(ctx, query, username)
	err := row.Scan(&u.Id, &u.Username, &u.Password, &u.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}

	return u, nil
}

func (r *UserRepository) Create(ctx context.Context, username string, password string) (models.User, error) {
	var u models.User

	u.Username = username
	u.Password = password

	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id, created_at"
	err := r.pool.QueryRow(ctx, query, username, password).Scan(&u.Id, &u.CreatedAt)
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}
