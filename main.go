package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	webPort = "80"
)

type Config struct{}

func main() {
	log.Printf("Starting broker service on port %s\n", webPort)

	// initialize the connection of rabbitmq
	conn, err := connect()
	if err != nil {
		log.Println("Failed to connect to RabbitMQ", err)
		os.Exit(1)
	}

	defer conn.Close()

	app := Config{}

	// Calling consumer function in its goroutine to run concurrently with the http server
	// and passing rabbitmq connection to it
	go func() {
		app.RecivedReviewToRabbitmq(conn)
	}()

	// define http server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// don't continue until rabbit is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			log.Println("Connected to RabbitMQ!")
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}
