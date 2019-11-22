package runes

import (
	"fmt"
	"unicode/utf8"
)

func RuneUsage() {
	hi := "Hello"
	fmt.Println(hi)

	nh := "你好"
	fmt.Println(nh)
	fmt.Println(len(nh))                    // 6 bytes !
	fmt.Println(utf8.RuneCountInString(nh)) // 2 chars
	nhRunes := []rune(nh)
	fmt.Println(string(nhRunes[1:]))
}
