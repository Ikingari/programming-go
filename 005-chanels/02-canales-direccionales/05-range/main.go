package main

import "fmt"

func main() {

	c := make(chan int)

	//enviar
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println("Lo que hay en el canal es:", v)
	}

	fmt.Println("Finalizado")

}
