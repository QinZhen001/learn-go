package main

import (
	"fmt"
	"sync"
	"time"
)

type Goods struct {
	v map[string]int
	m sync.Mutex
	// sync.RWMutex
}

func (g *Goods) Increment(key string, number int) {
	g.m.Lock()
	defer g.m.Unlock()
	fmt.Println("Increment", number)
	g.v[key]++
}

func (g *Goods) Value(key string) int {
	g.m.Lock()
	defer g.m.Unlock()
	fmt.Println("Value")
	return g.v[key]
}

func main() {
	g := Goods{
		v: make(map[string]int),
		m: sync.Mutex{},
	}
	for io := 0; io < 100; io++ {
		go g.Increment("a", io)
	}
	time.Sleep(time.Second * 1)
	fmt.Println("last value", g.Value("a"))
}
