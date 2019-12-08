package main

import (
	"fmt"

	"../utils"
	"github.com/fatih/color"
)

//Layer a layer
type Layer struct {
	pixels [][]int
}

//Task1 do stuff
func Task1(dataFile string, width int, height int) int {
	data, _ := utils.GetData(8, dataFile)
	intData := utils.ConvertToIntArr(string(data))
	layers := getLayers(intData, width, height)
	fewest0 := utils.GetMaxInt()
	var fewest0Layer Layer
	for _, l := range layers {
		noPix := getNoPixels(0, l)
		if noPix < fewest0 {
			fewest0 = noPix
			fewest0Layer = l
		}
	}

	printLayer(fewest0Layer, 0)
	pixVal := getMultiValue(fewest0Layer, 1, 2)
	fmt.Printf("PixValue: %d\n", pixVal)
	fmt.Printf("number of layers: %d\n", len(layers))

	return pixVal
}

func printLayers(layers []Layer) {
	for i, l := range layers {
		printLayer(l, i)
	}
}
func printLayer(l Layer, layerNo int) {
	fmt.Printf("Layer %d\n", layerNo)
	for _, r := range l.pixels {
		for _, p := range r {
			if p == 0 {
				c := color.New(color.BgBlack)
				c.Print(" ")
			}
			if p == 1 {
				c := color.New(color.BgWhite)
				c.Print(" ")
			}
			if p == 2 {
				fmt.Printf(" ")
			}
		}
		fmt.Println("")
	}
}

func getNoPixels(pixel int, l Layer) int {
	count := 0
	for _, row := range l.pixels {
		for _, p := range row {
			if p == pixel {
				count++
			}
		}
	}
	return count
}

func getMultiValue(l Layer, p1 int, p2 int) int {
	val1 := getNoPixels(p1, l)
	val2 := getNoPixels(p2, l)

	res := val1 * val2
	return res
}

func getLayers(intData []int, width int, height int) []Layer {
	layers := []Layer{}
	currLayer := Layer{pixels: make([][]int, height)}
	currRow := 0
	for i, pixel := range intData {
		if i > 0 && i%width == 0 {
			currRow++
		}
		if currRow == height {
			layers = append(layers, currLayer)
			currLayer = Layer{pixels: make([][]int, height)}
			currRow = 0
		}
		if currLayer.pixels[currRow] == nil {
			currLayer.pixels[currRow] = []int{}
		}
		currLayer.pixels[currRow] = append(currLayer.pixels[currRow], pixel)
	}
	layers = append(layers, currLayer)
	return layers
}

//Task2 do stuff
func Task2(dataFile string, width int, height int) Layer {
	data, _ := utils.GetData(8, dataFile)
	intData := utils.ConvertToIntArr(string(data))
	layers := getLayers(intData, width, height)
	layer := getTransparentLayer(width, height)
	for _, l := range layers {
		for h := 0; h < height; h++ {
			for w := 0; w < width; w++ {
				if layer.pixels[h][w] == 2 {
					layer.pixels[h][w] = l.pixels[h][w]
				}
			}
		}
	}
	printLayer(layer, 0)
	return layer
}

func getTransparentLayer(width int, height int) Layer {
	layer := Layer{}
	layer.pixels = make([][]int, height)
	for i := 0; i < height; i++ {
		layer.pixels[i] = make([]int, width)
		for j := 0; j < width; j++ {
			layer.pixels[i][j] = 2
		}
	}
	return layer
}
