package main

import (
	"fmt"
	"math/rand"
)

// channels[n] connects node n and node (n+1)%WORKERS, where WORKERS
// is the total number of nodes.
var channels []chan int

// initialize initializes the channels.
func initialize(WORKERS int) {
	channels = make([]chan int, WORKERS)
	for i := 0; i != WORKERS; i++ {
		channels[i] = make(chan int)
	}
}

// system executes the concurrent system with WORKERS many nodes
func system(WORKERS int, seed int64) {
	printChannel := make(chan string)
	rand.Seed(seed)
	for i := 0; i != WORKERS; i++ {
		go node(WORKERS, i, rand.Intn(10*WORKERS), printChannel)
	}
	console(WORKERS, printChannel)
}

func node(WORKERS, id, value int, printChannel chan string) {
	max := value
	printChannel <- fmt.Sprintf("Node %d holds %d.", id, value)
	if id != 0 {
		max = maxInt(max, <-channels[id])
	}
	channels[(id+1)%WORKERS] <- max
	max = <-channels[id]
	printChannel <- fmt.Sprintf("Node %d now has %d", id, max)
	if id != WORKERS-1 {
		channels[id+1] <- max
	}
}

// console reads in a value and prints it out to the standard input.
func console(WORKERS int, printChannel chan string) {
	for i := 0; i != 2*WORKERS; i++ {
		fmt.Println(<-printChannel)
	}
}

func maxInt(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func main() {
	var seed int64
	fmt.Println("Please enter an integer to be used as a seed to generate pseudo random numbers.")
	fmt.Scan(&seed)

	workers := 0
	for workers < 1 {
		fmt.Println("Please specify the number of nodes. It must be positve.")
		fmt.Scan(&workers)
	}

	initialize(workers)
	system(workers, seed)
}
