package main

// This version has an extra mutex
// to surround the addition and removal from
// the semaphore channel

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type empty struct{}

const bufferCap = 10

var buffer = make([]int, bufferCap)
var fullSemaphore = make(chan empty, bufferCap-1)
var emptySemaphore = make(chan empty, bufferCap-1)
var mutex = &sync.Mutex{}

func producer() {
	go func() {
		for {
			if rand.Float32() < 0.5 {
				<-emptySemaphore
				length := len(fullSemaphore)

				rand := rand.Intn(100) + 1
				mutex.Lock()
				buffer[length] = rand
				fmt.Println("Buffer len is", length)
				mutex.Unlock()
				fmt.Println("Put in a ", rand)

				fullSemaphore <- empty{}
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
				<-fullSemaphore
				length := len(fullSemaphore)
				mutex.Lock()
				num := buffer[length]
				buffer[len(fullSemaphore)] = -1
				fmt.Println("Buffer len is", length)
				mutex.Unlock()
				fmt.Println("Pulled out a", num)

				emptySemaphore <- empty{}
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
	for i := 0; i < bufferCap-1; i++ {
		emptySemaphore <- empty{}
	}
	for i := 0; i < 1; i++ {
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
