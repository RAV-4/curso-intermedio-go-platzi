package inventarioclases

import (
	"fmt"
	"time"
)

func repasoClase2() {
	c := make(chan int)
	go doSomething(c)
	<-c

	g := 25
	fmt.Println("Valor de g", g)
	h := &g
	fmt.Println("Apuntador a g", h)
	fmt.Println("Valor de g, referenciando desde el apuntador", *h)
}

func doSomething(c chan int) {
	time.Sleep(time.Second)
	fmt.Println("Done")
	c <- 1
}
