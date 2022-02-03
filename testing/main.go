package main

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
