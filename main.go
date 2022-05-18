package main

import (
	"log"
	"net/http"
	"os"
	"stumble/router"
)

func main() {
	r := router.Router()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
