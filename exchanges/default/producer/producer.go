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

	ch.QueueDeclare("Q_default_1", true, false, false, false, nil)
	ch.QueueDeclare("Q_default_2", true, false, false, false, nil)

	err = ch.Publish("", "Q_default_1", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("this is message of queue 1!"),
	})
	err = ch.Publish("", "Q_default_2", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("this is message of queue 2!"),
	})

	if err != nil {
		fmt.Println("Error in push:", err)
		os.Exit(1)
	}

	fmt.Println("Message Pushed 🎉")
}
