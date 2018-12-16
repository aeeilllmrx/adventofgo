package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// input: a list of coordinates x, y
// find bounding box for all
// assign each point to closest square
// find coordinate with the most squares
// but disqualify any coordinate with a square on the border

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

func assignCoordinates(boundaries []int, data []Coord) map[Coord]Coord {
	x_min := boundaries[0]
	x_max := boundaries[1]
	y_min := boundaries[2]
	y_max := boundaries[3]

	matrix := make(map[Coord]Coord)

	var dist, currDist int
	var thisCoord, closestCoord Coord

	var lastTied bool
	zeroCoord := Coord{0, 0} // assign ties to this

	for i := x_min; i <= x_max; i++ {
		for j := y_min; j <= y_max; j++ {
			dist = 1000
			thisCoord = Coord{i, j}
			for _, thatCoord := range data {
				currDist = getDistance(thisCoord, thatCoord)
				if currDist < dist {
					dist = currDist
					closestCoord = thatCoord
					lastTied = false
				} else if currDist == dist {
					lastTied = true
				}
			}
			if lastTied == true {
				matrix[thisCoord] = zeroCoord
			} else {
				matrix[thisCoord] = closestCoord
			}
		}
	}
	return matrix
}

func getValidCoordinates(boundaries []int, matrix map[Coord]Coord, data []Coord) map[Coord]bool {
	x_min := boundaries[0]
	x_max := boundaries[1]
	y_min := boundaries[2]
	y_max := boundaries[3]

	invalidCoordMap := make(map[Coord]bool)

	for k, v := range matrix {
		if k.x == x_min || k.x == x_max || k.y == y_min || k.y == y_max {
			invalidCoordMap[v] = true
		}
	}

	validCoordMap := make(map[Coord]bool)
	for _, coord := range data {
		if invalidCoordMap[coord] != true {
			validCoordMap[coord] = true
		}
	}

	return validCoordMap
}

func findLargestValidCoordinate(matrix map[Coord]Coord, validCoords map[Coord]bool) int {
	coordCounts := make(map[Coord]int)

	for _, closestCoord := range matrix {
		if validCoords[closestCoord] == true {
			coordCounts[closestCoord] += 1
		}
	}

	var maxValue int

	for _, v := range coordCounts {
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}

func main() {
	var data []Coord = getLines()

	//data := []Coord{{1,1},{1,6},{8,3},{3,4},{5,5},{8,9}}

	var boundaries []int = getBoundaries(data)

	var matrix map[Coord]Coord = assignCoordinates(boundaries, data)

	var validCoords map[Coord]bool = getValidCoordinates(boundaries, matrix, data)

	result := findLargestValidCoordinate(matrix, validCoords)

	fmt.Print("\n", result)
}
