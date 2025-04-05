package main

import (
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	const port = "8080"
	const filepathRoot = "."

	mux := http.NewServeMux()

	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
	}

	handler := http.FileServer(http.Dir(filepathRoot))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app/", handler)))
	mux.HandleFunc("GET /api/healthz", handlerRead)
	mux.HandleFunc("GET /api/error", handlerError)
	mux.HandleFunc("GET /admin/metrics", apiCfg.metricsHandler)
	mux.HandleFunc("POST /admin/reset", apiCfg.resetHandler)

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
