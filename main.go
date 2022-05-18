package main

import (
	"log"
	"net/http"
	"os"
	"stumble/router"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
