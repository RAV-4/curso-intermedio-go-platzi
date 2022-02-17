package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int, 2)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		c <- 1 //Con este canal podemos controlar que se ejecuten solo dos go routines al tiempo, debido al tamaño de buffer
		wg.Add(1)
		go doSomethingSem(i, &wg, c)
	}
}

func doSomethingSem(i int, group *sync.WaitGroup, c chan int) {
	defer group.Done()
	fmt.Printf("Id %d started\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("Id %d finished\n", i)
	<-c
}
