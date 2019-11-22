package consumer

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

func ConsumeMessage() {
	log.SetOutput(os.Stdout)

	// config := sarama.NewConfig()
	// config.Consumer.Group.Member

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitions, err := consumer.Partitions("test")
	log.Println("partition IDs: ", partitions)

	for _, id := range partitions {
		go consume(consumer, "test", id)
	}

	// consume(consumer, "test", partitions[0])

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	shut := <-signals
	log.Println("shutdown, ", shut)
}

func consume(consumer sarama.Consumer, topic string, partitionID int32) {
	log.Printf("start consumer messages from topic %s on partition #%d\n", topic, partitionID)
	partitionConsumer, err := consumer.ConsumePartition(topic, partitionID, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	defer func() {
		log.Println("Close partitionConsumer")
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0

ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message from partition #%d offset %d, message: %s\n", partitionID, msg.Offset, msg.Value)
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("In total consumed from topic %s patition %d: %d messages\n", topic, partitionID, consumed)
}
