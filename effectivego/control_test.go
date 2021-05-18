package effectivego

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

func TestControl(t *testing.T) {

	fmt.Println("flow control examples")

	// for init; condition; post { }
	var sum int
	for i := 0; i < 10; i++ { // no parentheses, but has braces
		sum += i
		// if i++ > 1 { // not compile
		// 	break
		// }
	}
	fmt.Println(sum)

	// There is no do or while loop, only a slightly generalized for;
	// For is Go's "while"
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	// forever
	// Like a C for(;;)
	// for {
	// }

	for key, value := range os.Environ() {
		fmt.Printf("%d=%s\n", key, value)
		if key > 3 {
			break
		}
	}

	// switch is more flexible;
	fmt.Print("Go runs on ")
	// if and switch accept an optional initialization statement like that of for;
	// It's therefore possible—and idiomatic—to write an if-else-if-else chain as a switch.
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
