package producer

import (
	"log"
	"os"
	"strconv"

	"github.com/Shopify/sarama"
)

func ProduceMessage() {
	log.SetOutput(os.Stdout)

	brokerURL := "localhost:9092"
	topicName := "test"
	// kafka producer
	producer, err := sarama.NewSyncProducer([]string{brokerURL}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	for i := 0; i < 10; i++ {
		// demo same key will hash to same partition. assume we have 5 keys, and have 3 partitions,
		// you will see massage with same key, will always to to same partition
		key := "PLMN-" + strconv.Itoa(i%5)
		msg := &sarama.ProducerMessage{
			Topic: topicName,
			Key:   sarama.StringEncoder(key),
			Value: sarama.StringEncoder("testing message from golang#" + strconv.Itoa(i) + ", key=" + key),
		}
		partition, offset, err := producer.SendMessage(msg)

		if err != nil {
			log.Printf("FAILED to send message: %s\n", err)
		} else {
			log.Printf("> message sent to topic: %s, partition:%d, offset: %d\n", topicName, partition, offset)
		}
	}
	// kafka producer end
}
