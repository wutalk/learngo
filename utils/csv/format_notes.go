package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 	in := `first_name,last_name,username
	// "Rob","Pike",rob
	// Ken,Thompson,ken
	// "Robert","Griesemer","gri"
	// `
	// r := csv.NewReader(strings.NewReader(in))

	f := flag.String("f", "safari-annotations-export.csv", "csv file path")
	c := flag.String("c", "Docker in Action, Second Edition", "csv Book Title field")
	flag.Parse()

	fmt.Printf("start format notes of book [%s] in file [%s]\n\n", *c, *f)

	cf, err := os.Open(*f)
	if err != nil {
		log.Fatal("file not found. ", err)
	}
	r := csv.NewReader(bufio.NewReader(cf))

	title := *c
	notes := make(map[int][]string)
	for i := 0; ; {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("failed read line of record", err)
		}

		if title == record[0] {
			// notes = append(notes, record)
			notes[i] = record
			i++
			// fmt.Println(record)
		}
	}
	// fmt.Println(len(record), record)
	l := len(notes)
	fmt.Printf("There are %d notes\n", l)
	preCh := ""
	for i := l - 1; i >= 0; i-- {
		// fmt.Println(notes[i])
		r := notes[i]
		chTitle := r[2]
		if preCh != chTitle {
			fmt.Println("\n", chTitle)
			preCh = chTitle
		}
		fmt.Println(r[7], r[8])
		fmt.Println()
	}

}
