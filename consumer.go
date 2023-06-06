package scratch

import (
	"log"

	"github.com/Shopify/sarama"
)

func NewConsumerGroup() (sarama.ConsumerGroup, error) {
	c := sarama.NewConfig()

	c.ClientID = "sarama-localhost"
	c.ChannelBufferSize = 256
	c.ApiVersionsRequest = true
	c.Version = sarama.V3_2_3_0

	c.Net.MaxOpenRequests = 4
	c.Consumer.Return.Errors = true
	c.Consumer.Offsets.Initial = sarama.OffsetOldest
	c.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRoundRobin}

	return sarama.NewConsumerGroup(brokers, "test-group", c)
}

type Consumer struct {
	MsgCount  uint32
	MsgLenSum uint32
}

func (c *Consumer) IsDone() bool {
	return c.MsgCount >= 10_000
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
			c.MsgCount++
			c.MsgLenSum += uint32(len(message.Value))
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
