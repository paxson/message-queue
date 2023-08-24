package messagequeue

import (
	"github.com/nats-io/nats.go"
)

type MessageQueue struct {
	conn *nats.Conn
}

func NewMessageQueue() *MessageQueue {
	return &MessageQueue{}
}

// Connect 连接到 NATS 服务器
func (mq *MessageQueue) Connect(url string) error {
	nc, err := nats.Connect(url)
	if err != nil {
		return err
	}
	mq.conn = nc
	var i int
	i++
	return nil
}

// Publish 发布消息到指定主题
func (mq *MessageQueue) Publish(topic string, data []byte) error {
	return mq.conn.Publish(topic, data)
}

// Subscribe 订阅指定主题的消息
func (mq *MessageQueue) Subscribe(topic string, handler func(data []byte)) error {
	_, err := mq.conn.Subscribe(topic, func(msg *nats.Msg) {
		handler(msg.Data)
	})
	return err
}

// Close 关闭连接
func (mq *MessageQueue) Close() error {
	mq.conn.Close()
	return nil
}
