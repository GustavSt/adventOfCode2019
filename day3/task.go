package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"../utils"
)

//Task1 do stuff
func Task1(dataFile string) int {
	data, _ := utils.SplitData(3, dataFile)
	wire1 := strings.Split(data[0], ",")
	wire2 := strings.Split(data[1], ",")

	wire1Pos := getVectorsFromWire(wire1)
	wire2Pos := getVectorsFromWire(wire2)

	// fmt.Printf("wire1: %v \n", wire1Pos)
	// fmt.Printf("wire2: %v \n", wire2Pos)
	crossingVectors := getCrossingVectors(wire1Pos, wire2Pos)
	// fmt.Printf("%v \n", crossingVectors)
	origo := utils.Vector2{X: 0, Y: 0}
	closestVector := getClosestVector(origo, crossingVectors)
	distance := manhattanDistance(origo, closestVector)
	fmt.Printf("distance: %d \n", distance)
	fmt.Printf("%+v\n", closestVector)
	return distance
}

//Task2 do stuff
func Task2(dataFile string) int {
	data, _ := utils.SplitData(3, dataFile)
	wire1 := strings.Split(data[0], ",")
	wire2 := strings.Split(data[1], ",")

	wire1Pos := getVectorsFromWire(wire1)
	wire2Pos := getVectorsFromWire(wire2)

	crossingVectors := getCrossingVectors(wire1Pos, wire2Pos)

	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	val := MaxInt
	for _, v := range crossingVectors {
		w1 := numOfSteps(v, wire1Pos)
		w2 := numOfSteps(v, wire2Pos)
		if w1+w2 < val {
			val = w1 + w2
		}
	}
	return val
}

func numOfSteps(dest utils.Vector2, wire []utils.Vector2) int {
	var prevV utils.Vector2
	steps := 0
	for i, v := range wire {
		if i == 0 {
			prevV = v
			continue
		}
		if (prevV.X < dest.X && dest.X < v.X && prevV.Y == dest.Y && dest.Y == v.Y) ||
			(prevV.X > dest.X && dest.X > v.X && prevV.Y == dest.Y && dest.Y == v.Y) ||
			(prevV.Y < dest.Y && dest.Y < v.Y && prevV.X == dest.X && dest.X == v.X) ||
			(prevV.Y > dest.Y && dest.Y > v.Y && prevV.X == dest.X && dest.X == v.X) {
			steps = steps + manhattanDistance(prevV, dest)
			return steps
		}
		steps = steps + manhattanDistance(prevV, v)
		prevV = v
	}
	return steps
}

func getCrossingVectors(wire1 []utils.Vector2, wire2 []utils.Vector2) []utils.Vector2 {
	crossingVectors := []utils.Vector2{}

	var prevV1 utils.Vector2
	for i, v1 := range wire1 {
		if i == 0 {
			prevV1 = v1
			continue
		}
		var prevV2 utils.Vector2
		for j, v2 := range wire2 {
			if j == 0 {
				prevV2 = v2
				continue
			}
			cross, err := wireSection(prevV1, v1, prevV2, v2)
			if err == nil {
				crossingVectors = append(crossingVectors, cross)
			}
			prevV2 = v2
		}
		prevV1 = v1
	}
	return crossingVectors
}

func wireSection(wire1a utils.Vector2, wire1b utils.Vector2, wire2a utils.Vector2, wire2b utils.Vector2) (utils.Vector2, error) {
	if wire1a.X == wire1b.X && wire2a.X != wire2b.X {
		if wire2a.Y >= wire1a.Y && wire2a.Y <= wire1b.Y || wire2a.Y <= wire1a.Y && wire2a.Y >= wire1b.Y {
			if wire2a.X >= wire1a.X && wire2b.X <= wire1a.X || wire2b.X >= wire1a.X && wire2a.X <= wire1a.X {
				//Intersect
				v := utils.Vector2{X: wire1a.X, Y: wire2a.Y}
				if v.X == 0 && v.Y == 0 {
					return utils.Vector2{}, errors.New("origo intersect")
				}
				return v, nil
			}
		}
	}

	if wire1a.Y == wire1b.Y && wire2a.Y != wire2b.Y {
		if wire2a.X >= wire1a.X && wire2a.X <= wire1b.X || wire2a.X <= wire1a.X && wire2a.X >= wire1b.X {
			if wire2a.Y >= wire1a.Y && wire2b.Y <= wire1a.Y || wire2b.Y >= wire1a.Y && wire2a.Y <= wire1a.Y {
				//Intersect
				v := utils.Vector2{X: wire2a.X, Y: wire1a.Y}
				if v.X == 0 && v.Y == 0 {
					return utils.Vector2{}, errors.New("origo intersect")
				}
				return v, nil
			}
		}
	}

	// s1 := utils.Vector2{X: wire1b.X - wire1a.X, Y: wire1b.Y - wire1a.Y}
	// s2 := utils.Vector2{X: wire2b.X - wire2b.X, Y: wire2b.Y - wire2a.Y}
	// var s, t float32
	// d := float32((-s2.X*s1.Y + s1.X*s2.Y))
	// if d == 0 {
	// 	return utils.Vector2{}, errors.New("no intersect1")
	// }
	// s = float32((-s1.Y*(wire1a.X-wire2a.X) + s1.X*(wire1a.Y-wire2a.Y))) / d
	// t = float32((s2.X*(wire1a.Y-wire2a.Y) - s2.Y*(wire1a.X-wire2a.X))) / d

	// if s >= 0 && s <= 1 && t >= 0 && t <= 1 {
	// 	v := utils.Vector2{X: wire1a.X + (t * s1.X)}
	// }
	return utils.Vector2{}, errors.New("no intersect1")
}

func getClosestVector(origo utils.Vector2, vectors []utils.Vector2) utils.Vector2 {
	// var closesVector utils.Vector2
	var closestVector utils.Vector2
	// closestDistance := ^uint(0)
	for i, v := range vectors {
		if i == 0 {
			closestVector = v
			continue
		}
		closestDistance := manhattanDistance(origo, closestVector)
		currDistance := manhattanDistance(origo, v)

		if currDistance < closestDistance {
			// closestDistance = currDistance
			closestVector = v
		}
	}
	return closestVector
}

func manhattanDistance(v1 utils.Vector2, v2 utils.Vector2) int {
	return utils.Abs(v1.X-v2.X) + utils.Abs(v1.Y-v2.Y)
}

func getVectorsFromWire(wire []string) []utils.Vector2 {
	currVector := utils.Vector2{X: 0, Y: 0}
	positions := []utils.Vector2{currVector}

	for _, pos := range wire {
		dir := pos[:1]
		nextVecotor := utils.Vector2{X: currVector.X, Y: currVector.Y}
		dirValue, err := strconv.Atoi(pos[1:])
		if err != nil {
			panic(err)
		}
		switch dir {
		case "R":
			nextVecotor.X = nextVecotor.X + dirValue
		case "L":
			nextVecotor.X = nextVecotor.X - dirValue
		case "U":
			nextVecotor.Y = nextVecotor.Y + dirValue
		case "D":
			nextVecotor.Y = nextVecotor.Y - dirValue
		}
		positions = append(positions, nextVecotor)
		currVector = nextVecotor
	}
	return positions
}
