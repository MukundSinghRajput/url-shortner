package routes

import (
	"encoding/json"
	"net/http"

	"github.com/MukundSinghRajput/goblog/internals/db"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dh bool
	if err := db.DB.Ping(); err == nil {
		dh = true
	} else {
		dh = false
	}
	json.NewEncoder(w).Encode(map[string]bool{
		"alive": true,
		"db":    dh,
	})
}
