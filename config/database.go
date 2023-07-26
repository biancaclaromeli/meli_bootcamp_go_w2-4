package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mercadolibre/go-meli-toolkit/gomelipass"
)

func ConnectDb(database string) *sql.DB {
	dbUsername := "bgow2s444_RPROD"
	dbPassword := gomelipass.GetEnv("DB_MYSQL_DESAENV08_BGOW2S444_BGOW2S444_RPROD")
	dbHost := gomelipass.GetEnv("DB_MYSQL_DESAENV08_BGOW2S444_BGOW2S444_ENDPOINT")
	dbName := "bgow2s444"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	databaseURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbUsername, dbPassword, dbHost, dbName)

	db, err := sql.Open(database, databaseURL)

	if err != nil {
		log.Fatal(database+"failed to start", err)
	}

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("could not ping the database", err)
	}

	return db
}
