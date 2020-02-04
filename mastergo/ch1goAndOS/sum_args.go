package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	fmt.Println("args:", args)
	if len(args) == 0 {
		fmt.Println("no argument found")
		os.Exit(0)
	}
	exp, sum := SumArgs(args)
	fmt.Println(exp, "=", sum)
}

func SumArgs(args []string) (string, float64) {
	var sum float64
	var buf bytes.Buffer
	for _, arg := range args {
		n, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Println(arg, "is not a number, continue")
			continue
		}
		sum += n
		if n >= 0 {
			buf.WriteString("+")
		}
		buf.WriteString(arg)
	}
	exp := buf.String()
	begin := 0
	if strings.HasPrefix(exp, "+") {
		begin = 1
	}
	exp = exp[begin:buf.Len()]
	return exp, sum
}
