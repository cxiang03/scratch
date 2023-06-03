package scratch

import (
	"context"
	"testing"
)

// ~ 6k msg/s
func BenchmarkConsumerMsg1k(b *testing.B) {
	c, err := NewConsumerGroup()
	if err != nil {
		b.Fatal(err)
	}

	ctx := context.Background()
	topics := []string{"test-events"}
	for i := 0; i < b.N; i++ {
		consumer := &Consumer{}
		for !consumer.IsDone() {
			if err := c.Consume(ctx, topics, consumer); err != nil {
				b.Fatal(err)
			}
		}
	}
}
