package apps

import (
	"fmt"
	"github.com/fahturr/xsis-test/apps/router"
	"github.com/fahturr/xsis-test/pkg/database"
)

const (
	ROUTER_GIN = "gin"
)

type Router interface {
	BuildRoutes()
	Run() error
}

type RouterFactory struct {
	db *database.Database
}

func NewRouterFactory(db *database.Database) *RouterFactory {
	return &RouterFactory{db}
}

func (r *RouterFactory) Create(typeRouter string, port string) (Router, error) {
	switch typeRouter {
	case ROUTER_GIN:
		return router.NewGinRouter(port, r.db), nil

	default:
		return nil, fmt.Errorf("router with type %s not found", typeRouter)
	}
}

type RouterExecutor struct{}

func NewRouterExecutor() *RouterExecutor {
	return &RouterExecutor{}
}

func (r *RouterExecutor) Execute(router Router) error {
	router.BuildRoutes()
	err := router.Run()

	return err
}
