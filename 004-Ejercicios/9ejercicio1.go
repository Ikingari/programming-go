package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("El número de Procesadores es:", runtime.NumCPU())
	fmt.Println("El número de GoRutines es:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hola desde la GoRutine 1")
		wg.Done()
	}()
	go func() {

		fmt.Println("Hola desde la GoRutine 2")
		wg.Done()
	}()

	fmt.Println("El número de Procesadores es:", runtime.NumCPU())
	fmt.Println("El número de GoRutines es:", runtime.NumGoroutine())
	fmt.Println("A punto de finalizar")
	wg.Wait()
	fmt.Println("El número de Procesadores es:", runtime.NumCPU())
	fmt.Println("El número de GoRutines es:", runtime.NumGoroutine())

}
