package router

import (
	"database/sql"
	"html/template"
	"msget/logger"
	"net/http"
)

func Mux(db *sql.DB, tmp *template.Template, log *logger.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	// static
	static := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	mux.Handle("/static/", static)

	// index
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index(w, r, tmp, log)
	})

	return mux
}
