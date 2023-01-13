package movie

import (
	"github.com/fahturr/xsis-test/apps/domains/movie/controllers"
	"github.com/fahturr/xsis-test/apps/domains/movie/repositories/mysql"
	"github.com/fahturr/xsis-test/apps/domains/movie/services/usecase"
	"github.com/fahturr/xsis-test/pkg/database"
	"github.com/gin-gonic/gin"
)

type RouteMovie interface {
	RegisterAuthRoutes()
}

type RouteMovieImpl struct {
	route *gin.RouterGroup
	movie *controllers.MovieController
}

func NewRouterMovie(r *gin.RouterGroup, db *database.Database) *RouteMovieImpl {
	repo := mysql.NewMovieRepo(db.DbSql)
	svc := usecase.NewMovieSvc(repo)
	handler := controllers.NewMovieController(svc)

	return &RouteMovieImpl{
		route: r,
		movie: handler,
	}
}

func (r *RouteMovieImpl) RegisterAuthRoutes() {
	movie := r.route.Group("movies")
	{
		movie.GET("list", r.movie.GetListMovie)
		movie.GET("detail/:id", r.movie.GetDetailMovie)
		movie.DELETE("delete/:id", r.movie.DeleteMovie)
		movie.POST("create", r.movie.CreateMovie)
		movie.PUT("update/:id", r.movie.UpdateMovie)
	}
}
