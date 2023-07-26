package main

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-4/cmd/server/routes"
	"github.com/extmatperez/meli_bootcamp_go_w2-4/config"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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
