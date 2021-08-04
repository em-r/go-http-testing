package main

import "net/http"

func viewIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{
		"message": "Hello Gophers!"
	}`))
}

func viewService(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{
		"success": true
	}`))
}
