package scratch

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	HandleFn  func(context.Context, []byte) error
	MaxCount  uint32
	MsgCount  uint32
	MsgLenSum uint32
}

func (c *Consumer) Handle(ctx context.Context, data []byte) error {
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

func (c *Consumer) IsDone() bool {
	return c.MsgCount >= c.MaxCount
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	log.Println("consumer setup...")
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("consumer cleanup...")
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message := <-claim.Messages():
			// log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
			if err := c.Handle(session.Context(), message.Value); err != nil {
				log.Println("consumer handle error:", err)
			}
			session.MarkMessage(message, "")
			if c.IsDone() {
				return nil
			}

		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/Shopify/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}
