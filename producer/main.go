package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	log.Println("Hello there")

	conn, err := amqp.Dial("amqp://app:app@localhost:5672/")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	err = ch.Publish(
		"",        // exchange
		"ha-test", // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello world !"),
		},
	)

	if err != nil {
		log.Fatal(err)
	}
}
