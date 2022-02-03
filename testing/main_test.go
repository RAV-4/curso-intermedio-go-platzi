package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	total := sum(1, 2)

	if total != 3 {
		t.Errorf("Sum was incorrect, got %d expected %d", total, 3)
	}
}

func TestSumMasCompleto(t *testing.T) {

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 7, 8},
		{4, 2, 6},
	}

	for _, item := range tables {
		total := sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.n)
		}
	}
}

func TestGetValues(t *testing.T) {
	tables := []struct {
		n int
		a int
		b int
		c int
	}{
		{2, 4, 6, 8},
		{3, 6, 9, 12},
		{10, 20, 30, 40},
	}
	for _, item := range tables {
		g1, g2, g3 := getValues(item.n)
		if g1 != item.a {
			t.Errorf("the double value was incorrect, got %d expected %d", g1, item.a)
		}
		if g2 != item.b {
			t.Errorf("the triple value was incorrect, got %d expected %d", g2, item.b)
		}
		if g3 != item.c {
			t.Errorf("the quad value was incorrect, got %d expected %d", g3, item.c)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}
	for _, item := range tables {
		fib := Fibonacci(item.a)
		if fib != item.n {
			t.Errorf("Fobonacci was incorrect, got %d expected %d", fib, item.n)
		}
	}
}

// go test -cover = comando que nos indica cual es la cobertura de nuestro codigo
// go test -coverprofile=coverage.out = nos genera un archivo con el indicador de cobertura
// go tool cover -func=coverage.out /nos permite ver en la terminal de una forma mas diciente como esta la cobertura
// go tool cover -html=coverage.out /nos permite ver en html la cobertura de codigo
// go test -cpuprofile=cpu.out / genera un archivo que usaremos para valida el rendimiento de nuestro codigo
// go tool pprof cpu.out / nos muestra un reporte del rendimiento del codigo
