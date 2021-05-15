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

	ch.ExchangeDeclare("X_one", "direct", true, false, false, false, nil)
	ch.ExchangeDeclare("X_two", "direct", true, false, false, false, nil)
	ch.QueueDeclare("Q_one", true, false, false, false, nil)
	ch.QueueDeclare("Q_two", true, false, false, false, nil)

	ch.QueueBind("Q_one", "one", "X_one", false, nil)
	ch.QueueBind("Q_two", "two", "X_two", false, nil)

	ch.ExchangeBind("X_two", "two", "X_one", false, nil)

	err = ch.Publish("X_one", "one", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Message to exchange `X_one` with key `one`"),
	})

	err = ch.Publish("X_one", "two", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Message to exchange `X_one` with key `two`"),
	})

	if err != nil {
		fmt.Println("Error in push:", err)
		os.Exit(1)
	}

	fmt.Println("Message Pushed 🎉")
}
