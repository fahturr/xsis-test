package controllers

import (
	"github.com/fahturr/xsis-test/apps/commons/response"
	"github.com/fahturr/xsis-test/apps/domains/movie/params"
	"github.com/fahturr/xsis-test/apps/domains/movie/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

type MovieController struct {
	movieSvc services.MovieSvc
}

func NewMovieController(movieSvc services.MovieSvc) *MovieController {
	return &MovieController{movieSvc: movieSvc}
}

func (m *MovieController) GetListMovie(ctx *gin.Context) {
	movies, cErr := m.movieSvc.GetListMovie(ctx)
	if cErr != nil {
		ctx.AbortWithStatusJSON(cErr.StatusCode, cErr)
		return
	}

	res := response.GeneralSuccessCustomMessageAndPayload("success fetch list movie", movies)
	ctx.JSON(res.StatusCode, res)
}

func (m *MovieController) GetDetailMovie(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := response.BadRequestErrorWithAdditionalInfo(err.Error())
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	movies, cErr := m.movieSvc.GetDetailMovie(ctx, id)
	if cErr != nil {
		ctx.AbortWithStatusJSON(cErr.StatusCode, cErr)
		return
	}

	res := response.GeneralSuccessCustomMessageAndPayload("success fetch detail movie", movies)
	ctx.JSON(res.StatusCode, res)
}

func (m *MovieController) CreateMovie(ctx *gin.Context) {
	req := params.CreateMovieRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		res := response.BadRequestErrorWithAdditionalInfo(err.Error())
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	cErr := m.movieSvc.CreateMovie(ctx, &req)
	if cErr != nil {
		ctx.AbortWithStatusJSON(cErr.StatusCode, cErr)
		return
	}

	res := response.CreatedSuccessCustomMessagePayload("success create movie", nil)
	ctx.JSON(res.StatusCode, res)
}

func (m *MovieController) UpdateMovie(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := response.BadRequestErrorWithAdditionalInfo(err.Error())
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	req := params.UpdateMovieRequest{Id: id}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		res := response.BadRequestErrorWithAdditionalInfo(err.Error())
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	cErr := m.movieSvc.UpdateMovie(ctx, &req)
	if cErr != nil {
		ctx.AbortWithStatusJSON(cErr.StatusCode, cErr)
		return
	}

	res := response.GeneralSuccessCustomMessageAndPayload("success update movie", nil)
	ctx.JSON(res.StatusCode, res)
}

func (m *MovieController) DeleteMovie(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := response.BadRequestErrorWithAdditionalInfo(err.Error())
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	cErr := m.movieSvc.DeleteMovie(ctx, id)
	if cErr != nil {
		ctx.AbortWithStatusJSON(cErr.StatusCode, cErr)
		return
	}

	res := response.GeneralSuccessCustomMessageAndPayload("success delete movie", nil)
	ctx.JSON(res.StatusCode, res)
}
