package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

type operationStatus struct {
	OperationID string    `json:"operationId"`
	Timestamp   time.Time `json:"timestamp"`
	Status      string    `json:"status"`
}

func main() {
	opID := flag.String("opID", "7465fc42f22a4a87b047b1bb2430e56f", "operation ID")
	st := flag.String("st", "started", "operation status")
	brokerURL := flag.String("b", "localhost:9092", "broker URL")
	topicName := flag.String("t", "operations.status", "topic name from which to receive")
	flag.Parse()
	fmt.Println(*opID, *st)

	opStatus := &operationStatus{
		OperationID: *opID,
		Timestamp:   time.Now(),
		Status:      *st,
	}
	fmt.Printf("message prepared: %#v\n", *opStatus)
	ProduceOneMessage(*brokerURL, *topicName, opStatus)
}

func ProduceOneMessage(brokerURL, topicName string, status *operationStatus) {
	log.SetOutput(os.Stdout)
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

	v, err := json.Marshal(status)
	if err != nil {
		log.Fatal("failed to marshal", err)
	}
	msg := &sarama.ProducerMessage{
		Topic: topicName,
		Key:   sarama.StringEncoder(status.OperationID), // msg with same operationId will come to same partition
		Value: sarama.StringEncoder(v),
	}
	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("message sent to topic: %s, partition:%d, offset: %d\n", topicName, partition, offset)
	}
	// kafka producer end
}
