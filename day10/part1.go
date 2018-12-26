package main

import (
	"adventofgo/helpers"
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
	x, y int
}

type Velocity struct {
	x, y int
}

type Point struct {
	p Position
	v Velocity
}

func (point *Point) step() {
	point.p.x += point.v.x
	point.p.y += point.v.y
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func extract(data []string) []Point {
	var posData, velData []string
	var xp, yp, xv, yv int
	var points []Point

	for _, line := range data {
		posData = strings.Split(line[10:24], ", ")
		velData = strings.Split(line[36:42], ", ")

		xp, _ = strconv.Atoi(strings.TrimSpace(posData[0]))
		yp, _ = strconv.Atoi(strings.TrimSpace(posData[1]))

		xv, _ = strconv.Atoi(strings.TrimSpace(velData[0]))
		yv, _ = strconv.Atoi(strings.TrimSpace(velData[1]))

		points = append(points, Point{Position{xp, yp}, Velocity{xv, yv}})
	}

	return points
}

func takeStep(points []Point) {
	// why can't i update a slice of structs by reference?
	var newPoints = points[:0]

	for _, point := range points {
		point.step()
		newPoints = append(newPoints, point)
	}
}

func closeEnough(points []Point) bool {
	y_max := -10000
	y_min := 10000

	for _, point := range points {
		y_max = max(y_max, point.p.y)
		y_min = min(y_min, point.p.y)
	}

	if y_max-y_min < 10 {
		return true
	} else {
		return false
	}
}

func printGrid(points []Point) {
	var pmap = make(map[Position]bool)
	var x_min, y_min, x_max, y_max = 1000, 1000, -1000, -1000
	var byteSlice []byte
	var line string

	for _, point := range points {
		x_min = min(x_min, point.p.x)
		x_max = max(x_max, point.p.x)
		y_min = min(y_min, point.p.y)
		y_max = max(y_max, point.p.y)
		pmap[Position{point.p.x, point.p.y}] = true
	}

	for i := x_min; i <= x_max; i++ {
		byteSlice = []byte{}
		for j := y_min; j <= y_max; j++ {
			if pmap[Position{i, j}] {
				byteSlice = append(byteSlice, 'x')
			} else {
				byteSlice = append(byteSlice, '.')
			}
		}
		line = string(byteSlice[:])
		fmt.Println(line)
	}
}

func main() {
	var data []string = helpers.GetLines()
	//position=< 10253, -50152> velocity=<-1,  5>

	points := extract(data)

	for {
		takeStep(points)

		if closeEnough(points) {
			printGrid(points)
			break
		}
	}
}
