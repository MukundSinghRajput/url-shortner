package routes

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/MukundSinghRajput/goblog/internals/db"
	"github.com/MukundSinghRajput/goblog/internals/db/database"
	"github.com/MukundSinghRajput/goblog/internals/web/components"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		components.Base("Shrink", components.Shotner()).Render(w)
	} else {
		id := r.URL.Path[len("/"):]
		if url, ok := database.New(db.DB).GetByShort(context.Background(), id); ok != nil {
			if errors.Is(ok, sql.ErrNoRows) {
				http.Error(w, "Sorry the short url with this id doesn't exists", http.StatusNotFound)
			} else {
				http.Error(w, fmt.Sprintf("Error: %v\n\nReport it in github issues: https://github.com/MukundSinghRajput/url-shortner", ok), http.StatusNotFound)
			}
		} else {
			if !strings.HasPrefix(url.Origin, "http://") && !strings.HasPrefix(url.Origin, "https://") {
				url.Origin = "https://" + url.Origin
			}
			http.Redirect(w, r, url.Origin, http.StatusFound)
		}
	}
}
