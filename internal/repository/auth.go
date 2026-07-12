package repository

import (
	"database/sql"
)

type AuthRepository struct {
	db *sql.DB
}
