package scratch

import (
	"testing"

	"github.com/Shopify/sarama"
	"github.com/icrowley/fake"
)

// ~ 40k msg/s [mac-host -> docker]
func BenchmarkProduceMsg(b *testing.B) {
	topic := "test-events"

	b.Run("produce 1k msg", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			b.StopTimer()
			p, err := NewProducer()
			if err != nil {
				b.Fatal(err)
			}

			procMsgCh := make(chan struct{})
			go func() {
				for msg := range p.Successes() {
					b.Logf("msg: %v", msg)
				}
				close(procMsgCh)
			}()

			procErrCh := make(chan struct{})
			go func() {
				for err := range p.Errors() {
					b.Logf("err: %v", err)
				}
				close(procErrCh)
			}()
			b.StartTimer()

			for j := 0; j < 1000; j++ {
				msg := fake.Title()
				p.Input() <- &sarama.ProducerMessage{
					Topic: topic,
					Key:   nil,
					Value: sarama.StringEncoder(msg),
				}
			}

			p.Close()
			<-procMsgCh
			<-procErrCh
		}
	})
}

// ~ 2k msg/s
func BenchmarkProduceMsgSync(b *testing.B) {
	topic := "test-events"
	b.Run("sync produce 1k msg", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			b.StopTimer()
			p, err := NewSyncProducer()
			if err != nil {
				b.Fatal(err)
			}

			b.StartTimer()
			for j := 0; j < 1000; j++ {
				msg := fake.Title()
				_, _, err := p.SendMessage(&sarama.ProducerMessage{
					Topic: topic,
					Key:   nil,
					Value: sarama.StringEncoder(msg),
				})
				if err != nil {
					b.Fatal(err)
				}
			}
			p.Close()
		}
	})
}
