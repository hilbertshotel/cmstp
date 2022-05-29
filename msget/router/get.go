package router

import (
	"database/sql"
	"encoding/json"
	"msget/logger"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func get(w http.ResponseWriter, r *http.Request, log *logger.Logger, db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users;")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Error(err)
		return
	}

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Error(err)
			return
		}
		users = append(users, user)
	}

	response, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Error(err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(response)
}
