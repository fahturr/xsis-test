package router

import (
	"github.com/fahturr/xsis-test/apps/domains/movie"
	"github.com/fahturr/xsis-test/pkg/database"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	port   string
	router *gin.Engine
	db     *database.Database
}

func NewGinRouter(port string, db *database.Database) *Gin {
	router := gin.Default()

	return &Gin{
		port:   port,
		router: router,
		db:     db,
	}
}

func (r *Gin) BuildRoutes() {
	r.router.Use(CORS)

	v1 := r.router.Group("v1")

	movieRoute := movie.NewRouterMovie(v1, r.db)
	movieRoute.RegisterAuthRoutes()
}

func (r *Gin) Run() error {
	err := r.router.Run(r.port)

	return err
}

func CORS(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	ctx.Next()
}
