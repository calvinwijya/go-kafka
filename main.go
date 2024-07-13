package main

import (
	"github.com/calvinwijya/go-kafka/consumer"
	"github.com/calvinwijya/go-kafka/producer"
)

func main() {
	producer.PublishMessage()
	consumer.ConsumeMessage()
}
