package models

import "time"

type Tweet struct {
	Id        int64     `json:"id"`
	UserId    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
