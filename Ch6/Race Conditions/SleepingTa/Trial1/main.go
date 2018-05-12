package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numStudents = 6
	numChairs   = 3
)

var mutex = &sync.Mutex{}
var taStatus = make(chan int, 1)
var waitingRoom = make(chan int, numChairs+1)

func helpStudent(s *student) {
	fmt.Println(s.id, "is receiving help")
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	mutex.Lock()
	taStatus <- 1
	<-waitingRoom
	mutex.Unlock()
}

type student struct {
	id int
}

func (s *student) program() {
	fmt.Println(s.id, "is programming!")
	time.Sleep(time.Duration(rand.Intn(750)) * time.Millisecond)
	fmt.Println(s.id, "is done programming!")
}

func (s *student) getHelp() {
	fmt.Println(s.id, "is getting help!")
	if len(waitingRoom) == numChairs+1 {
		fmt.Println("Waiting room is full for", s.id)
	} else {
		mutex.Lock()
		waitingRoom <- 1
		mutex.Unlock()
		helpStudent(s)
		mutex.Lock()
		<-taStatus
		mutex.Unlock()
		fmt.Println(s.id, "is done getting help!")
	}
}

func main() {

	for i := 0; i < numStudents; i++ {
		go func(i int) {
			s := student{i}
			for {
				s.program()
				s.getHelp()
			}
		}(i)
	}
	v := make(chan int)
	<-v
}
