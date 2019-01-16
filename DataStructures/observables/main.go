package main

import (
	"fmt"
	"time"
)

var foreverChannel = make(chan int)

func main() {
	obs := NewObservable(func(o *Observable) {
		go func() {
			i := 0
			for {
				time.Sleep(time.Second * 2)
				o.Next(i)
				i++
			}
		}()
	})

	obs.Subscribe(func(val int) {
		fmt.Print("(1st) Got value: ")
		fmt.Println(val)
	})
	obs.Subscribe(func(val int) {
		fmt.Print("(2nd) Got value: ")
		fmt.Println(val)
	})

	// go func() {
	// 	j := 0
	// 	for {
	// 		time.Sleep(time.Second)
	// 		j++
	// 		if j == 8 {
	// 			thing1.Unsubscribe()
	// 		}
	// 	}
	// }()
	<-foreverChannel
}
