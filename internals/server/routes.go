package server

import (
	"net/http"
	"path/filepath"

	"github.com/MukundSinghRajput/goblog/internals/routes"
)

func (server Server) RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	assetsDir := "internals/web/assets"
	fs := http.FileServer(http.Dir(filepath.Clean(assetsDir)))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", routes.Home)
	mux.HandleFunc("POST /short", routes.Short)
	mux.HandleFunc("/health", routes.Health)
	return mux
}
