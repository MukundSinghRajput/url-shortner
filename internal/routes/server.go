package routes

import (
	"net/http"
	"time"
)

type Server struct{}

func NewServer() http.Server {
	server := Server{}

	return http.Server{
		Addr:         ":8080",
		Handler:      server.RegisterRoutes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}
}
