package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mercadolibre/fury_bootcamp-go-w2-s4-4-4/cmd/api/routes"
	"github.com/mercadolibre/fury_bootcamp-go-w2-s4-4-4/config"
)

func main() {
	// NO MODIFICAR
	db := config.ConnectDb("mysql")

	eng := gin.Default()
	router := routes.NewRouter(eng, db)
	router.MapRoutes()

	if err := eng.Run(); err != nil {
		panic(err)
	}
}
