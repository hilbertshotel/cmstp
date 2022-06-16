package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func get(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var data string
	err := db.QueryRow("SELECT data FROM test;").Scan(&data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		fmt.Println(err)
		return
	}

	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		fmt.Println(err)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
	w.Write(response)
}

func main() {
	addr := os.Getenv("HOST_ADDR")
	url := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		get(w, r, db)
	})

	fmt.Println("Server listening:", addr)
	http.ListenAndServe(addr, nil)
}
