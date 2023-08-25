package main

import (
	"github.com/paxson/messagequeue"
)

func main() {
	mq := messagequeue.NewMessageQueue()
	mq.Connect("nats://localhost:4222")
}
