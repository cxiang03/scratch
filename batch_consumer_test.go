package scratch

import (
	"context"
	"testing"
	"time"
)

func BenchmarkBatchConsumeMsg1kAsync(b *testing.B) {
	topics := []string{"test-events"}

	b.Run("batch async consume 1k, sleep 1 ms", func(b *testing.B) {
		c, err := NewConsumerGroup()
		if err != nil {
			b.Fatal(err)
		}
		ctx := context.Background()

		for i := 0; i < b.N; i++ {
			consumer := &BatchConsumer{
				MaxCount: 1000,
				HandleFn: func(ctx context.Context, data []byte) error {
					time.Sleep(time.Millisecond)
					return nil
				},
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
