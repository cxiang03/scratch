package scratch

import (
	"context"
	"testing"
)

// ~ 4k msg/s
func BenchmarkConsumerMsg1k(b *testing.B) {
	c, err := NewConsumerGroup()
	if err != nil {
		b.Fatal(err)
	}

	ctx := context.Background()
	topics := []string{"test-events"}

	b.Run("consume 100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			consumer := &Consumer{
				MaxCount: 100,
			}
			for !consumer.IsDone() {
				if err := c.Consume(ctx, topics, consumer); err != nil {
					b.Fatal(err)
				}
			}
			b.Logf("processed count %d, total length %d", consumer.MsgCount, consumer.MsgLenSum)
		}
	})

	b.Run("consume 1k", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			consumer := &Consumer{
				MaxCount: 1_000,
			}
			for !consumer.IsDone() {
				if err := c.Consume(ctx, topics, consumer); err != nil {
					b.Fatal(err)
				}
			}
			b.Logf("processed count %d, total length %d", consumer.MsgCount, consumer.MsgLenSum)
		}
	})

	b.Run("consume 10k", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			consumer := &Consumer{
				MaxCount: 10_000,
			}
			for !consumer.IsDone() {
				if err := c.Consume(ctx, topics, consumer); err != nil {
					b.Fatal(err)
				}
			}
			b.Logf("processed count %d, total length %d", consumer.MsgCount, consumer.MsgLenSum)
		}
	})
}
