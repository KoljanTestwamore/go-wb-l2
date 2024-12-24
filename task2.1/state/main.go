package main

import "fmt"

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) GetCount() int {
	return c.count
}

func main() {
	counter := Counter{}

	counter.Increment()
	counter.Increment()
	counter.Increment()
	counter.Increment()

	fmt.Println(counter.GetCount())
}