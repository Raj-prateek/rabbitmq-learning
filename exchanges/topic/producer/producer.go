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
	ch.ExchangeDeclare("X_topic", "topic", true, false, false, false, nil)
	ch.QueueDeclare("Q_topic_1", true, false, false, false, nil)
	ch.QueueDeclare("Q_topic_2", true, false, false, false, nil)
	ch.QueueDeclare("Q_topic_3", true, false, false, false, nil)
	ch.QueueBind("Q_topic_1", "*.image.*", "X_topic", false, nil)
	ch.QueueBind("Q_topic_2", "#.image", "X_topic", false, nil)
	ch.QueueBind("Q_topic_3", "image.#", "X_topic", false, nil)

	err = ch.Publish("X_topic", "image.jpeg", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello World!"),
	})

	err = ch.Publish("X_topic", "bitmap.image.jpeg", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello World!"),
	})

	err = ch.Publish("X_topic", "bitmap.image", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello World!"),
	})

	if err != nil {
		fmt.Println("Error in push:", err)
		os.Exit(1)
	}

	fmt.Println("Message Pushed 🎉")
}
