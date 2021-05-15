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

	msgs, err := ch.Consume("Q_response", "", true, false, false, false, nil)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	msgCount := 0
	go func() {
		for msg := range msgs {
			msgCount = msgCount + 1
			fmt.Println("Messages: ", string(msg.Body))
		}
	}()

	go takeInput(ch)

	<-time.After(time.Second * 50)
}

func takeInput(ch *amqp.Channel) {
	var request string
	fmt.Println("Enter your request:")
	fmt.Scanln(&request)

	err := ch.Publish("", "Q_request", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(request),
	})

	if err != nil {
		fmt.Println("Error in push:", err)
		os.Exit(1)
	}
}
