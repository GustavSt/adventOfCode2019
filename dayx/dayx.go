package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Result for data.txt")
	start := time.Now()
	task1 := Task1("data.txt")
	fmt.Printf("task1: %d \n", task1)
	fmt.Println("task 1 finished in: ", time.Since(start))

	start = time.Now()
	task2 := Task2("data.txt")
	fmt.Printf("task2: %d \n", task2)
	fmt.Println("task 2 finished in: ", time.Since(start))
	fmt.Println("-----------------------------")

	fmt.Println("Result for dataSpotify.txt")
	start = time.Now()
	task1Spotify := Task1("dataSpotify.txt")
	fmt.Printf("task1: %d \n", task1Spotify)
	fmt.Println("task 1 finished in: ", time.Since(start))

	start = time.Now()
	task2Spotify := Task2("dataSpotify.txt")
	fmt.Printf("task2: %d \n", task2Spotify)
	fmt.Println("task 2 finished in: ", time.Since(start))
	fmt.Println("-----------------------------")
}
