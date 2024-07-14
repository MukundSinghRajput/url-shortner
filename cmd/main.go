package main

import (
	"log"

	"github.com/MukundSinghRajput/url-shotner/internal/routes"
)

func main() {
	s := routes.NewServer()
	log.Println("Listening on port 8080")
	s.ListenAndServe()
}
