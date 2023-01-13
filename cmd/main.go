package main

import (
	"fmt"
	"github.com/fahturr/xsis-test/apps"
	"github.com/fahturr/xsis-test/config"
	"github.com/fahturr/xsis-test/pkg/database"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	database.VerifyDBConnection()

	sqlDb, err := database.ConnectSQLDB()
	if err != nil {
		panic(err)
	}

	db := database.NewDatabase()
	db.SetDbSql(sqlDb)

	port := os.Getenv(config.APP_PORT)
	factory := apps.NewRouterFactory(db)
	router, err := factory.Create(apps.ROUTER_GIN, fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	executor := apps.NewRouterExecutor()
	err = executor.Execute(router)
	if err != nil {
		panic(err)
	}
}
