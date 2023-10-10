package main

import "yqnft/YQRabbitMQ/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("KP")
	rabbitmq.ConsumeSimple()
}
