package main

import (
	"cmstp_get/config"
	"cmstp_get/logger"
	"cmstp_get/router"
	"database/sql"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	// logger
	log := logger.Init()

	// config
	cfg, err := config.Init()
	if err != nil {
		log.Error(err)
		return
	}

	// database
	db, err := sql.Open("postgres", cfg.ConnStr)
	if err != nil {
		log.Error(err)
		return
	}

	// server
	server := http.Server{
		Addr:         cfg.HostAddr,
		Handler:      router.Mux(db, log),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// run
	errch := make(chan error)

	go func(errch chan error) {
		log.Ok("Now listening on " + cfg.HostAddr)
		errch <- server.ListenAndServe()
	}(errch)

	log.Error(<-errch)
}
