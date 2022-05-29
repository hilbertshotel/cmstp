package router

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"msadd/logger"
	"net/http"
)

func add(w http.ResponseWriter, r *http.Request, log *logger.Logger, db *sql.DB) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Error(err)
		return
	}

	var name string
	err = json.Unmarshal(body, &name)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Error(err)
		return
	}

	_, err = db.Exec("INSERT INTO users (name) VALUES ($1)", name)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Error(err)
		return
	}
}
