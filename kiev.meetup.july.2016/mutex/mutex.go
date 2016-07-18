package main

import (
	"fmt"
	"sync"
)

// CounterMux is a cumulative metric that represents a single numerical value that only ever goes up.
type CounterMux struct {
	name  string
	value uint64
	sync.Mutex
}

// Get returns counter value.
func (c *CounterMux) Get() uint64 {
	c.Lock()
	defer c.Unlock()
	return c.value
}

// Add adds delta to counter value.
func (c *CounterMux) Add(delta uint64) {
	c.Lock()
	c.value += delta
	c.Unlock()
}

// NewCounterMux returns new counter that satsfies Metric interface.
func NewCounterMux(name string) *CounterMux {
	return &CounterMux{name: name}
}

var wg sync.WaitGroup

func main() {
	c := NewCounter("")
	wg.Add(2)
	go func() {
		c.Add(14)
		wg.Done()
	}()
	go func() {
		c.Add(28)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(c.Get())
}
