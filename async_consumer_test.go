package scratch

import (
	"context"
	"testing"
	"time"
)

func BenchmarkConsumerMsgAsync(b *testing.B) {
	ctx := context.Background()
	topics := []string{"test-events"}

	b.Run("async consume 100, sleep 1 ms", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			c, err := NewConsumerGroup()
			if err != nil {
				b.Fatal(err)
			}
			b.StartTimer()

			consumer := &AsyncConsumer{
				ConsumerGroup: c,
				HandleFn: func(ctx context.Context, data []byte) error {
					time.Sleep(time.Millisecond)
					return nil
				},
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

	b.Run("async consume 1k, sleep 1 ms", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			c, err := NewConsumerGroup()
			if err != nil {
				b.Fatal(err)
			}
			b.StartTimer()

			consumer := &AsyncConsumer{
				ConsumerGroup: c,
				HandleFn: func(ctx context.Context, data []byte) error {
					time.Sleep(time.Millisecond)
					return nil
				},
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

	b.Run("async consume 10k, sleep 1 ms", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			c, err := NewConsumerGroup()
			if err != nil {
				b.Fatal(err)
			}
			b.StartTimer()

			consumer := &AsyncConsumer{
				ConsumerGroup: c,
				HandleFn: func(ctx context.Context, data []byte) error {
					time.Sleep(time.Millisecond)
					return nil
				},
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
