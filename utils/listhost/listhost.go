package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"text/scanner"
)

// grep 'Host ' C:\Users\owu\.ssh\config | cut -d ' ' -f2 | sort
func main() {
	home := os.Getenv("HOME")
	// fmt.Println(home)
	sep := string(os.PathSeparator)
	// f := "/home/owen/.ssh/config"
	f := fmt.Sprintf("%s%s.ssh%sconfig", home, sep, sep)
	// fmt.Println(f)
	fmt.Println("configured host list")
	cf, err := os.Open(f)
	if err != nil {
		log.Fatal("file not found. ", err)
	}
	defer cf.Close()

	r := bufio.NewReader(cf)

	var s scanner.Scanner
	s.Init(r)
	var prev string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		// fmt.Printf("%s: [%s]\n", s.Position, s.TokenText())
		txt := s.TokenText()
		if prev == "Host" {
			fmt.Println(txt)
		}
		prev = txt
	}
}
