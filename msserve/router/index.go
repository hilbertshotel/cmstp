package router

import (
	"html/template"
	"msget/logger"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request, tmp *template.Template, log *logger.Logger) {
	if err := tmp.ExecuteTemplate(w, "index.html", nil); err != nil {
		log.Error(err)
		return
	}
}
