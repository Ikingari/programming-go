package main

import "fmt"

func main() {
	// Buffered channel (canal sin búfer)
	ca := make(chan int, 2)
	ca <- 42
	ca <- 43

	fmt.Println(<-ca)
	fmt.Println(<-ca)
}
