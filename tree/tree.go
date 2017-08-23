package main

import (
	"fmt"
	"math/rand"
)

// channels[n] connects node n and its parent if it exists.
var channels []chan int

// initialize initializes the channels.
func initialize(WORKERS int) {
	channels = make([]chan int, WORKERS)
	for i := 0; i != WORKERS; i++ {
		channels[i] = make(chan int)
	}
}

// system executes the concurrent system with WORKERS many nodes.
func system(WORKERS int, seed int64) {
	rand.Seed(seed)
	printChannel := make(chan string)
	for i := 0; i != WORKERS; i++ {
		go node(WORKERS, i, rand.Intn(WORKERS*10), printChannel)
	}
	console(WORKERS, printChannel)
}

func node(WORKERS, id, value int, printChannel chan string) {
	leftChild := 2*id + 1
	rightChild := 2*id + 2
	max := value
	printChannel <- fmt.Sprintf("Node %d holds %d.", id, value)
	if leftChild < WORKERS {
		max = maxInt(max, <-channels[leftChild])
	}
	if rightChild < WORKERS {
		max = maxInt(max, <-channels[rightChild])
	}

	if id != 0 {
		channels[id] <- max
		max = <-channels[id]
	}
	printChannel <- fmt.Sprintf("Node %d now has %d.", id, max)

	if leftChild < WORKERS {
		channels[leftChild] <- max
	}
	if rightChild < WORKERS {
		channels[rightChild] <- max
	}
}

// console reads in a value and prints it out to the standard input.
func console(WORKERS int, in chan string) {
	for i := 0; i != 2*WORKERS; i++ {
		fmt.Println(<-in)
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
