package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MukundSinghRajput/goblog/internals/configs"
)

type Server struct{}

func NewServer() http.Server {
	server := Server{}

	return http.Server{
		Addr:         fmt.Sprintf(":%v", configs.Config.PORT),
		Handler:      server.RegisterRoutes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}
}
