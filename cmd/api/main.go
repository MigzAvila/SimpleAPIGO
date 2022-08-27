//Filename: cmd/api/main.go

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string // dev, stg, prd, etc...
}

// dependencies injections
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	//read in the flag that are needed to populate the config
	flag.IntVar(&cfg.port, "port", 4000, "API port")
	flag.StringVar(&cfg.env, "env", "dev", "dev, stg, prd")
	flag.Parse()

	//create a logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	//create install of out app
	app := &application{
		config: cfg,
		logger: logger,
	}
	//create out new servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	//create our http server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server ats %s", cfg.env, srv.Addr)
	//start the server
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
