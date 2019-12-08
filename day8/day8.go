package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Result for data.txt")
	start := time.Now()
	task1 := Task1("data1.txt", 25, 6)
	fmt.Printf("task1: %d \n", task1)
	fmt.Println("task 1 finished in: ", time.Since(start))

	start = time.Now()
	Task2("data1.txt", 25, 6)
	fmt.Println("task 2 finished in: ", time.Since(start))
	fmt.Println("-----------------------------")

	fmt.Println("Result for data2.txt")
	start = time.Now()
	task1data2 := Task1("data2.txt", 25, 6)
	fmt.Printf("task1: %d \n", task1data2)
	fmt.Println("task 1 finished in: ", time.Since(start))

	start = time.Now()
	Task2("data2.txt", 25, 6)
	// fmt.Printf("task2: %d \n", task2data2)
	fmt.Println("task 2 finished in: ", time.Since(start))
	fmt.Println("-----------------------------")
}