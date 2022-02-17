package main

import "fmt"

func channels() {
	//c := make(chan int)    //Canal sin buffer
	c := make(chan int, 3) //Canal con buffer

	c <- 1
	c <- 20
	c <- 334

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
