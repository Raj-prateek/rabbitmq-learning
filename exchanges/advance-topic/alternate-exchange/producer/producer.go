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

	ch.ExchangeDeclare("X_alternate_two", "fanout", true, false, false, false, nil)
	ch.ExchangeDeclare("X_alternate_one", "direct", true, false, false, false, amqp.Table{
		"alternate-exchange": "X_alternate_two",
	})

	ch.QueueDeclare("Q_alternate_one", true, false, false, false, nil)
	ch.QueueDeclare("Q_alternate_two", true, false, false, false, nil)
	ch.QueueDeclare("Q_alternate_unrouting", true, false, false, false, nil)

	ch.QueueBind("Q_alternate_one", "one", "X_alternate_one", false, nil)
	ch.QueueBind("Q_alternate_two", "two", "X_alternate_one", false, nil)
	ch.QueueBind("Q_alternate_unrouting", "", "X_alternate_two", false, nil)

	err = ch.Publish("X_alternate_one", "one", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Message to routed to queue with key `one`"),
	})

	err = ch.Publish("X_alternate_one", "two", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Message to routed to queue with key `two`"),
	})

	err = ch.Publish("X_alternate_one", "three", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Message to routed to queue with key `three`"),
	})

	if err != nil {
		fmt.Println("Error in push:", err)
		os.Exit(1)
	}

	fmt.Println("Message Pushed 🎉")
}
