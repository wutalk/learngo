package envparse

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"time"
)

/*
Set some environment variables:

export MYAPP_DEBUG=false
export MYAPP_PORT=8080
export MYAPP_USER=Kelsey
export MYAPP_RATE="0.5"
export MYAPP_TIMEOUT="3m"
export MYAPP_USERS="rob,ken,robert"
export MYAPP_COLOR_CODES="red:1,green:2,blue:3"
export KAFKA_CONSUMER_GROUP="test-group"
*/

// specification can be lowercase
// but fields name must be capitalized
type specification struct {
	Debug      bool
	Port       int
	User       string
	Users      []string
	Rate       float32
	Timeout    time.Duration
	ColorCodes map[string]int `split_words:"true"`
	Group      string         `envconfig:"KAFKA_CONSUMER_GROUP"`
}

func EnvParse() {
	var s specification
	err := envconfig.Process("myapp", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	format := "Debug: %v\nPort: %d\nUser: %s\nRate: %0.2f\nTimeout: %s\n"
	_, err = fmt.Printf(format, s.Debug, s.Port, s.User, s.Rate, s.Timeout)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Users:")
	for _, u := range s.Users {
		fmt.Printf("  %s\n", u)
	}

	fmt.Println("Color codes:")
	for k, v := range s.ColorCodes {
		fmt.Printf("  %s: %d\n", k, v)
	}
	fmt.Println("group:", s.Group)
}
