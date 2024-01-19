package main

import (
	"flag"
	"log"
	"net/http"
)

type application struct {
}

func main() {
	app := &application{}

	addr := flag.String("addr", ":8000", "HTTP network address")
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting the server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
