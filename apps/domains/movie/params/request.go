package params

import (
	"github.com/fahturr/xsis-test/apps/domains/movie/models"
)

type (
	CreateMovieRequest struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Rating      float64 `json:"rating"`
		Image       string  `json:"image"`
	}

	UpdateMovieRequest struct {
		Id          int
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Rating      float64 `json:"rating"`
		Image       string  `json:"image"`
	}
)

func (c *CreateMovieRequest) ToMovieModel() *models.Movie {
	return &models.Movie{
		Title:       c.Title,
		Description: c.Description,
		Rating:      c.Rating,
		Image:       c.Image,
	}
}

func (c *UpdateMovieRequest) ToMovieModel() *models.Movie {
	return &models.Movie{
		Id:          c.Id,
		Title:       c.Title,
		Description: c.Description,
		Rating:      c.Rating,
		Image:       c.Image,
	}
}
