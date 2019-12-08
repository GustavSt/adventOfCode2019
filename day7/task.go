package main

import (
	"fmt"
	"strconv"

	"../utils"
)

//Task1 do stuff
func Task1(dataFile string) int {
	data, _ := utils.SplitOnChar(7, dataFile, ",")
	intData := utils.ConvertToInt(data)

	maxThruster := 0
	bestPhaseSettings := make([]int, 5)

	combinations := getCombinations([]int{0, 1, 2, 3, 4})
	for _, phaseSettings := range combinations {
		output := AmplifierController(0, phaseSettings, intData)
		if output > maxThruster {
			maxThruster = output
			bestPhaseSettings = phaseSettings
		}
	}

	fmt.Printf("Max thrust was %d, given by phaseSetting %v \n", maxThruster, bestPhaseSettings)
	return maxThruster
}
func getCombinations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

//Task2 do stuff
func Task2(dataFile string) int {
	data, _ := utils.SplitOnChar(7, dataFile, ",")
	intData := utils.ConvertToInt(data)

	maxThruster := 0
	bestPhaseSettings := make([]int, 5)
	combinations := getCombinations([]int{5, 6, 7, 8, 9})
	for _, phaseSettings := range combinations {
		output := AmplifierController(0, phaseSettings, intData)
		if output > maxThruster {
			maxThruster = output
			bestPhaseSettings = phaseSettings
		}
	}
	fmt.Printf("Max thrust was %d, given by phaseSetting %v \n", maxThruster, bestPhaseSettings)

	return maxThruster
}

//AmplifierController c
func AmplifierController(ampInput int, phaseSettings []int, intData []int) int {
	ampAInput := make(chan int)
	ampBInput := make(chan int)
	ampCInput := make(chan int)
	ampDInput := make(chan int)
	ampEInput := make(chan int)

	ampAOutput := make(chan int)
	ampBOutput := make(chan int)
	ampCOutput := make(chan int)
	ampDOutput := make(chan int)
	ampEOutput := make(chan int)

	ampADone := make(chan bool)

	go IntCodeProgram(ampAInput, intData, ampAOutput, ampADone) // A
	go IntCodeProgram(ampBInput, intData, ampBOutput, nil)      // B
	go IntCodeProgram(ampCInput, intData, ampCOutput, nil)      // C
	go IntCodeProgram(ampDInput, intData, ampDOutput, nil)      // D
	go IntCodeProgram(ampEInput, intData, ampEOutput, nil)      // E

	ampAInput <- phaseSettings[0]
	ampBInput <- phaseSettings[1]
	ampCInput <- phaseSettings[2]
	ampDInput <- phaseSettings[3]
	ampEInput <- phaseSettings[4]
	ampAInput <- 0
	go listenToChannelAndSend("B", ampBInput, ampAOutput)
	go listenToChannelAndSend("C", ampCInput, ampBOutput)
	go listenToChannelAndSend("D", ampDInput, ampCOutput)
	go listenToChannelAndSend("E", ampEInput, ampDOutput)

	latestThrust := 0
	for {
		_, aOpen := <-ampADone
		if !aOpen {
			break
		}
		ampEThrust := <-ampEOutput
		latestThrust = ampEThrust
		go sendInputToChannel(latestThrust, ampAInput)
	}
	return latestThrust
}

func listenToChannelAndSend(receivername string, receiverChan chan int, senderChan chan int) {
	for val := range senderChan {
		receiverChan <- val
	}
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
			intData[intData[index+1]] = <-inputChan
			index = index + 2
		case 4:
			output := getVal(opInst.first, intData, index+1)
			outputs = append(outputs, output)
			outputChan <- output
			if endOfProgram != nil {
				endOfProgram <- false
			}
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

	if endOfProgram != nil {
		close(endOfProgram)
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
