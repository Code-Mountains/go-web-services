package main

import (
	"flag"
	"log"
	"net/http"

	"readinglist.thecodemountains.com/internal/models"
)

type application struct {
	readlinglist *models.ReadinglistModel
}

func main() {

	addr := flag.String("addr", ":80", "HTTP network address")
	endpoint := flag.String("endpoint", "http://localhost:4000/v1/books", "Endpoint for the readinglist web service")

	app := &application{
		readlinglist: &models.ReadinglistModel{
			Endpoint: *endpoint,
		},
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting server on port %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
