package mysql

import (
	"context"
	"database/sql"
	"github.com/fahturr/xsis-test/apps/domains/movie/models"
	"time"
)

type MovieRepoImpl struct {
	db *sql.DB
}

func NewMovieRepo(db *sql.DB) *MovieRepoImpl {
	return &MovieRepoImpl{db: db}
}

func (m *MovieRepoImpl) CreateMovie(ctx context.Context, movie *models.Movie) error {
	queryMovie := `INSERT INTO movies(title, description, rating, image, created_at, updated_at)
				   VALUE (?, ?, ?, ?, ?, ?)`

	stmt, err := m.db.PrepareContext(ctx, queryMovie)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		movie.Title,
		movie.Description,
		movie.Rating,
		movie.Image,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	return stmt.Close()
}

func (m *MovieRepoImpl) UpdateMovie(ctx context.Context, movie *models.Movie) error {
	queryMovie := `UPDATE movies 
    			   SET title = ?, description = ?, rating = ?, image = ?, updated_at = ?
				   WHERE id = ?`

	stmt, err := m.db.PrepareContext(ctx, queryMovie)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		movie.Title,
		movie.Description,
		movie.Rating,
		movie.Image,
		time.Now(),
		movie.Id,
	)
	if err != nil {
		return err
	}

	return stmt.Close()
}

func (m *MovieRepoImpl) DeleteMovie(ctx context.Context, id int) error {
	queryMovie := `DELETE FROM movies
				   WHERE id = ?`

	stmt, err := m.db.PrepareContext(ctx, queryMovie)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return stmt.Close()
}

func (m *MovieRepoImpl) FindAllMovie(ctx context.Context) ([]*models.Movie, error) {
	queryMovie := `SELECT id, title, description, image, rating, created_at, updated_at
				   FROM movies`

	stmt, err := m.db.PrepareContext(ctx, queryMovie)
	if err != nil {
		return nil, err
	}

	result := []*models.Movie{}

	rowsMovie, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	for rowsMovie.Next() {
		tempMovie := models.Movie{}

		err = rowsMovie.Scan(
			&tempMovie.Id,
			&tempMovie.Title,
			&tempMovie.Description,
			&tempMovie.Image,
			&tempMovie.Rating,
			&tempMovie.CreatedAt,
			&tempMovie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, &tempMovie)
	}

	return result, nil
}

func (m *MovieRepoImpl) FindMovieById(ctx context.Context, id int) (*models.Movie, error) {
	queryMovie := `SELECT id, title, description, image, rating, created_at, updated_at
				   FROM movies
				   WHERE id = ?`

	stmt, err := m.db.PrepareContext(ctx, queryMovie)
	if err != nil {
		return nil, err
	}

	result := &models.Movie{}

	rowMovie := stmt.QueryRowContext(ctx, id)

	err = rowMovie.Scan(
		&result.Id,
		&result.Title,
		&result.Description,
		&result.Image,
		&result.Rating,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}
