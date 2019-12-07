package main

import (
	"fmt"
	"strconv"

	"../utils"
)

//Task1 do stuff
func Task1(dataFile string) []int {
	data, _ := utils.SplitOnChar(5, dataFile, ",")
	intData := utils.ConvertToInt(data)

	outputs := intCodeProgram(1, intData)
	return outputs
}

//Task2 do stuff
func Task2(dataFile string, input int) []int {
	data, _ := utils.SplitOnChar(5, dataFile, ",")
	intData := utils.ConvertToInt(data)

	return intCodeProgram(input, intData)
}

func intCodeProgram(input int, intData []int) []int {
	index := 0
	opInst := GetOpInstructions(intData[index])
	outputs := []int{}
	for opInst.opCode != 99 {
		switch opInst.opCode {
		case 1:
			first := getVal(opInst.first, intData, index+1)
			second := getVal(opInst.second, intData, index+2)
			intData[intData[index+3]] = first + second
			index = index + 4
		case 2:
			first := getVal(opInst.first, intData, index+1)
			second := getVal(opInst.second, intData, index+2)
			intData[intData[index+3]] = first * second
			index = index + 4
		case 3:
			intData[intData[index+1]] = input
			index = index + 2
		case 4:
			output := getVal(opInst.first, intData, index+1)
			fmt.Printf("output: %d \n", output)
			outputs = append(outputs, output)
			index = index + 2
		case 5:
			val := getVal(opInst.first, intData, index+1)
			if val != 0 {
				index = getVal(opInst.second, intData, index+2)
			} else {
				index = index + 3
			}
		case 6:
			val := getVal(opInst.first, intData, index+1)
			if val == 0 {
				index = getVal(opInst.second, intData, index+2)
			} else {
				index = index + 3
			}
		case 7:
			val1 := getVal(opInst.first, intData, index+1)
			val2 := getVal(opInst.second, intData, index+2)
			if val1 < val2 {
				intData[intData[index+3]] = 1
			} else {
				intData[intData[index+3]] = 0
			}
			index = index + 4
		case 8:
			val1 := getVal(opInst.first, intData, index+1)
			val2 := getVal(opInst.second, intData, index+2)
			if val1 == val2 {
				intData[intData[index+3]] = 1
			} else {
				intData[intData[index+3]] = 0
			}
			index = index + 4
		}
		opInst = GetOpInstructions(intData[index])
	}
	return outputs
}

//OpInstruct cs
type OpInstruct struct {
	opCode, first, second, third int
}

//GetOpInstructions c
func GetOpInstructions(instructions int) OpInstruct {
	iStr := strconv.Itoa(instructions)
	if len(iStr) == 1 || instructions == 99 {
		return OpInstruct{opCode: instructions, first: 0, second: 0, third: 0}
	}

	opIns, _ := strconv.Atoi(iStr[len(iStr)-2:])
	opResult := OpInstruct{opCode: opIns}
	y := 0
	for i := len(iStr) - 3; i >= 0; i-- {
		opCode, _ := strconv.Atoi(iStr[i : i+1])
		switch y {
		case 0:
			opResult.first = opCode
		case 1:
			opResult.second = opCode
		case 2:
			opResult.third = opCode
		}
		y++
	}
	return opResult
}

func getVal(parameterMode int, intData []int, index int) int {
	if parameterMode == 0 {
		return intData[intData[index]]
	}
	return intData[index]
}
