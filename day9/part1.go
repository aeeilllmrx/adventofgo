package main

import (
	"adventofgo/helpers"
	"fmt"
	"regexp"
	"strconv"
)

func getInts(data string) (int, int) {
	re := regexp.MustCompile("[0-9]+")
	match := re.FindAllString(data, -1)

	numPlayers, _ := strconv.Atoi(match[0])
	lastMarble, _ := strconv.Atoi(match[1])

	return numPlayers, lastMarble
}

func insertMarble(marbles *[]int, prev *int, curr int) {
	prevIndex := helpers.IntIndexSlice(*prev, *marbles)
	insertionIndex := (prevIndex + 2) % len(*marbles)

	start := append([]int{}, (*marbles)[:insertionIndex]...)
	end := append([]int{}, (*marbles)[insertionIndex:]...)
	*marbles = append(start, curr)
	*marbles = append(*marbles, end...)

	*prev = curr
}

func removeMarble(marbles *[]int, scores *map[int]int, prev *int, curr int) {
	prevIndex := helpers.IntIndexSlice(*prev, *marbles)

	// modulo operator doesn't work on negative numbers, so hack around
	N := len(*marbles)
	deletionIndex := ((prevIndex-7)%N + N) % N

	acc := (*marbles)[deletionIndex]

	start := append([]int{}, (*marbles)[:deletionIndex]...)
	end := append([]int{}, (*marbles)[deletionIndex+1:]...)
	*marbles = append(start, end...)

	(*scores)[curr] += acc
	*prev = (*marbles)[deletionIndex]
}

func placeMarbles(numPlayers int, lastMarble int) int {
	var prevMarble, currMarble int
	var currPlayer int
	var marbles []int

	playerScores := make(map[int]int)

	for currMarble <= lastMarble {
		currPlayer = currPlayer % numPlayers

		if len(marbles) <= 1 {
			marbles = append(marbles, currMarble)
			prevMarble = currMarble
		} else if currMarble%23 != 0 {
			insertMarble(&marbles, &prevMarble, currMarble)
		} else {
			playerScores[currPlayer] += currMarble
			removeMarble(&marbles, &playerScores, &prevMarble, currPlayer)
		}

		currMarble += 1
		currPlayer += 1
	}

	maxVal := 0
	for _, v := range playerScores {
		if v > maxVal {
			maxVal = v
		}
	}

	return maxVal
}

func main() {
	var data string = helpers.GetStringInput()

	//data = "9 players; last marble is worth 25 points" // 32
	//data = "10 players; last marble is worth 1618 points" // 8317

	numPlayers, lastMarble := getInts(data)

	playerScores := placeMarbles(numPlayers, lastMarble)

	fmt.Print("\n", playerScores)

}
