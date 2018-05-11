package main

// Getting deadlocks relatively fast in this version

import (
	"fmt"
	"math/rand"
	"sync"
)

type empty struct{}

const bufferCap = 10

var buffer = make([]int, bufferCap)
var count = 0

func addToBuffer(semaphore chan empty, mutex *sync.Mutex) {
	semaphore <- empty{}

	mutex.Lock()
	buffer[len(semaphore)%bufferCap] = len(semaphore)
	mutex.Unlock()
}
func removeFromBuffer(item int, semaphore chan empty, mutex *sync.Mutex) {
	if item == -1 {
		item = 0
	}
	<-semaphore

	mutex.Lock()
	buffer[item] = -1
	fmt.Println("Removed:", item)
	fmt.Println("Current status:", buffer)
	mutex.Unlock()
}
func main() {
	semaphore := make(chan empty, bufferCap)
	var mutex = &sync.Mutex{}
	// var chanMutex = &sync.Mutex{}
	for i := 0; i < 1000; i++ {
		go func() {
			if r := rand.Float32(); r < 0.5 {
				addToBuffer(semaphore, mutex)
			} else {
				removeFromBuffer(len(semaphore)-1, semaphore, mutex)
			}
		}()
	}
	randChan := make(chan empty)
	<-randChan
}
