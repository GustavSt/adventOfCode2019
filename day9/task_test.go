package main

import (
	"strconv"
	"testing"

	"../utils"
)

func TestTask1First(t *testing.T) {
	res := Task1("testData.txt")
	testData, _ := utils.SplitOnChar(9, "testData.txt", ",")
	expected := utils.ConvertToInt(testData)
	if len(expected) != len(res) {
		t.Errorf("Expected lendth of %v\n to be the same of %v", expected, res)
	}
	for i, r := range res {
		if r != expected[i] {
			t.Errorf("Task First was %d, expected %d", r, expected[i])
		}
	}
}

func TestTask1Second(t *testing.T) {
	res := Task1("testData2.txt")
	resStr := strconv.Itoa(res[0])
	expected := 16 //16 digit number
	if len(resStr) != expected {
		t.Errorf("Task Second was %d, expected %d", res[0], expected)
	}
}

func TestTask1Third(t *testing.T) {
	res := Task1("testData3.txt")
	expected := 1125899906842624
	if res[0] != expected {
		t.Errorf("Task third was %d, expected %d", res[0], expected)
	}
}
func TestTask2(t *testing.T) {
	res := Task2("testData.txt")
	expected := 4
	if res != expected {
		t.Errorf("Task2 was %d, expected %d", res, expected)
	}
}
