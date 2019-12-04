package main

import (
	"strconv"

	"../utils"
)

//Task1 do stuff
func Task1(dataFile string) int {
	data, _ := utils.SplitOnChar(4, dataFile, "-")
	startRange, _ := strconv.Atoi(data[0])
	endRange, _ := strconv.Atoi(data[1])

	noPasswords := 0
	for i := startRange; i <= endRange; i++ {
		if IsValidPassword1(i) {
			noPasswords++
		}
	}
	return noPasswords
}

//IsValidPassword1 for test
func IsValidPassword1(val int) bool {
	str := strconv.Itoa(val)
	var prevRune rune
	hasDouble := false
	// fmt.Printf("string number: %s \n", str)
	for i, r := range str {
		if i == 0 {
			prevRune = r
			continue
		}
		prevDigit := utils.RuneToInt(prevRune)
		currDigit := utils.RuneToInt(r)
		if currDigit < prevDigit {
			// fmt.Printf("invalid du to %d is lower than previous digit %d\n", currDigit, prevDigit)
			return false
		}
		// fmt.Printf("Comparing prevDigit %d, and currDigit %d \n", prevDigit, currDigit)
		if utils.RuneToInt(prevRune) == utils.RuneToInt(r) {
			// fmt.Printf("digits match \n")
			hasDouble = true
		}
		prevRune = r
	}

	return hasDouble
}

//Task2 do stuff
func Task2(dataFile string) int {
	data, _ := utils.SplitOnChar(4, dataFile, "-")
	startRange, _ := strconv.Atoi(data[0])
	endRange, _ := strconv.Atoi(data[1])

	noPasswords := 0
	for i := startRange; i <= endRange; i++ {
		if IsValidPassword2(i) {
			noPasswords++
		}
	}
	return noPasswords
}

//IsValidPassword2 for test
func IsValidPassword2(val int) bool {
	str := strconv.Itoa(val)
	var prevRune rune
	var prevPrevRune rune
	hasDouble := false
	pair := false
	for i, r := range str {
		if i == 0 {
			prevRune = r
			continue
		}

		prevDigit := utils.RuneToInt(prevRune)
		currDigit := utils.RuneToInt(r)
		if currDigit < prevDigit {
			return false
		}
		if prevDigit == currDigit {
			pair = true
			if prevPrevRune != 0 {
				prevPrevDigit := utils.RuneToInt(prevPrevRune)
				if prevPrevDigit == currDigit {
					pair = false
				}
			}
		} else if pair {
			hasDouble = true
		} else {
			pair = false
		}
		prevPrevRune = prevRune
		prevRune = r
	}

	return hasDouble || pair
}
