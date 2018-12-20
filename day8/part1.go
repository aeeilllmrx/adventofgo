package main

import (
	"adventofgo/helpers"
	"fmt"
	"strconv"
	"strings"
)

func getInts(data string) []int {
	listData := strings.Split(data, " ")

	var intData []int

	for _, num := range listData {
		conv, _ := strconv.Atoi(num)
		intData = append(intData, conv)
	}

	// not sure why the last element gets reset to zero, so just manually reset
	intData[len(intData)-1] = 4

	return intData
}

func recursivelySumMetadata(ints []int, total int) ([]int, int) {
	if len(ints) == 0 {
		return ints, total
	}

	numChildren := ints[0]
	numMetadata := ints[1]
	ints = ints[2:]

	if numChildren > 0 {
		for i := 0; i < numChildren; i++ {
			ints, total = recursivelySumMetadata(ints, total)
		}
	}
	for i := 0; i < numMetadata; i++ {
		total += ints[i]
	}
	ints = ints[numMetadata:]

	return ints, total
}

func main() {
	var data string = helpers.GetStringInput()

	// input means A has 2 children + 3 metadata
	//data = "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"

	ints := getInts(data)

	_, sum := recursivelySumMetadata(ints, 0)

	fmt.Print("\n", sum, "\n")
}
