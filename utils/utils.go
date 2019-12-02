package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"runtime"
	"strconv"
	"strings"
)

//GetData ads
func GetData(day int, dataFile string) ([]byte, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("failed with caller")
	}
	filepath := path.Join(path.Dir(filename), fmt.Sprintf("../day%d/%s", day, dataFile))
	f, errRead := ioutil.ReadFile(filepath)
	if errRead != nil {
		log.Fatal(errRead)
		return nil, errRead
	}
	return f, nil
}

//SplitData get data split on \n\r
func SplitData(day int, dataFile string) ([]string, error) {
	data, err := GetData(day, dataFile)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	strs := strings.Split(string(data), "\r\n")
	return strs, nil
}

//SplitOnChar splits input data on specified rune
func SplitOnChar(day int, dataFile string, split string) ([]string, error) {
	data, err := GetData(day, dataFile)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	strs := strings.Split(string(data), split)
	return strs, nil
}

//ConvertToInt converts slice of strings to slice of ints
func ConvertToInt(arr []string) []int {
	intArr := []int{}
	for _, i := range arr {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intArr = append(intArr, j)
	}
	return intArr
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
