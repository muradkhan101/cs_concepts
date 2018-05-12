package main

import (
	"fmt"
	"sync"
	"time"
)

const PHILOS = 5

type state int

const (
	awake = iota + 1
	sleeping
	eating
)

type philosopher struct {
	id    int
	state state
	life  int
}

func (p *philosopher) wake() {
	p.state = awake
	p.life--
	if p.life == 500 {
		fmt.Println("Philosopher", p.id, "is halfway dead!")
	} else if p.life == 200 {
		fmt.Println("Philosopher", p.id, "is pretty close to death!")
	} else if p.life < 50 {
		fmt.Println("Philosopher", p.id, "has", p.life, "life left!")
	}
	if !spoons[p.id] && !spoons[(p.id+1)%(PHILOS)] {
		mutex.Lock()
		spoons[p.id] = true
		spoons[(p.id+1)%(PHILOS)] = true
		mutex.Unlock()
		go func() { p.eat() }()
	} else {
		go func() { p.sleep() }()
	}
}

func (p *philosopher) sleep() {
	p.state = sleeping
	time.Sleep(time.Second)
	go func() { p.wake() }()
}

func (p *philosopher) eat() {
	p.state = eating
	p.life += 500
	fmt.Println("Philosopher", p.id, "eating", ps)
	time.Sleep(time.Millisecond * 250)
	mutex.Lock()
	spoons[p.id] = false
	spoons[(p.id+1)%(PHILOS)] = false
	mutex.Unlock()
	go func() { p.sleep() }()
}

var spoons [PHILOS]bool
var ps [PHILOS]philosopher
var mutex = &sync.Mutex{}

func main() {
	go func() {
		for i := 0; i < PHILOS; i++ {
			ps[i] = philosopher{i, sleeping, 1000}
			fmt.Println(i, ps)
			ps[i].wake()
		}
	}()
	t := make(chan int)
	<-t
}
