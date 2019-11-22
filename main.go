package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"

	"learngo/cg"
	"learngo/cg1"
	"learngo/cluster"
	"learngo/envparse"
	"learngo/json"
	"learngo/net"
	"learngo/producer"
	"learngo/runes"

	"github.com/Shopify/sarama"
)

func main() {
	var i int = -1
	args := os.Args[1:]
	fmt.Println("args:", args)
	if len(args) > 0 {
		id := args[0]
		n, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
		}
		i = n
	}
	fmt.Println("start main")

	// switch i {
	// case 0:
	// 	ConsumeMessage()
	// case 1:
	// 	cg.Consume()
	// }

	if i == 0 {
		ConsumeMessage()
	} else if i == 1 {
		cg.Consume()
	} else if i == 11 {
		cg1.Consume()
	} else if i == 2 {
		cluster.ConsumeCluster()
	} else if i == 3 {
		producer.ProduceMessage()
	}

	if i == -2 {
		envparse.EnvParse()
	}
	if i == -3 {
		net.ReadPage("https://example.com")
	}
	if i == -4 {
		net.StartHttpServer()
	}
	if i == -5 {
		// json.MarshalDemo()
		// json.UnmarshalDemo()
		json.MarshalAndUnmarshalDB()
	}
	if i == -6 {
		runes.RuneUsage()
	}

	fmt.Println("end main")
}

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
