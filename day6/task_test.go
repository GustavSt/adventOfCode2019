package main

import "testing"

func TestTask1(t *testing.T) {
	res := Task1("testData.txt")
	expected := 42
	if res != expected {
		t.Errorf("Task1 was %d, expected %d", res, expected)
	}
}

func TestTask2(t *testing.T) {
	res := Task2("testData2.txt")
	expected := 4
	if res != expected {
		t.Errorf("Task2 was %d, expected %d", res, expected)
	}
}
