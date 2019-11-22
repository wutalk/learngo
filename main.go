package main

import (
	"fmt"
	"os"
	"strconv"

	"learngo/cg"
	"learngo/cg1"
	"learngo/cluster"
	"learngo/consumer"
	"learngo/envparse"
	"learngo/json"
	"learngo/net"
	"learngo/producer"
	"learngo/runes"
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

	if i == 0 {
		consumer.ConsumeMessage()
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
		json.MarshalAndUnmarshalDB()
	}
	if i == -6 {
		runes.RuneUsage()
	}

	fmt.Println("end main")
}
