package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type empty struct{}

const maxCars = 10

var northBound = make(chan empty, maxCars)
var southBound = make(chan empty, maxCars)

func sendNorth(southCap *bool) {
	if *southCap == true {
		*southCap = false
	}
	fmt.Println("Going north!")
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Println("Made it north!")
	<-northBound
}

func sendSouth(northCap *bool) {
	if *northCap == true {
		*northCap = false
	}
	fmt.Println("Going south!")
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	fmt.Println("Made it south!")
	<-southBound
}

func main() {
	mutex := &sync.Mutex{}
	maxCapNorth := false
	maxCapSouth := false
	for {
		mutex.Lock()
		if len(southBound) == 0 && !maxCapNorth {
			northBound <- empty{}
			go sendNorth(&maxCapSouth)
		}
		if len(northBound) == maxCars && !maxCapNorth {
			fmt.Println("North is maxed outttt!")
			maxCapNorth = true
		}
		if len(northBound) == 0 && !maxCapSouth {
			southBound <- empty{}
			go sendSouth(&maxCapNorth)
		}
		if len(southBound) == maxCars && !maxCapSouth {
			fmt.Println("South is maxed outttt!")
			maxCapSouth = true
		}
		mutex.Unlock()
	}
}
