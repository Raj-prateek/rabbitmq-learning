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

	ch.ExchangeDeclare("X_fanout", "fanout", true, false, false, false, nil)
	ch.QueueDeclare("Q_fanout_1", true, false, false, false, nil)
	ch.QueueDeclare("Q_fanout_2", true, false, false, false, nil)
	ch.QueueBind("Q_fanout_1", "", "X_fanout", false, nil)
	ch.QueueBind("Q_fanout_2", "", "X_fanout", false, nil)

	body := "Hello World!"
	err = ch.Publish("X_fanout", "", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})

	if err != nil {
		fmt.Println("Error in push:", err)
		os.Exit(1)
	}

	fmt.Println("Message Pushed 🎉")
}
