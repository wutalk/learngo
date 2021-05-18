package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Concurrency examples")

	// goroutines() // uncomment to see
	channels()
	ranges()
	selectChannel()
	defaultSelection()
	syncMutexDemo()
}

func title(title string) {
	fmt.Printf("\n---------------\n%s\n", title)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutines() {
	go say("+ world")
	say("-  hello")
}

func channels() {
	title("Channels")

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 // panic deadlock
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func ranges() {
	title("ranges")
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func selectChannel() {
	title("selectChannel")

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 12; i++ {
			fmt.Print(<-c, " ")
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func defaultSelection() {
	title("defaultSelection")

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Print("tick.")
		case <-boom:
			fmt.Print("BOOM!")
			return
		default:
			fmt.Print(" * ")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func syncMutexDemo() {
	title("syncMutexDemo")

	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex // uncomment this: fatal error: concurrent map writes
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	// defer c.mux.Unlock() // defer unlock after lock to avoid forget unlock
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}
