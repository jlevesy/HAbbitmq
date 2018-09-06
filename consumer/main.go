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

	msgs, err := ch.Consume(
		"ha-test",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	for msg := range msgs {
		log.Println(string(msg.Body))
	}
}
