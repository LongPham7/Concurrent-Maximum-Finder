package main

import (
	"fmt"
	"math/rand"
)

// system executes the concurrent system consisting of WORKERS many nodes
// and a single coordinator.
func system(WORKERS int, seed int64) {
	rand.Seed(seed)
	fromN := make(chan int)
	toN := make(chan int)
	printChannel := make(chan string)
	for i := 0; i != WORKERS; i++ {
		go node(i, rand.Intn(10*WORKERS), toN, fromN, printChannel)
	}
	go coordinator(WORKERS, fromN, toN)
	console(WORKERS, printChannel)
}

// coordinator keeps track of the maximum of the values that have been
// sent by the nodes thus far.
func coordinator(WORKERS int, in, out chan int) {
	max := 0
	for i := 0; i != WORKERS; i++ {
		max = maxInt(max, <-in)
	}
	for i := 0; i != WORKERS; i++ {
		out <- max
	}
}

func node(id, value int, in, out chan int, printChannel chan string) {
	printChannel <- fmt.Sprintf("Node %d holds %d", id, value)
	out <- value
	max := <-in
	printChannel <- fmt.Sprintf("Node %d now has %d", id, max)
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

	system(workers, seed)
}
