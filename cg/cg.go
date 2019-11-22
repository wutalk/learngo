package cg

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
)

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// handle your message, here only prints it out
	fmt.Println("in consumerGroupSession.ConsumeClaim. My shares:")
	for t, p := range sess.Claims() {
		fmt.Println("topic:", t, ", partitions:", p)
	}
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d. msg: %s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func Consume() {
	groupName := "my-group"
	topicName := "test"
	// Init config, specify appropriate version
	config := sarama.NewConfig()
	config.Version = sarama.V1_0_0_0
	config.Consumer.Return.Errors = true
	// config.Consumer.Offsets.Initial = sarama.OffsetNewest // default
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Start with a client
	// client, err := sarama.NewClient([]string{"localhost:9092"}, config)
	// if err != nil {
	// 	panic(err)
	// }
	// defer func() { _ = client.Close() }()

	// Start a new consumer group
	// group, err := sarama.NewConsumerGroupFromClient("my-group", client)
	group, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, groupName, config)
	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		topics := []string{topicName}
		handler := exampleConsumerGroupHandler{}

		err := group.Consume(ctx, topics, handler) // will block here (until a rebalance ocurs)
		if err != nil {
			panic(err)
		}
		fmt.Println("partitions will be re-assigned") // no change to go here. If more consumers in consumergroup, will be here
	}
}
