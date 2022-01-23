package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("El sistema operativo es: ", runtime.GOOS)
	fmt.Println("La arquitectura de este equipo es: ", runtime.GOARCH)
}
