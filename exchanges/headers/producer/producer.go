package main

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Connecting to RabbitMQ...")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer ch.Close()

	ch.ExchangeDeclare("X_headers", "headers", true, false, false, false, nil)
	ch.QueueDeclare("Q_headers_1", true, false, false, false, nil)
	ch.QueueDeclare("Q_headers_2", true, false, false, false, nil)
	ch.QueueBind("Q_headers_1", "", "X_headers", false, amqp.Table{
		"x-match": "all",
		"job":     "convert",
		"format":  "jpeg",
	})
	ch.QueueBind("Q_headers_2", "", "X_headers", false, amqp.Table{
		"x-match": "any",
		"job":     "convert",
		"format":  "jpeg",
	})

	err = ch.Publish("X_headers", "", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("This messages is for match all "),
		Headers: amqp.Table{
			"job":    "convert",
			"format": "jpeg",
		},
	})

	err = ch.Publish("X_headers", "", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("This messages is for match any "),
		Headers: amqp.Table{
			"job":    "convert",
			"format": "bmp",
		},
	})

	if err != nil {
		fmt.Println("Error in push:", err)
		os.Exit(1)
	}

	fmt.Println("Message Pushed 🎉")
}
