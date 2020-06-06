package utils

import (
	"sync"
)

// Counter is go routine safe counter used to count events.
// We add Read/Write Mutex to prevent race conditions.
type Counter struct {
	sync.RWMutex
	counter map[string]uint64
}

// NewCounter creates and returns a new Counter
func NewCounter() *Counter {
	return &Counter{
		counter: make(map[string]uint64),
	}
}

// Incr increments counter for specified key
func (c *Counter) Incr(k string) {
	c.Lock()
	c.counter[k]++
	c.Unlock()
}

// Val returns current value for specified key
func (c *Counter) Val(k string) uint64 {
	var v uint64
	c.RLock()
	defer c.RUnlock()
	v = c.counter[k]
	return v
}

// Items returns all the counter items
func (c *Counter) Items() map[string]uint64 {
	c.RLock()
	items := make(map[string]uint64, len(c.counter))
	for k, v := range c.counter {
		items[k] = v
	}
	c.RUnlock()
	return items
}
