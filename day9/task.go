package main

import (
	"strconv"

	"../utils"
)

//Task1 do stuff
func Task1(dataFile string) []int {
	data, _ := utils.SplitOnChar(9, dataFile, ",")
	intData := utils.ConvertToInt(data)
	inputChan := make(chan int)
	outputChan := make(chan int)
	go IntCodeProgram(inputChan, intData, outputChan, nil)

	go sendInputToChannel(1, inputChan)
	output := []int{}
	for o := range outputChan {
		output = append(output, o)
	}
	return output
}

//Task2 do stuff
func Task2(dataFile string) []int {
	data, _ := utils.SplitOnChar(9, dataFile, ",")
	intData := utils.ConvertToInt(data)
	inputChan := make(chan int)
	outputChan := make(chan int)
	go IntCodeProgram(inputChan, intData, outputChan, nil)

	go sendInputToChannel(2, inputChan)
	output := []int{}
	for o := range outputChan {
		output = append(output, o)
	}
	return output
}
func sendInputToChannel(input int, c chan int) {
	c <- input
}

//IntCodeProgram i
func IntCodeProgram(inputChan chan int, intDataSet []int, outputChan chan int, endOfProgram chan bool) []int {
	intData := make([]int, len(intDataSet))
	copy(intData, intDataSet)
	index := 0
	opInst := GetOpInstructions(intData[index])
	outputs := []int{}
	relativeBase := 0
	for opInst.opCode != 99 {
		switch opInst.opCode {
		case opCodeAdd:
			first := getVal(opInst.first, intData, index+1, relativeBase)
			second := getVal(opInst.second, intData, index+2, relativeBase)
			write(&intData, intData[index+3], first+second, opInst.third, relativeBase)
			index = index + 4
		case opCodeMultiply:
			first := getVal(opInst.first, intData, index+1, relativeBase)
			second := getVal(opInst.second, intData, index+2, relativeBase)
			write(&intData, intData[index+3], first*second, opInst.third, relativeBase)
			index = index + 4
		case opCodeWriteFromInput:
			write(&intData, intData[index+1], <-inputChan, opInst.first, relativeBase)
			index = index + 2
		case opCodeOutput:
			output := getVal(opInst.first, intData, index+1, relativeBase)
			outputs = append(outputs, output)
			outputChan <- output
			if endOfProgram != nil {
				endOfProgram <- false
			}
			index = index + 2
		case opCodeJumpIfFalse:
			val := getVal(opInst.first, intData, index+1, relativeBase)
			if val != 0 {
				index = getVal(opInst.second, intData, index+2, relativeBase)
			} else {
				index = index + 3
			}
		case opCodeJumpIfTrue:
			val := getVal(opInst.first, intData, index+1, relativeBase)
			if val == 0 {
				index = getVal(opInst.second, intData, index+2, relativeBase)
			} else {
				index = index + 3
			}
		case opCodeLessThan:
			val1 := getVal(opInst.first, intData, index+1, relativeBase)
			val2 := getVal(opInst.second, intData, index+2, relativeBase)
			if val1 < val2 {
				write(&intData, intData[index+3], 1, opInst.third, relativeBase)
			} else {
				write(&intData, intData[index+3], 0, opInst.third, relativeBase)
			}
			index = index + 4
		case opcodeEquals:
			val1 := getVal(opInst.first, intData, index+1, relativeBase)
			val2 := getVal(opInst.second, intData, index+2, relativeBase)
			if val1 == val2 {
				write(&intData, intData[index+3], 1, opInst.third, relativeBase)
			} else {
				write(&intData, intData[index+3], 0, opInst.third, relativeBase)
			}
			index = index + 4
		case opCodeAdjustRelative:
			relativeBase = relativeBase + getVal(opInst.first, intData, index+1, relativeBase)
			index = index + 2
		}
		opInst = GetOpInstructions(intData[index])
	}

	if endOfProgram != nil {
		close(endOfProgram)
	}
	close(outputChan)
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

func getVal(parameterMode int, intData []int, index int, relativeBase int) int {
	if parameterMode == positionMode {
		return read(&intData, intData[index])
	}
	if parameterMode == immediateMode {
		return read(&intData, index)
	}
	return read(&intData, intData[index]+relativeBase)
}

func read(intData *[]int, index int) int {
	if len(*intData) <= index {
		sizeIncrease := index - len(*intData) + 1
		increasSlice(intData, sizeIncrease)
	}
	return (*intData)[index]
}

func write(intData *[]int, index int, val int, paramMode int, relativeBase int) {
	i := index
	if paramMode == relativeMode {
		i = i + relativeBase
	}
	if len(*intData) <= i {
		sizeIncrease := i - len(*intData) + 1
		increasSlice(intData, sizeIncrease)
	}
	(*intData)[i] = val
}

func increasSlice(intData *[]int, sizeIncrease int) {
	extension := make([]int, sizeIncrease)
	*intData = append(*intData, extension...)
}

const (
	positionMode  = 0
	immediateMode = 1
	relativeMode  = 2
)

const (
	opCodeAdd            = 1
	opCodeMultiply       = 2
	opCodeWriteFromInput = 3
	opCodeOutput         = 4
	opCodeJumpIfFalse    = 5
	opCodeJumpIfTrue     = 6
	opCodeLessThan       = 7
	opcodeEquals         = 8
	opCodeAdjustRelative = 9
)
