package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("begin sum time")

	filename := "time-data-dockerinaction.txt"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s", err)
		os.Exit(1)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	var minutes int = 0
	var seconds int = 0
	count := 0
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
		}

		// line := "Title Page <span class=\"minutes\">(01:09 mins)</span>"
		// r1 := regexp.MustCompile(`.*>\(\d\d:\d\d mins\)<.*`)
		r1 := regexp.MustCompile(`minutes\">\(\d\d:\d\d`)
		// r2 := regexp.MustCompile(`.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\] .*`)

		if r1.MatchString(line) {
			// if count > 10 {
			// 	break
			// }
			count++
			// match := r1.FindStringSubmatch(line)
			match := r1.FindAllString(line, 5)
			if len(match) == 1 {
				fmt.Println(line)
				mt := match[0]
				// fmt.Println(mt)
				t := strings.Split(mt, "(")[1]
				// fmt.Println(t)
				ms := strings.Split(t, ":")
				m, err := strconv.Atoi(ms[0])
				if err != nil {
					fmt.Println("fail to parse", ms[0])
					continue
				}
				minutes += m
				s, err := strconv.Atoi(ms[1])
				if err != nil {
					fmt.Println("fail to parse", ms[1])
					continue
				}
				seconds += s
			} else {
				fmt.Println(">>>too many matches", line)
			}
		}
	}
	// seconds = 90
	minutes += seconds / 60
	seconds = seconds % 60
	fmt.Printf("totally %d:%d mins in %d lines\n", minutes, seconds, count)
	// totally 754:24 mins in 358 lines
}
