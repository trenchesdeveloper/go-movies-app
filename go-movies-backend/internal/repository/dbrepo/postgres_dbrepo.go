package dbrepo

import (
	"context"
	"database/sql"
	"time"

	"github.com/trenchesdeveloper/go-backend/internal/models"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT id, title, release_date, runtime, mpaa_rating, coalesce(image, ''), created_at, updated_at
		FROM movies
		ORDER BY title
	`

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var movies []*models.Movie

	for rows.Next() {
		var movie models.Movie

		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.MPAARating,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		movies = append(movies, &movie)

	}

	return movies, nil
}



