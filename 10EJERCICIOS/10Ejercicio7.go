package main

import "fmt"

func main() {

	c := make(chan int)
	for i := 0; i < 10; i++ {
		go enviar(c)

	}

	recibir(c)
}

func enviar(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}

}

func recibir(c <-chan int) {

	for i := 0; i < 100; i++ {
		fmt.Println(i, <-c)
	}

}
