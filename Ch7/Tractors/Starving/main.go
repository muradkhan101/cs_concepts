package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type empty struct{}

var northBound = make(chan empty, 5)
var southBound = make(chan empty, 5)

func sendNorth() {
	fmt.Println("Going north!")
	time.Sleep(time.Duration(rand.Intn(1200)) * time.Millisecond)
	fmt.Println("Made it north!")
	<-northBound
}

func sendSouth() {
	fmt.Println("Going south!")
	time.Sleep(time.Duration(rand.Intn(1200)) * time.Millisecond)
	fmt.Println("Made it south!")
	<-southBound
}

func main() {
	mutex := &sync.Mutex{}
	for {
		// Originally had goroutine wrap inside of for loop but that caused lockup with mutexes
		mutex.Lock()
		if len(southBound) == 0 {
			northBound <- empty{}
			go func() {
				sendNorth()
			}()
		}
		if len(northBound) == 0 {
			southBound <- empty{}
			go func() {
				sendSouth()
			}()
		}
		mutex.Unlock()
	}
}
