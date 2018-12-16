package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
Input is a list of coordinates.
Find the bounding box of all relevant cells.
Then find the number of cells whose summed distance to all coordinates < 10000.
*/

type Coord struct {
	x, y int
}

func getLines() []Coord {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var coords []Coord

	var x, y int
	for _, line := range lines {
		coord := strings.Split(line, ", ")
		x, _ = strconv.Atoi(coord[0])
		y, _ = strconv.Atoi(coord[1])
		coords = append(coords, Coord{x, y})
	}

	return coords
}

func getBoundaries(data []Coord) []int {
	var x_min, x_max, y_min, y_max int

	x_min = 1000
	x_max = 0
	y_min = 1000
	y_max = 0

	var x, y int
	for _, coord := range data {
		x = coord.x
		y = coord.y
		if x < x_min {
			x_min = x
		} else if x > x_max {
			x_max = x
		}
		if y < y_min {
			y_min = y
		} else if y > y_max {
			y_max = y
		}
	}
	return []int{x_min, x_max, y_min, y_max}
}

func getDistance(a Coord, b Coord) int {
	x := math.Abs(float64(a.x - b.x))
	y := math.Abs(float64(a.y - b.y))
	return int(x + y)
}

func getNumCells(boundaries []int, data []Coord) int {
	x_min := boundaries[0]
	x_max := boundaries[1]
	y_min := boundaries[2]
	y_max := boundaries[3]

	regionDist := 10000

	var numCells int
	var currDist int
	var thisCoord Coord

	for i := x_min; i <= x_max; i++ {
		for j := y_min; j <= y_max; j++ {
			currDist = 0
			thisCoord = Coord{i, j}
			for _, thatCoord := range data {
				currDist += getDistance(thisCoord, thatCoord)
			}
			if currDist < regionDist {
				numCells += 1
			}
		}
	}
	return numCells
}

func main() {
	var data []Coord = getLines()

	//data := []Coord{{1,1},{1,6},{8,3},{3,4},{5,5},{8,9}}

	var boundaries []int = getBoundaries(data)

	result := getNumCells(boundaries, data)

	fmt.Print("\n", result)
}
