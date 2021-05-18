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

	switch i {
	case 0:
		consumer.ConsumeMessage()
	case 1:
		cg.Consume()
	case 11:
		cg1.Consume()
	case 2:
		cluster.ConsumeCluster()
	case 3:
		producer.ProduceMessage()
	// other
	case -2:
		envparse.EnvParse()
	case -3:
		net.ReadPage("https://example.com")
	case -4:
		s := net.NewMyServer()
		s.StartHttpServer()
	case -5:
		json.MarshalAndUnmarshalDB()
	case -6:
		runes.RuneUsage()
	default:
		fmt.Println("do nothing")
	}

	fmt.Println("end main")
}
