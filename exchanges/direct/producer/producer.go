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

	ch.ExchangeDeclare("X_direct", "direct", true, false, false, false, nil)
	ch.QueueDeclare("Q_direct_error", true, false, false, false, nil)
	ch.QueueDeclare("Q_direct_info", true, false, false, false, nil)
	ch.QueueDeclare("Q_direct_warning", true, false, false, false, nil)
	ch.QueueBind("Q_direct_error", "error", "X_direct", false, nil)
	ch.QueueBind("Q_direct_info", "info", "X_direct", false, nil)
	ch.QueueBind("Q_direct_warning", "warning", "X_direct", false, nil)

	err = ch.Publish("X_direct", "error", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("this is an error message!"),
	})
	err = ch.Publish("X_direct", "info", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("this is an info message!"),
	})
	err = ch.Publish("X_direct", "warning", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("this is an warning message!"),
	})

	if err != nil {
		fmt.Println("Error in push:", err)
		os.Exit(1)
	}

	fmt.Println("Message Pushed 🎉")
}
