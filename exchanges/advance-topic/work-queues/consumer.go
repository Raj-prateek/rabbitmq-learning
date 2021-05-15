package main

import (
	"fmt"
	"os"
	"strconv"
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
	msgCount := 0

	msgs, err := ch.Consume(
		"Q_fanout_1", //queue
		"",           //consumer
		true,         //auto-ack
		false,        //exclusive
		false,        //no-local
		false,        //no-wait
		nil,          //args
	)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	go func(msgCount int) {
		for msg := range msgs {
			msgCount = msgCount + 1
			fmt.Println("Messages: ", string(msg.Body))
			fmt.Println("Messages Time: ", time.Now())
			n, err := strconv.Atoi(string(msg.Body))
			if err != nil {
				fmt.Println("Error: not able to process msg.Body")
				continue
			}
			time.Sleep(time.Second * time.Duration(n))

		}
	}(msgCount)

	select {
	case <-time.After(time.Second * 100):
		fmt.Printf("Total Messages Fetched: %d\n", msgCount)
		fmt.Println("No more messages in queue. Timing out...")

	}
}
