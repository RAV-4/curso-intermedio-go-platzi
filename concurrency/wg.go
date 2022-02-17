package main

import (
	"fmt"
	"sync"
	"time"
)

func wg() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}
	wg.Wait() //Espera que el contador del wait group llege a cero
}

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Started ", i)
	time.Sleep(2 * time.Second)
	fmt.Println("End ", i)
}
