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
	env  string
}

// :BOOK PAGE 98

type application struct {
	config config
	logger *log.Logger
}

func main() {
	// declare an instance of the config struct
	var cfg config

	// Read the value of the port and env command-line
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Enviroment (developer | staging | production)")
	flag.Parse()


	// Initialize a new logger
	logger := log.New(os.Stdout, "__v1: ", log.Ldate | log.Ltime)

	// declare an instance of the application struct, containing the config struct and the logger
	app := &application{
		config: cfg,
		logger: logger,
	}


	// Declare a HTTP server with some sensibel timeout setting.

	// ??? app.routest() but can't application.routes()
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// start the HTTP server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()

	if err != nil {
		logger.Fatal(err)
	}

}
