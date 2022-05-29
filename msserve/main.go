package main

import (
	"database/sql"
	"html/template"
	"msget/config"
	"msget/logger"
	"msget/router"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	// logger
	log := logger.Init()

	// config
	cfg := config.Init()

	// database
	db, err := sql.Open("postgres", cfg.ConnStr)
	if err != nil {
		log.Error(err)
		return
	}

	// templates
	tmp, err := template.ParseGlob(cfg.TmpDir)
	if err != nil {
		log.Error(err)
		return
	}

	// server
	server := http.Server{
		Addr:         cfg.HostAddr,
		Handler:      router.Mux(db, tmp, log),
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
