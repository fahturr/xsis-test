package services

import (
	"context"
	"github.com/fahturr/xsis-test/apps/commons/response"
	"github.com/fahturr/xsis-test/apps/domains/movie/params"
)

type MovieSvc interface {
	CreateMovie(ctx context.Context, req *params.CreateMovieRequest) *response.CustomError
	UpdateMovie(ctx context.Context, req *params.UpdateMovieRequest) *response.CustomError
	GetListMovie(ctx context.Context) ([]*params.GetListMovieResponse, *response.CustomError)
	GetDetailMovie(ctx context.Context, id int) (*params.GetDetailMovieResponse, *response.CustomError)
	DeleteMovie(ctx context.Context, id int) *response.CustomError
}
