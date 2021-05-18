package effectivego

import (
	"fmt"
	"math"
	"os"
	"testing"
)

func TestFunctions(t *testing.T) {

	title("A method is a function with a special receiver argument.")
	v := Vertex{3, 4}
	fmt.Println(Abs(v))  // function
	fmt.Println(v.Abs()) // method

	title("Functions: 3 main differences")

	title("Multiple return values")
	// 1) improved from C's error reporting style.
	//    returned error is like exception in java, while panic is like Error in java
	// 2) useful when need a one time struct(object in java)
	bs, err := os.ReadFile("/home/owen/testdata/note.txt")
	if err != nil {
		fmt.Printf("failed to read file %v\n", err)
	}
	fmt.Println(string(bs))

	title("Defer")
	fmt.Println(deferDemo())

}

func deferDemo() int {
	// Stacking defers
	i := 1
	defer fmt.Println(i, "first defer should be printed at last")
	i++
	defer fmt.Println(i, "second defer should be printed previous to last")

	fmt.Println("doing complex job...")

	var err error
	defer commitOrRollback(err, "nil err")
	err = fmt.Errorf("oh err")
	defer commitOrRollback(err, "nonnil err")

	title("End")

	return i
}

func commitOrRollback(err error, msg string) {
	if err != nil {
		fmt.Println("rollback", msg)
	} else {
		fmt.Println("commit", msg)
	}
}

func title(title string) {
	fmt.Printf("\n====\t%s\t====\n", title)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
