package scratch

import (
	"github.com/Shopify/sarama"
)

func NewProducer() (sarama.AsyncProducer, error) {
	c := sarama.NewConfig()

	c.ClientID = "sarama-localhost"
	c.ChannelBufferSize = 256
	c.ApiVersionsRequest = true
	c.Version = sarama.V3_2_3_0

	c.Net.MaxOpenRequests = 4
	c.Producer.Idempotent = false
	c.Producer.RequiredAcks = sarama.WaitForLocal
	c.Producer.Return.Successes = false
	c.Producer.Return.Errors = true
	c.Producer.Partitioner = sarama.NewRoundRobinPartitioner

	return sarama.NewAsyncProducer(brokers, c)
}

// sync producer is a variant of the async producer
// ~/go/pkg/mod/github.com/!shopify/sarama@v1.38.1/sync_producer.go:75
func NewSyncProducer() (sarama.SyncProducer, error) {
	c := sarama.NewConfig()

	c.ClientID = "sarama-localhost"
	c.ChannelBufferSize = 8
	c.ApiVersionsRequest = true
	c.Version = sarama.V3_2_3_0

	c.Net.MaxOpenRequests = 1
	c.Producer.Idempotent = true
	c.Producer.RequiredAcks = sarama.WaitForAll
	c.Producer.Return.Successes = true
	c.Producer.Return.Errors = true
	c.Producer.Partitioner = sarama.NewRandomPartitioner

	return sarama.NewSyncProducer(brokers, c)
}
