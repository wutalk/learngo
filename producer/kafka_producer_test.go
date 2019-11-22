package producer_test

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/require"
)

func TestProduceMessage(t *testing.T) {
	t.Log("Just use as main")
	log.SetOutput(os.Stdout)

	// kafka producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	for i := 0; i < 5; i++ {
		msg := &sarama.ProducerMessage{Topic: "test", Value: sarama.StringEncoder("testing message from golang#" + strconv.Itoa(i))}
		partition, offset, err := producer.SendMessage(msg)

		require.True(t, partition < 3)

		if err != nil {
			log.Printf("FAILED to send message: %s\n", err)
		} else {
			log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
			fmt.Printf("> message sent to partition %d at offset %d\n", partition, offset)
		}
	}
	// kafka producer end
}
