package main

import (
	"fmt"
	"os"
	"time"

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

	msgs, err := ch.Consume("Q_request", "", true, false, false, false, nil)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	msgCount := 0
	go func() {
		for msg := range msgs {
			msgCount = msgCount + 1
			fmt.Println("Messages: ", string(msg.Body))
			err = ch.Publish("", "Q_response", false, false, amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(fmt.Sprintf("Message recieved: %s", msg.Body)),
			})
		}
	}()

	select {
	case <-time.After(time.Second * 10):
		fmt.Println("No more messages in queue. Timing out...")
	}
}
