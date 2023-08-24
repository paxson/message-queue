package main

import (
	"github.com/paxson/message-queue/mq"
)

func main() {
	var mq mq.MessageQueue
	mq = mq.NewMessageQueue()
	mq.Connect("nats://localhost:4222")
}
