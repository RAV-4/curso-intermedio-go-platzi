package main

import "fmt"

func mainWorkers() {
	task := []int{2, 3, 4, 5, 6, 7, 10, 12, 40} //Se definen las tareas
	nWorkers := 3                               //Se definen el numero de workers
	jobs := make(chan int, len(task))
	results := make(chan int, len(task))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results) //Aqui se crean los workers
	}

	for _, value := range task {
		jobs <- value //Con esa se disparan las tareas
	}
	close(jobs)

	for r := 0; r < len(task); r++ {
		<-results
	}
}

func Worker(id int, jobs <-chan int, result chan<- int) { //esta leyendo del canal
	for job := range jobs {
		fmt.Printf("Worker whit id %d started fib whit %d\n", id, job)
		fib := Fibbonacci(job)
		fmt.Printf("Worker with id %d, job %d and fib %d\n", id, job, fib)
		result <- fib
	}
}

func Fibbonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibbonacci(n-1) + Fibbonacci(n-2)
}
