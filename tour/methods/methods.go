package main

import (
	"fmt"
	"io"
	"math"
	"strings"
)

func main() {
	fmt.Println("methods examples")
	methodsOnTypes()
	interfaces()
	typeAssertions()
	stringers()

	readers()
}
func title(title string) {
	fmt.Printf("\n---------------\n%s\n", title)
}

func methodsOnTypes() {
	title("methodsOnTypes")
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	v.Scale(2)
	fmt.Println(v.Abs())
}

type Vertex struct {
	X, Y float64
}

// With a value receiver, the Scale method operates on a copy of the original Vertex value.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Since methods often need to modify their receiver,
// pointer receivers are more common than value receivers.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func interfaces() {
	title("interfaces")
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())

	// No null pointer exception
	// If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
	// posibly because there is "zero" value, like int is 0, string is empty "", struct is just a collection of fields
	var vnil Vertex
	a = &vnil
	fmt.Printf("(%v, %T)\n", a, a)
	fmt.Println(a.Abs())

	// ! however, Calling a method on a nil interface is a run-time error
	// because there is no type inside the interface tuple to indicate which concrete method to call.
	var b Abser
	fmt.Printf("(%v, %T)\n", b, b)
	// un-comment following line and run to see the error
	// b.Abs()
}

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func typeAssertions() {
	title("Type assertions")
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	fmt.Println(f)
}

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func stringers() {
	title("Stringers")
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %s\n", name, ip)
	}
}

func readers() {
	title("Readers")
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
