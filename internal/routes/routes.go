package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/MukundSinghRajput/url-shortner/internal/database/controller"
	"github.com/MukundSinghRajput/url-shortner/internal/utils"
	base "github.com/MukundSinghRajput/url-shortner/internal/view"
)

func (server Server) RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	assetsDir := "internal/web/assets"
	fs := http.FileServer(http.Dir(filepath.Clean(assetsDir)))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			base.Base("Shotner").Render(context.Background(), w)
		} else {
			id := r.URL.Path[len("/"):]
			if url, ok := controller.Get(id); !ok {
				http.Error(w, "Sorry the short url with this id doesn't exists", http.StatusNotFound)
			} else {
				if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
					url = "https://" + url
				}
				http.Redirect(w, r, url, http.StatusFound)
			}
		}
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]bool{
			"alive": true,
		})
	})

	mux.HandleFunc("POST /short", func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			URL string `json:"url"`
		}

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Please pass a valid body", http.StatusBadRequest)
			return
		}

		shortUrl := utils.CreateUrl(data.URL)
		response := struct {
			ShortURL string `json:"shortURL"`
		}{ShortURL: shortUrl}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&response)
	})

	return mux
}
