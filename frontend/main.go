package main

import (
	"log"
	"net/http"
)

func main() {
	setupRouters()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupRouters() {
	http.HandleFunc("/", rootHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
