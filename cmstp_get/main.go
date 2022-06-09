package main

import (
	"cmstp_get/logger"
	"cmstp_get/router"
	"database/sql"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	// env
	hostAddr := os.Getenv("HOST_ADDR")
	dbUrl := os.Getenv("DB_URL")

	// logger
	log := logger.Init()

	// database
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Error(err)
		return
	}

	// server
	server := http.Server{
		Addr:         hostAddr,
		Handler:      router.Mux(db, log),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// run
	errch := make(chan error)

	go func(errch chan error) {
		log.Ok("Now listening on " + hostAddr)
		errch <- server.ListenAndServe()
	}(errch)

	log.Error(<-errch)
}
