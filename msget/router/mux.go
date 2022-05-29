package router

import (
	"database/sql"
	"msget/logger"
	"net/http"
)

func Mux(db *sql.DB, log *logger.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	// index
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		get(w, r, log, db)
	})

	return mux
}
