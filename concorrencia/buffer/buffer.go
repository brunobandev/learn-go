package main

import "fmt"

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Println("Executou!")
	ch <- 6
}

func main() {
	ch := make(chan int, 3) // buffer de 3 posições
	go rotina(ch)

	fmt.Println(<-ch)
}