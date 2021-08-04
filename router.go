package main

import "net/http"

func router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", viewIndex)
	mux.HandleFunc("/service", viewService)
	return mux
}
