package main

import (
	"fmt"
	"time"
)

type A struct {
	id int
}

type B struct {
	A
	name string
}

func (a *A) show() {
	fmt.Printf("a id=%d\n", a.id)
}

func (b *B) show() {
	fmt.Printf("b id=%d name=%d\n", b.id, b.name)
	b.A.show()
}

func Run() {
	var b B = B{A: A{id: 1}, name: "b"}
	b.show()
}
func main() {
	task_start := time.Now()
	Run()
	time_cost := time.Since(task_start)
	cost := uint64(float64(time_cost) / float64(time.Millisecond))
	fmt.Printf("=========== end time pass = %dms\n", cost)
}
