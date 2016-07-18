package main

import "sync/atomic"

// Counter is a cumulative metric that represents a single numerical value that only ever goes up.
type Counter struct {
	name  string
	value uint64
}

// Get returns counter value.
func (c *Counter) Get() uint64 {
	return atomic.LoadUint64(&c.value)
}

// Add adds delta to counter value.
func (c *Counter) Add(delta uint64) {
	atomic.AddUint64(&c.value, delta)
}

// NewCounter returns new counter that satsfies Metric interface.
func NewCounter(name string) *Counter {
	return &Counter{name: name}
}
