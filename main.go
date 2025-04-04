package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	const port = "8080"
	const filepathRoot = "."

	mux := http.NewServeMux()

	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))

	mux.HandleFunc("/healthz", handlerRead)
	mux.HandleFunc("/error", handlerError)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server is running on port " + port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
