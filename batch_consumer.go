package scratch

import (
	"log"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

type BatchConsumer struct {
	sync.Mutex

	MsgCount  uint32
	MsgLenSum uint32

	ticker  *time.Ticker
	session sarama.ConsumerGroupSession
	msgBuf  []*sarama.ConsumerMessage
	msgsCh  chan []*sarama.ConsumerMessage
}

func (c *BatchConsumer) IsDone() bool {
	return c.MsgCount >= 1000
}

func (c *BatchConsumer) acceptMsg(msg *sarama.ConsumerMessage) {
	c.Lock()
	c.msgBuf = append(c.msgBuf, msg)
	if len(c.msgBuf) >= 16 {
		c.flushBuf()
	}
	c.Unlock()
}

func (c *BatchConsumer) flushBuf() {
	if len(c.msgBuf) > 0 {
		c.msgsCh <- c.msgBuf
		c.msgBuf = make([]*sarama.ConsumerMessage, 0, 128)
	}
}

func (c *BatchConsumer) Run() {
	for {
		msgs := <-c.msgsCh
		for _, msg := range msgs {
			c.MsgCount++
			c.MsgLenSum += uint32(len(msg.Value))
			time.Sleep(2 * time.Millisecond)
			c.session.MarkMessage(msg, "")
		}
	}
}

func (c *BatchConsumer) Setup(sarama.ConsumerGroupSession) error {
	log.Println("consumer setup...")
	c.msgBuf = make([]*sarama.ConsumerMessage, 0, 16)
	c.msgsCh = make(chan []*sarama.ConsumerMessage, 16)
	c.ticker = time.NewTicker(100 * time.Microsecond)
	for i := 0; i < 16; i++ {
		go c.Run()
	}
	return nil
}

func (c *BatchConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	log.Println("start claim")
	c.Lock()
	c.session = session
	c.Unlock()

	for {
		select {
		case msg := <-claim.Messages():
			// log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
			if c.IsDone() {
				return nil
			}
			c.acceptMsg(msg)

		case <-c.ticker.C:
			c.Lock()
			c.flushBuf()
			c.Unlock()

		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/Shopify/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}

func (c *BatchConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("consumer cleanup...")
	return nil
}
