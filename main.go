package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"stumble/router"
)

func main() {
	r := router.Router()
	fs := http.FileServer(http.Dir("build"))
	http.Handle("/", fs)
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
