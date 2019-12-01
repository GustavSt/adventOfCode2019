package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	task1 := Task1("data.txt")
	fmt.Printf("task1: %d \n", task1)
	fmt.Println("task 1 finished in: ", time.Since(start))

	start = time.Now()
	task2 := Task2("data.txt")
	fmt.Printf("task2: %d \n", task2)
	fmt.Println("task 2 finished in: ", time.Since(start))
}
