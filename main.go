package main

import (
	"fmt"
	"log"
	"net/http"
)

func viewIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{
		"message": "Hello Gophers!"
	}`))
}

func router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", viewIndex)
	return mux
}

func main() {
	port := 5000
	log.Printf("Running on http://127.0.0.1:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router()))
}
