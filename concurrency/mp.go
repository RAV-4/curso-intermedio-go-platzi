package main

import (
	"fmt"
	"time"
)

func mainMultiplexacion() {
	c1 := make(chan int)
	c2 := make(chan int)

	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go DoSomethingMP(d1, c1, 1)
	go DoSomethingMP(d2, c2, 2)

	for i := 0; i < 2; i++ {
		select {
		case channel1 := <-c1:
			fmt.Println(channel1)
		case channel2 := <-c2:
			fmt.Println(channel2)
		}
	}
}

func DoSomethingMP(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
