package message_queue

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

func InitQueue() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open RabbitMQ channel: %v", err)
	}
	return ch
}

func PublishMessage(queueName string, message interface{}) {
	ch := InitQueue()
	defer ch.Close()

	msg, _ := json.Marshal(message)
	ch.Publish("", queueName, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        msg,
	})
}
