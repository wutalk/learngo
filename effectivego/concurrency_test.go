package effectivego

import (
	"fmt"
	"runtime"
	"testing"
)

func TestConcurrency(t *testing.T) {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(0))

	silkCaves := []string{"cave1", "cave2", "cave3", "cave4"}
	c := make(chan SpiderSpirit, len(silkCaves))
	for i := 0; i < len(silkCaves); i++ {
		go search(c, silkCaves[i])
	}
	for s := range c {
		fmt.Println(s)
	}
}

func search(c chan SpiderSpirit, location string) {
	fmt.Printf("searching %s...\n", location)
	s := fmt.Sprintf("%s#SpiderSpirit", location)
	c <- SpiderSpirit(s)
}

type SpiderSpirit string

func TestVariablesScope(t *testing.T) {
	s := "hello"
	c := make(chan int)
	go func() {
		fmt.Println(1, s)
		c <- 1
	}()

	// s = "world"
	go func(x string) {
		fmt.Println(2, x)
		c <- 1
	}(s)

	s = "china"

	<-c
	<-c
}

func TestVariablesScopeRefer(t *testing.T) {
	s := "hello"
	rs := &s
	c := make(chan int)
	go func() {
		fmt.Println(1, *rs)
		c <- 1
	}()

	// s = "world"
	go func(x *string) {
		fmt.Println(2, *x)
		c <- 1
	}(rs)

	s = "china"

	<-c
	<-c
}

/*
Go programs express error state with error values. (with multiple return values)
C, bash, etc with int values (a single value mix normal return with error return)
*/

// write our own Error type
// PathError records an error and the operation and
// file path that caused it.
type PathError struct {
	Op   string // "open", "unlink", etc.
	Path string // The associated file.
	Err  error  // Returned by the system call.
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func TestCustomError(t *testing.T) {
	err := foo()
	fmt.Printf("fail to foo: %v\n", err)
	fmt.Printf("fail to foo: %v\n", fmt.Errorf("unlink /tmp/data.log due to %v", fmt.Errorf("no space")))
}

func foo() PathError {
	//
	return PathError{Op: "unlink", Path: "/tmp/data.log", Err: fmt.Errorf("no space")}
}

/*
	resp, err := c.Do(req)
	// resp, err := http.Get(link)
	if err != nil {
		if err, ok := err.(*url.Error); ok {
			fmt.Printf("is timeout: %v\n", err.Timeout())
		}
		log.Fatalf("failed to request url %s\n%#v", link, err)
	}
*/

func TestPanicRecover(t *testing.T) {
	c := make(chan int)
	f(c)
	<-c
}

func f(c chan int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("work failed:", err)
		}
		c <- 1
	}()
	g()
}

func g() {
	fmt.Println("panic")
	panic("panic in g")
}
