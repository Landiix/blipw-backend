package repository

import (
	"context"
	"time"

	"blipw/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TweetRepository struct {
	pool *pgxpool.Pool
}

func NewTweetRepository(pool *pgxpool.Pool) *TweetRepository {
	return &TweetRepository{
		pool: pool,
	}
}

func (r *TweetRepository) GetAll(ctx context.Context) ([]models.Tweet, error) {
	query := `SELECT id, user_id,content,created_at FROM tweets`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tweets := make([]models.Tweet, 0)
	for rows.Next() {
		var t models.Tweet

		err := rows.Scan(&t.Id, &t.UserId, &t.Content, &t.CreatedAt)
		if err != nil {
			return nil, err
		}

		tweets = append(tweets, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tweets, nil
}

func (r *TweetRepository) Create(ctx context.Context, userId int64, content string) (models.Tweet, error) {
	var newId int64
	var newCreatedAt time.Time
	query := "INSERT INTO tweets (user_id, content) VALUES ($1, $2) RETURNING id, created_at"
	row := r.pool.QueryRow(ctx, query, userId, content)
	err := row.Scan(&newId, &newCreatedAt)
	if err != nil {
		return models.Tweet{}, err
	}

	var t models.Tweet
	t.Id = newId
	t.CreatedAt = newCreatedAt
	t.UserId = int(userId)
	t.Content = content

	return t, nil
}
