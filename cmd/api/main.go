package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/healthcheck", healthcheck)

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println(err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "env: %s\n", "dev")
	fmt.Fprintf(w, "version: %s\n", "0.0.1")
}
