package main

import "testing"

func TestTask2(t *testing.T) {
	res := Task2("testData2.txt", 7)
	expected := 999
	if res[0] != expected {
		t.Errorf("Task2 was %d, expected %d", res, expected)
	}
	res = Task2("testData2.txt", 8)
	expected = 1000
	if res[0] != expected {
		t.Errorf("Task2 was %d, expected %d", res, expected)
	}
	res = Task2("testData2.txt", 10)
	expected = 1001
	if res[0] != expected {
		t.Errorf("Task2 was %d, expected %d", res, expected)
	}
}

func TestGetOpInstructions(t *testing.T) {
	res := GetOpInstructions(1002)

	if res.opCode != 2 {
		t.Errorf("Expected %d to be 2", res.opCode)
	}
	if res.first != 0 {
		t.Errorf("Expected %d to be 0", res.first)
	}
	if res.second != 1 {
		t.Errorf("Expected %d to be 1", res.second)
	}
}
