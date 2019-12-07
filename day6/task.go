package main

import (
	"strings"

	"../utils"
)

type spaceObject struct {
	name   string
	orbits *spaceObject
}

//Task1 do stuff
func Task1(dataFile string) int {
	data, _ := utils.SplitData(6, dataFile)
	orbitMap := getOrbitMap(data)
	noOrbits := 0
	for _, orbit := range orbitMap {
		noOrbits = noOrbits + getNoOrbits(orbit)
	}

	return noOrbits
}
func getNoOrbits(spaceObj *spaceObject) int {
	noOrbits := 0
	currOrbit := spaceObj.orbits
	for currOrbit != nil {
		noOrbits++
		currOrbit = currOrbit.orbits
	}
	return noOrbits
}
func getOrbitMap(data []string) map[string]*spaceObject {
	orbitMap := make(map[string]*spaceObject)
	for _, o := range data {
		or := strings.Split(o, ")")
		o1, o1Ok := orbitMap[or[0]]
		if !o1Ok {
			o1 = &spaceObject{name: or[0]}
			orbitMap[o1.name] = o1
		}
		o2, o2Ok := orbitMap[or[1]]
		if !o2Ok {
			o2 = &spaceObject{name: or[1], orbits: o1}
			orbitMap[o2.name] = o2
		} else {
			o2.orbits = o1
		}
	}
	return orbitMap
}

//Task2 do stuff
func Task2(dataFile string) int {
	data, _ := utils.SplitData(6, dataFile)
	orbitMap := getOrbitMap(data)

	start := orbitMap["YOU"].orbits
	dest := orbitMap["SAN"].orbits
	distance := getDistance(start, dest)
	return distance
}

func getDistance(so1 *spaceObject, so2 *spaceObject) int {
	dest := getClosestSame(so1, so2)

	distance := 0
	so1Curr := so1
	for so1Curr != dest {
		distance++
		so1Curr = so1Curr.orbits
	}
	so2Curr := so2
	for so2Curr != dest {
		distance++
		so2Curr = so2Curr.orbits
	}
	return distance
}

func getClosestSame(so1 *spaceObject, so2 *spaceObject) *spaceObject {
	so1Map := make(map[string]*spaceObject)
	curr1 := so1
	for curr1 != nil {
		so1Map[curr1.name] = curr1
		curr1 = curr1.orbits
	}
	curr2 := so2
	for curr2 != nil {
		_, existIn1 := so1Map[curr2.name]
		if existIn1 {
			return curr2
		}
		curr2 = curr2.orbits
	}
	return nil
}
