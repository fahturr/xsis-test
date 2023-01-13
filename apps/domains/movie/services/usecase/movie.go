package usecase

import (
	"context"
	"github.com/fahturr/xsis-test/apps/commons/response"
	"github.com/fahturr/xsis-test/apps/domains/movie/params"
	"github.com/fahturr/xsis-test/apps/domains/movie/repositories"
)

type MovieSvcImpl struct {
	movieRepo repositories.MovieRepo
}

func NewMovieSvc(movieRepo repositories.MovieRepo) *MovieSvcImpl {
	return &MovieSvcImpl{movieRepo: movieRepo}
}

func (m *MovieSvcImpl) CreateMovie(ctx context.Context, req *params.CreateMovieRequest) *response.CustomError {
	reqMovie := req.ToMovieModel()

	err := m.movieRepo.CreateMovie(ctx, reqMovie)
	if err != nil {
		return response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	return nil
}

func (m *MovieSvcImpl) UpdateMovie(ctx context.Context, req *params.UpdateMovieRequest) *response.CustomError {
	reqMovie := req.ToMovieModel()

	err := m.movieRepo.UpdateMovie(ctx, reqMovie)
	if err != nil {
		return response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	return nil
}

func (m *MovieSvcImpl) GetListMovie(ctx context.Context) ([]*params.GetListMovieResponse, *response.CustomError) {
	result, err := m.movieRepo.FindAllMovie(ctx)
	if err != nil {
		return nil, response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	res := []*params.GetListMovieResponse{}

	for _, movie := range result {
		tempMovie := params.GetListMovieResponse{
			Id:          movie.Id,
			Title:       movie.Title,
			Description: movie.Description,
			Rating:      movie.Rating,
			Image:       movie.Image,
			CreatedAt:   movie.CreatedAt,
			UpdatedAt:   movie.UpdatedAt,
		}

		res = append(res, &tempMovie)
	}

	return res, nil
}

func (m *MovieSvcImpl) GetDetailMovie(ctx context.Context, id int) (*params.GetDetailMovieResponse, *response.CustomError) {
	result, err := m.movieRepo.FindMovieById(ctx, id)
	if err != nil {
		return nil, response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	res := &params.GetDetailMovieResponse{
		Id:          result.Id,
		Title:       result.Title,
		Description: result.Description,
		Rating:      result.Rating,
		Image:       result.Image,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}

	return res, nil
}

func (m *MovieSvcImpl) DeleteMovie(ctx context.Context, id int) *response.CustomError {
	err := m.movieRepo.DeleteMovie(ctx, id)
	if err != nil {
		return response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	return nil
}
