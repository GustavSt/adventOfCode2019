package main

import (
	"flag"
	"fmt"
)

var dayPtr = flag.Int("day", 1, "which day to run")
var taskPtr = flag.Int("task", 1, "run task number one(1) or two(2)")

func main() {
	flag.Parse()
	fmt.Printf("day : %d", *dayPtr)
	fmt.Println()
	fmt.Printf("task: %d", *taskPtr)
	fmt.Println()
	switch *dayPtr {
	default:
		fmt.Println("default case")
	}
}
