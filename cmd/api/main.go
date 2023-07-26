package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mercadolibre/fury_go-core/pkg/web"
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	app, err := fury.NewWebApplication()
	if err != nil {
		return err
	}

	app.Post("/hello", helloHandler)

	return app.Run()
}

func helloHandler(w http.ResponseWriter, r *http.Request) error {
	return web.EncodeJSON(w, fmt.Sprintf("%s, world!", r.URL.Path[1:]), http.StatusOK)
}
