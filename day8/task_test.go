package main

import "testing"

func TestTask1(t *testing.T) {
	res := Task1("testData.txt", 3, 2)
	expected := 1
	if res != expected {
		t.Errorf("Task1 was %d, expected %d", res, expected)
	}
}

func TestTask2(t *testing.T) {
	res := Task2("testData2.txt", 2, 2)
	row1 := []int{0, 1}
	row2 := []int{1, 0}
	for i, r := range res.pixels {
		for j, p := range r {
			if i == 0 {
				if row1[j] != p {
					t.Errorf("Task2 expected pixel %d but got pixel %d\n", row1[j], p)
				}
			} else {
				if row2[j] != p {
					t.Errorf("Task2 expected pixel %d but got pixel %d\n", row1[j], p)
				}
			}
		}
	}
}
