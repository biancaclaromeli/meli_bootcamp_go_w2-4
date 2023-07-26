package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mercadolibre/go-meli-toolkit/gomelipass"
)

// Library go-meli-toolkit should be inside of vendor folder.
// E.g: app-name/src/api/vendor/go-meli-toolkit
// Import client:

var dbUsername = "bgow2s444_RPROD"
var dbPassword = gomelipass.GetEnv("DB_MYSQL_DESAENV08_BGOW2S444_BGOW2S444_RPROD")

// You can use the variables (with READ ONLY permissions):
// gomelipass.GetEnv("DB_MYSQL_DESAENV08_BGOW2S444_BGOW2S444_LOCAL_REPLICA_ENDPOINT"; to connect to the local replica, or...
// gomelipass.GetEnv("DB_MYSQL_DESAENV08_BGOW2S444_BGOW2S444_REMOTE_REPLICA_ENDPOINT"; for the remote replica
var dbHost = gomelipass.GetEnv("DB_MYSQL_DESAENV08_BGOW2S444_BGOW2S444_ENDPOINT")
var dbName = "bgow2s444"

func ConnectDb(database string) *sql.DB {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// load environment variables
	// err := godotenv.Load(".env")

	// handle error, if any
	// if err != nil {
	// 	_ = fmt.Errorf("error loading .env-example file")
	// }

	// format connection string from environment variables
	databaseURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbUsername, dbPassword, dbHost, dbName)

	// create a database object which can be used to communicate with database
	db, err := sql.Open(database, databaseURL)

	// handle error, if any
	if err != nil {
		log.Fatal(database+"failed to start", err)
	}

	// test database connection
	if err := db.PingContext(ctx); err != nil {
		log.Fatal("could not ping the database", err)
	}

	// return database object
	return db
}
