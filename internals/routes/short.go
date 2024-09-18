package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/MukundSinghRajput/goblog/internals/configs"
	"github.com/MukundSinghRajput/goblog/internals/db"
	"github.com/MukundSinghRajput/goblog/internals/db/database"
	utils "github.com/MukundSinghRajput/goblog/internals/util"
)

func Short(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Please pass a valid body", http.StatusBadRequest)
		return
	}

	if data.URL == "" {
		http.Error(w, "The url body can't be empty", http.StatusBadRequest)
		return
	}

	q := database.New(db.DB)

	w.Header().Set("Content-Type", "application/json")

	if u, err := q.GetByOrigin(context.Background(), data.URL); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			short := utils.CreateUrl(data.URL)
			json.NewEncoder(w).Encode(map[string]string{
				"shortURL": short,
			})
		}
	} else {
		json.NewEncoder(w).Encode(map[string]string{
			"shortURL": fmt.Sprintf("%v/%v", configs.Config.URL, u.Short),
		})
	}
}
