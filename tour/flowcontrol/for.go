package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Stacking defers
	defer fmt.Println("this first line should be printed at last")
	defer fmt.Println("this second line should be printed previous to last")

	fmt.Println("flow control examples")
	var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//For is Go's "while"
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	// forever
	// for {
	// }

	// switch
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
		//  no need break
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

}
