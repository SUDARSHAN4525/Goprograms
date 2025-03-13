package main

import "fmt"

type Counter struct {
	count int
}

func (c Counter) getCount() int {
	return c.count
}

func (c *Counter) increment() {
	c.count++
}

func main() {
	c := Counter{count: 5}
	c.increment()
	fmt.Println("Count after increment:", c.getCount())
}
