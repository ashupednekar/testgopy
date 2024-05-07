package pycall

import (
	"fmt"
	"sync"
	"time"
)

type Caller struct {
	Ch   chan string
	Iter int
}

func NewCaller(n int) Caller {
	c := Caller{
		Ch:   make(chan string, n),
		Iter: n,
	}
	return c
}

func (c *Caller) run(f func()) {
	fmt.Printf("Worker starting\n")
	time.Sleep(time.Second)
	f()
	c.Ch <- "hi"
	fmt.Printf("Worker done\n")
}

func (c *Caller) Start(f func()) {
	var wg sync.WaitGroup

	for range c.Iter {
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.run(f)
		}()
	}

	go func() {
		wg.Wait()
		close(c.Ch)
	}()
}
