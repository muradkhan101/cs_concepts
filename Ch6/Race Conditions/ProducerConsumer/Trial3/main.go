package main

// This version has the entire while loop
// wrapped in a go function to see if that helps
// and creates multiple producers and consumers

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type empty struct{}

const bufferCap = 10

var buffer = make([]int, bufferCap)
var semaphore = make(chan empty, bufferCap-1)
var mutex = &sync.Mutex{}

func producer() {
	go func() {
		for {
			if rand.Float32() < 0.5 {
				rand := rand.Intn(100) + 1
				semaphore <- empty{}
				mutex.Lock()
				buffer[len(semaphore)] = rand
				fmt.Println("Buffer len is", len(semaphore))
				mutex.Unlock()
				fmt.Println("Put in a ", rand)
			} else {
				time.Sleep(time.Second)
			}
		}
	}()
}

func consumer() {
	go func() {
		for {
			if rand.Float32() < 0.5 {
				<-semaphore
				mutex.Lock()
				num := buffer[len(semaphore)]
				buffer[len(semaphore)] = -1
				fmt.Println("Buffer len is", len(semaphore))
				mutex.Unlock()
				fmt.Println("Pulled out a", num)
				if num == 0 || num == -1 {
					fmt.Println("OMG THE BUFFERS ARE SCREWQWWWWED!!!")
				}
			} else {
				time.Sleep(time.Second)
			}
		}
	}()
}

func main() {
	for i := 0; i < 3; i++ {
		go func() {
			consumer()
		}()
		go func() {
			producer()
		}()
	}

	randChan := make(chan empty)
	<-randChan
}
