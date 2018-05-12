package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

func isSafe(need *Matrix, allocation *Matrix, a Vector) bool {
	work := a
	var finish = make([]bool, len(need.Entries))
	for i := 0; i < len(finish); i++ {
		if finish[i] == false && need.Entries[i].LessThanEqual(&work) {
			work = *work.Add(&allocation.Entries[i])
			finish[i] = true
			i = 0
		}
	}
	for i := 0; i < len(finish); i++ {
		if finish[i] == false {
			return false
		}
	}
	return true
}

func release(need *Matrix, allocation *Matrix, pid int, req *Vector, available *Vector) {
	fmt.Println("RELEASING:", pid, req)
	mutex.Lock()
	newAllocation := Matrix{make([]Vector, 0), allocation.Vlen}
	for i := 0; i < len(allocation.Entries); i++ {
		if i == pid {
			newAllocation.AddEntry(allocation.Entries[i].Subtract(req))
		} else {
			newAllocation.AddEntry(&allocation.Entries[i])
		}
	}
	mutex.Unlock()
	newNeed := Matrix{make([]Vector, 0), need.Vlen}
	mutex.Lock()
	for i := 0; i < len(allocation.Entries); i++ {
		if i == pid {
			newNeed.AddEntry(need.Entries[i].Add(req))
		} else {
			newNeed.AddEntry(&need.Entries[i])
		}
	}
	newAvailable := *available.Add(req)
	mutex.Unlock()
	mutex.Lock()
	available = &newAvailable
	allocation = &newAllocation
	need = &newNeed
	mutex.Unlock()
}

func request(need *Matrix, allocation *Matrix, pid int, req *Vector, available *Vector) {
	if req.LessThanEqual(&need.Entries[pid]) {
		for !req.LessThanEqual(available) {
			fmt.Println("GOING TO SLEEP:", pid)
			time.Sleep(time.Second)
		}
		mutex.Lock()
		newAvailable := *available.Subtract(req)
		mutex.Unlock()

		newAllocation := Matrix{make([]Vector, 0), allocation.Vlen}
		mutex.Lock()
		for i := 0; i < len(allocation.Entries); i++ {
			if i == pid {
				newAllocation.AddEntry(allocation.Entries[i].Add(req))
			} else {
				newAllocation.AddEntry(&allocation.Entries[i])
			}
		}
		mutex.Unlock()
		newNeed := Matrix{make([]Vector, 0), need.Vlen}
		mutex.Lock()
		for i := 0; i < len(allocation.Entries); i++ {
			if i == pid {
				newNeed.AddEntry(need.Entries[i].Subtract(req))
			} else {
				newNeed.AddEntry(&need.Entries[i])
			}
		}
		fmt.Println("\nnewAvailable:", newAvailable,
			"\noldAvailable", *available)
		fmt.Println("\nnewAllocation:", newAllocation,
			"\noldAllocation:", *allocation)
		fmt.Println("\nnewNeed:", newNeed,
			"\noldNeed:", *need, "for process:", pid)
		mutex.Unlock()
		if isSafe(&newNeed, &newAllocation, newAvailable) {
			fmt.Println("ALLOCATING RESOURCES FOR REQUEST:", pid)
			mutex.Lock()
			available = &newAvailable
			allocation = &newAllocation
			need = &newNeed
			mutex.Unlock()

			go func() {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
				release(need, allocation, pid, req, available)
			}()

		} else {
			fmt.Println("UNSAFE STATE:", newAvailable, newAllocation, newNeed, req)
		}
	} else {
		fmt.Println("LIAR PROCESS!", pid, "REQUESTING MORE THAN IT NEEDS")
	}
}

func main() {
	// 6 different resource types available
	available := Vector{[]int{10, 5, 7, 4, 2, 12}}
	// 5 processes competing for resources
	max := Matrix{[]Vector{
		Vector{[]int{4, 2, 5, 1, 0, 8}},
		Vector{[]int{6, 3, 3, 2, 2, 5}},
		Vector{[]int{7, 1, 4, 2, 1, 3}},
		Vector{[]int{1, 1, 2, 1, 0, 3}},
		Vector{[]int{3, 2, 2, 3, 1, 7}}}, 6}
	allocation := Matrix{[]Vector{
		Vector{[]int{1, 0, 1, 0, 0, 2}},
		Vector{[]int{0, 1, 0, 0, 0, 1}},
		Vector{[]int{1, 0, 1, 0, 1, 0}},
		Vector{[]int{0, 0, 1, 1, 0, 2}},
		Vector{[]int{0, 1, 0, 1, 0, 1}}}, 6}
	need := *max.Subtract(&allocation)
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1250)))
		go func() {
			index := rand.Intn(len(allocation.Entries))
			mutex.Lock()
			fmt.Println("SENDING TO RANDOMVECTOR", need.Entries[index])
			req := RandomVector(need.Entries[index])
			mutex.Unlock()
			fmt.Println("Requesting:", req, "\nFor process:", index)
			request(&need, &allocation, index, req, &available)
		}()
	}
}
