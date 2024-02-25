package structs

import "sync"

type Counter struct {
	mutex sync.Mutex
	count int
}

func NewCounter() *Counter {
	return &Counter{
		count: 0,
		mutex: sync.Mutex{},
	}
}

func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
}

func (c *Counter) Decrement() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count--
}

func (c *Counter) Value() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.count
}
