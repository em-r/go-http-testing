package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 5000
	log.Printf("Running on http://127.0.0.1:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router()))
}
