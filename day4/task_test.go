package main

import "testing"

func TestIsValidPassword2(t *testing.T) {

	validNos := []int{122345, 112233, 111122, 112223}
	invalidNos := []int{135679, 123789, 123444, 111111, 111123}
	// validNos := []int{122345}
	for _, i := range validNos {
		res := IsValidPassword2(i)
		if !res {
			t.Errorf("Valid no %d is marked as invalid", i)
		}
	}
	for _, i := range invalidNos {
		res := IsValidPassword2(i)
		if res {
			t.Errorf("Valid no %d is marked as valid", i)
		}
	}
}

func TestIsValidPassword1(t *testing.T) {

	validNos := []int{111111, 111123, 122345}
	invalidNos := []int{135679, 123789}
	// validNos := []int{122345}
	for _, i := range validNos {
		res := IsValidPassword1(i)
		if !res {
			t.Errorf("Valid no %d is marked as invalid", i)
		}
	}
	for _, i := range invalidNos {
		res := IsValidPassword1(i)
		if res {
			t.Errorf("Valid no %d is marked as valid", i)
		}
	}
}

// func TestTask1(t *testing.T) {
// 	res := Task1("testData.txt")
// 	expected := 4
// 	if res != expected {
// 		t.Errorf("Task1 was %d, expected %d", res, expected)
// 	}
// }

// func TestTask2(t *testing.T) {
// 	res := Task2("testData.txt")
// 	expected := 4
// 	if res != expected {
// 		t.Errorf("Task2 was %d, expected %d", res, expected)
// 	}
// }
