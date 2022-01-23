package main

import "fmt"

func main() {

	c := make(chan int)

	go enviar(c)
	recibir(c)
}

func enviar(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func recibir(c <-chan int) {

	for v := range c {
		fmt.Println(v)
	}

}
