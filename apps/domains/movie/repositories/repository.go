package repositories

import (
	"context"
	"github.com/fahturr/xsis-test/apps/domains/movie/models"
)

type MovieRepo interface {
	CreateMovie(ctx context.Context, movie *models.Movie) error
	UpdateMovie(ctx context.Context, movie *models.Movie) error
	DeleteMovie(ctx context.Context, id int) error
	FindAllMovie(ctx context.Context) ([]*models.Movie, error)
	FindMovieById(ctx context.Context, id int) (*models.Movie, error)
}
