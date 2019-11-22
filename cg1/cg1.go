// Package cg1 consumer group several goroutines messages handled in 1 place
package cg1

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
)

type payload struct {
	msg  *sarama.ConsumerMessage
	mark func(metadata string)
}

type exampleConsumerGroupHandler struct {
	central chan payload
}

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
		fmt.Println("redirect msg to central process")
		h.central <- payload{
			msg: msg,
			mark: func(metadata string) {
				sess.MarkMessage(msg, metadata) // will marked later
				fmt.Println("msg", string(msg.Value), "marked with metadata", metadata)
			},
		}
		// sess.MarkMessage(msg, "")
	}
	return nil
}

func handlePayload(central chan payload) {
	for {
		p := <-central
		fmt.Println("handle payload, msg", string(p.msg.Value))
		p.mark("handled")
	}
}

func Consume() {
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
	group, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "my-group", config)
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
		c := make(chan payload)
		go handlePayload(c)

		handler := exampleConsumerGroupHandler{central: c}
		topics := []string{"test"}

		err := group.Consume(ctx, topics, handler) // will block here (until a rebalance ocurs)
		if err != nil {
			panic(err)
		}
		fmt.Println("partitions will be re-assigned") // If more consumers in consumergroup, will be here
	}
}
