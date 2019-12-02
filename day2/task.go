package main

import (
	"fmt"

	"../utils"
)

//Task1 do stuff
func Task1(dataFile string) int {
	data, _ := utils.SplitOnChar(2, dataFile, ",")
	intData := utils.ConvertToInt(data)

	return intCodeProgram(12, 2, intData)
}

//Task2 do stuff
func Task2(dataFile string) int {
	data, _ := utils.SplitOnChar(2, dataFile, ",")
	intData := utils.ConvertToInt(data)

	noun := 0
	verb := 0
	res := 0
	expected := 19690720

	for res != expected {
		if noun == 100 {
			noun = 0
			verb++
		}
		res = intCodeProgram(noun, verb, intData)
		intData = utils.ConvertToInt(data)
		if res == expected {
			break
		}
		noun++
	}

	fmt.Printf("noun: %d. verb %d", noun, verb)
	return noun*100 + verb
}

func intCodeProgram(noun int, verb int, intData []int) int {
	index := 0
	currentOpcode := intData[index]
	intData[1] = noun
	intData[2] = verb
	for currentOpcode != 99 {
		switch currentOpcode {
		case 1:
			intData[intData[index+3]] = intData[intData[index+1]] + intData[intData[index+2]]
		case 2:
			intData[intData[index+3]] = intData[intData[index+1]] * intData[intData[index+2]]
		}
		index = index + 4
		currentOpcode = intData[index]
	}
	return intData[0]
}
