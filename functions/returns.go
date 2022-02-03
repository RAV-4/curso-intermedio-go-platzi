package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func getValues(x int) (doble int, triple int, cuadruple int) {
	doble = 2 * x
	triple = 3 * x
	cuadruple = 4 * x
	return
}
func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(getValues(2))
}
