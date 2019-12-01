package main

import (
	"strconv"

	"../utils"
)

//Task1 do stuff
func Task1(dataFile string) int {
	data, _ := utils.SplitData(1, dataFile)
	res := 0
	for _, module := range data {
		i, _ := strconv.Atoi(module)
		res += calcFuel1(i)
	}
	return res
}

func calcFuel1(moduleMass int) int {
	d3 := moduleMass / 3
	res := d3 - 2
	return res
}

//Task2 do stuff
func Task2(dataFile string) int {
	data, _ := utils.SplitData(1, dataFile)
	res := 0
	for _, module := range data {
		i, _ := strconv.Atoi(module)
		res += calcFuel2(i)
	}
	return res
}

func calcFuel2(mass int) int {
	if mass <= 0 {
		return 0
	}
	fuel := calcFuel1(mass)
	if fuel <= 0 {
		fuel = 0
	}
	fuel += calcFuel2(fuel)
	return fuel
}
