package router

import (
	"database/sql"
	"msadd/logger"
	"net/http"
)

func Mux(db *sql.DB, log *logger.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		add(w, r, log, db)
	})

	return mux
}
