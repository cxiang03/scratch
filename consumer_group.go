package scratch

import "github.com/Shopify/sarama"

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

	return sarama.NewConsumerGroup(brokers, "cg-unit-test", c)
}
