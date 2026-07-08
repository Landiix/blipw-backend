package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Tweet struct {
	Id        int
	UserId    int
	Content   string
	CreatedAt time.Time
}

type TweetRepository struct {
	pool *pgxpool.Pool
}

func NewTweetRepository(pool *pgxpool.Pool) *TweetRepository {
	return &TweetRepository{
		pool: pool,
	}
}

func (r *TweetRepository) GetAll(ctx context.Context) ([]Tweet, error) {
	query := `SELECT id, user_id,content,created_at FROM tweets`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tweets := make([]Tweet, 0)
	for rows.Next() {
		var t Tweet

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
