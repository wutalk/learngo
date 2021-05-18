package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Hello, 世界", math.Pi)
	fmt.Println("My favorite number is", rand.Intn(100))
	var total int
	x, y := addAndMulti(5, 3)
	fmt.Println(x, y)
	total = x + y
	fmt.Println(total)

	// zero values
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// The expression T(v) converts the value v to the type T.
	// Unlike in C, in Go assignment between items of different type requires an explicit conversion.
	i = 42
	f = float64(i)
	u := uint(f)
	fmt.Printf("%v, %v, %v", i, f, u)

	//

}

func addAndMulti(a, b int) (int, int) {
	return a + b, a * b
}
