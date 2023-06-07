package scratch

import (
	"context"
	"log"
	"sync"

	"github.com/Shopify/sarama"
)

type AsyncConsumer struct {
	sync.Mutex
	ConsumerGroup sarama.ConsumerGroup
	HandleFn      func(context.Context, []byte) error
	MaxCount      uint32
	MsgCount      uint32
	MsgLenSum     uint32
	Closed        bool
}

func (c *AsyncConsumer) Handle(ctx context.Context, data []byte) error {
	if c.HandleFn == nil {
		return nil
	}
	if err := c.HandleFn(context.Background(), data); err != nil {
		return err
	}
	c.MsgCount++
	c.MsgLenSum += uint32(len(data))
	return nil
}

func (c *AsyncConsumer) IsDone() bool {
	return c.MsgCount >= c.MaxCount
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *AsyncConsumer) Setup(sarama.ConsumerGroupSession) error {
	log.Println("async consumer setup...")
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (c *AsyncConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("async consumer cleanup...")
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (c *AsyncConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message := <-claim.Messages():
			if c.IsDone() {
				return nil
			}

			go func() {
				// log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
				if err := c.Handle(session.Context(), message.Value); err != nil {
					log.Println("consumer handle error:", err)
				}

				session.MarkMessage(message, "")
				if !c.Closed && c.IsDone() {
					c.Lock()
					c.ConsumerGroup.Close()
					c.Closed = true
					c.Unlock()
				}
			}()

		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/Shopify/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}
