package scratch

import (
	"context"
	"testing"
)

func BenchmarkBatchConsumerMsg1k(b *testing.B) {
	c, err := NewConsumerGroup()
	if err != nil {
		b.Fatal(err)
	}

	ctx := context.Background()
	topics := []string{"test-events"}
	for i := 0; i < b.N; i++ {
		consumer := &BatchConsumer{}
		for !consumer.IsDone() {
			if err := c.Consume(ctx, topics, consumer); err != nil {
				b.Fatal(err)
			}
		}
		b.Logf("processed count %d, total length %d", consumer.MsgCount, consumer.MsgLenSum)
	}
}
