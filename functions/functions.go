package main

import (
	"fmt"
	"time"
)

func mainFunctions() {
	x := 5
	//Funcion anonima
	y := func() int {
		return x + 2
	}() //Los parentesis al final es para indicar la llamada la funcion anonima
	fmt.Println(y)

	c := make(chan int)
	go func() {
		fmt.Println("Starting function")
		time.Sleep(2 * time.Second)
		fmt.Println("End")
		c <- 1
	}()
	<-c
}
