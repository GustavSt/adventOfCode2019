package main

import (
	"testing"

	"../utils"
)

func TestAmplifierController(t *testing.T) {
	data1, _ := utils.SplitOnChar(7, "testData.txt", ",")
	intData1 := utils.ConvertToInt(data1)

	phaseSetting1 := []int{4, 3, 2, 1, 0}
	output1 := AmplifierController(0, phaseSetting1, intData1)
	expected1 := 43210
	if output1 != expected1 {
		t.Errorf("Expected output for phaseSetting %v \n was %d, but got %d", phaseSetting1, expected1, output1)
	}

	data2, _ := utils.SplitOnChar(7, "testData2.txt", ",")
	intData2 := utils.ConvertToInt(data2)

	phaseSetting2 := []int{0, 1, 2, 3, 4}
	output2 := AmplifierController(0, phaseSetting2, intData2)
	expected2 := 54321
	if output2 != expected2 {
		t.Errorf("Expected output for phaseSetting %v \n was %d, but got %d", phaseSetting2, expected2, output2)
	}

	data3, _ := utils.SplitOnChar(7, "testData3.txt", ",")
	intData3 := utils.ConvertToInt(data3)

	phaseSetting3 := []int{1, 0, 4, 3, 2}
	output3 := AmplifierController(0, phaseSetting3, intData3)
	expected3 := 65210
	if output3 != expected3 {
		t.Errorf("Expected output for phaseSetting %v \n was %d, but got %d", phaseSetting3, expected3, output3)
	}

}

func TestTask1(t *testing.T) {
	res := Task1("testData.txt")
	expected := 43210
	if res != expected {
		t.Errorf("Task1 was %d, expected %d", res, expected)
	}

	res = Task1("testData2.txt")
	expected = 54321
	if res != expected {
		t.Errorf("Task1 was %d, expected %d", res, expected)
	}

	res = Task1("testData3.txt")
	expected = 65210
	if res != expected {
		t.Errorf("Task1 was %d, expected %d", res, expected)
	}
}

func TestTask2(t *testing.T) {
	res := Task2("testData4.txt")
	expected := 139629729
	if res != expected {
		t.Errorf("Task2 was %d, expected %d", res, expected)
	}

	res = Task2("testData5.txt")
	expected = 18216
	if res != expected {
		t.Errorf("Task2 was %d, expected %d", res, expected)
	}
}
