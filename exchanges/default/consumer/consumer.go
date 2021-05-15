package main

import (
	"fmt"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func main() {

	fmt.Println("Connecting to RabbitMQ...")
	conn, err := amqp.Dial("amqp://localhost:5672")
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

	fQ := []string{
		"Q_default_1",
		"Q_default_2",
	}

	msgCount := 0

	for _, queue := range fQ {
		msgs, err := ch.Consume(
			queue, //queue
			"",    //consumer
			true,  //auto-ack
			false, //exclusive
			false, //no-local
			false, //no-wait
			nil,   //args
		)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		go func() {
			for msg := range msgs {
				msgCount = msgCount + 1
				fmt.Println("Messages: ", string(msg.Body))
			}
		}()

	}

	select {
	case <-time.After(time.Second * 10):
		fmt.Printf("Total Messages Fetched: %d\n", msgCount)
		fmt.Println("No more messages in queue. Timing out...")

	}
}
