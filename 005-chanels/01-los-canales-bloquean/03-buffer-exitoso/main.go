package main

import "fmt"

func main() {
	// Buffered channel (canal sin búfer)
	ca := make(chan int, 1)
	ca <- 42

	fmt.Println(<-ca)
}
