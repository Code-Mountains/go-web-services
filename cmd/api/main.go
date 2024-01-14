package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "0.0.1"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 8000, "API Server Port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stage|prod)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	addr := fmt.Sprintf(":%d", cfg.port)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheck)

	logger.Printf("starting %s server on %s", cfg.env, addr)

	err := http.ListenAndServe(addr, mux)

	if err != nil {
		fmt.Println(err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "env: %s\n", cfg.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
