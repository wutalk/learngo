package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("types examples")
	pointers()

	structDemo()

	arrays()

	slices()
	sliceLenCap()
	makeSlice()

	maps()

	functionValues()
	functionClosures()
	fibonacciClosure()
}

func pointers() {
	title("pointers")

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	fmt.Printf("p type is %v\n", p)
	*p = 21        // set i through the pointer
	fmt.Println(i) // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

// A struct is a collection of fields.
type Vertex struct {
	X int
	Y int
}

func structDemo() {
	title("struct")
	var vtx Vertex
	vtx = Vertex{}
	vtx = Vertex{Y: 100}
	fmt.Printf("vtx=%v\n", vtx)
	vtx.X = 29
	fmt.Printf("vtx=%v\n", vtx)

	p := &vtx
	fmt.Printf("p.X=%v, (*p).X=%v\n", p.X, (*p).X) // same as (*p).X
}

func title(title string) {
	fmt.Printf("\n---------------\n%s\n", title)
}

func arrays() {
	title("arrays")

	// An array's length is part of its type, so arrays cannot be resized.
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slices() {
	title("slices")

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	/*
		A slice does not store any data, it just describes a section of an underlying array.
		Changing the elements of a slice modifies the corresponding elements of its underlying array.
	*/
	s[0] = 30
	fmt.Println(s)
	fmt.Println(primes)

	a := [3]bool{true, true, false}
	fmt.Println(a)

	sl := []bool{true, true, false}
	// same as
	// sl := a[:]
	fmt.Println(sl)

}

func sliceLenCap() {
	title("sliceLenCap")

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	// Extend its length. again
	// Drop is permernent, cannot go back to cap 6
	s = s[:4]
	printSlice(s)

	// but you can append, which doubles original capacity
	s = append(s, 99)
	printSlice(s)

	for i, v := range s {
		fmt.Println(i, v)
	}

	var nilSlice []int
	printSlice(nilSlice)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func makeSlice() {
	title("makeSlice")

	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)
}

func maps() {
	title("maps")

	var m map[string]Vertex

	// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
	// m["cd"] = Vertex{} // panic: assignment to entry in nil map

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40, -74,
	}
	fmt.Println(m["Bell Labs"])

	d := map[string]string{} // empty map can be added
	d["hello"] = "world"
	fmt.Println(d)

	e := map[string]string{
		"hello": "world",
		"foo":   "bar"}
	fmt.Println(e)
	for k, v := range e {
		fmt.Printf("%s=%s\n", k, v)
	}

	delete(e, "foo")
	fmt.Println(e)

	elem, ok := e["foo"]
	if ok {
		fmt.Printf("e=%v\n", elem)
	} else {
		fmt.Println("not found")
	}
}

func functionValues() {
	title("functionValues")

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func functionClosures() {
	title("functionClosures")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func fibonacciClosure() {
	title("fibonacciClosure")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
func fibonacci() func() int {
	f, g := 1, 0
	return func() int {
		f, g = g, f+g
		return f
	}
}
