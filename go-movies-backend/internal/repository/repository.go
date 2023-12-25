package repository

import (
	"database/sql"

	"github.com/trenchesdeveloper/go-backend/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
