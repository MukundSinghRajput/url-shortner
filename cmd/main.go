package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MukundSinghRajput/goblog/internals/configs"
	"github.com/MukundSinghRajput/goblog/internals/db"
	"github.com/MukundSinghRajput/goblog/internals/server"
	"github.com/charmbracelet/log"
)

func main() {
	s := server.NewServer()
	logger := log.NewWithOptions(os.Stderr, log.Options{
		Prefix: "TINY",
	})
	logger.Infof("Server has started on port :%v", configs.Config.PORT)
	db := db.NewDB()
	defer db.Close()
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Server error: ", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	logger.Info("Received shutdown signal, starting graceful shutdown...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		logger.Error("Graceful shutdown failed: ", err)
	} else {
		logger.Info("Shutdown complete.")
	}
}
