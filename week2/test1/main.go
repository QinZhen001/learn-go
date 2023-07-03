package main

import "fmt"

type Run interface {
	Running()
}

type Swim interface {
	Swimming()
}

type Sport interface {
	Run
	Swim
}

type Boy struct {
	Name string
}

func (b *Boy) Running() {
	fmt.Println(b.Name + "在跑步。。。。")
}

func (b *Boy) Swimming() {
	fmt.Println(b.Name + "在游泳。。。。")
}

func GoSports(s Sport) {
	s.Running()
	s.Swimming()
}

func main() {
	b := Boy{Name: "111"}
	GoSports(&b)
}
