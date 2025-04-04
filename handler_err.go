package main

import "net/http"

func handlerError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Something went wrong... Not UwU"))
}
