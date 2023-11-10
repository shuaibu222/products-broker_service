package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	webPort = "9000"
)

type Config struct{}

func main() {
	log.Printf("Starting broker service on port %s\n", webPort)

	app := Config{}

	// define http server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
