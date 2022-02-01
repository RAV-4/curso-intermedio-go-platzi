package inventarioclases

import (
	"fmt"
	"strconv"
)

func repasoClase1() {
	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)
	myValue, err := strconv.ParseInt("7s", 0, 64)
	// Los errores en go se manejan de manera explicita
	if err != nil {
		fmt.Println("\n", err)
	} else {
		fmt.Println(myValue)
	}

	//Map
	m := make(map[string]int)
	m["key"] = 6
	fmt.Println(m["key"])

	//Slice
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Printf("indice %d valor %d\n", index, value)
	}
	s = append(s, 6)
	for index, value := range s {
		fmt.Printf("indice %d valor %d\n", index, value)
	}
}
