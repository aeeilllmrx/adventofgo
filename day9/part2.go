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

type Marble struct {
	value int
	prev  *Marble
	next  *Marble
}

func insertMarble(marble *Marble, curr int) *Marble {
	newMarble := &Marble{curr, marble.next, marble.next.next}
	marble.next.next.prev = newMarble
	marble.next.next = newMarble

	return marble.next.next
}

func removeMarble(marble *Marble) *Marble {
	marble.prev.next = marble.next
	marble.next.prev = marble.prev

	return marble
}

func placeMarbles(numPlayers int, lastMarble int) int {
	var currPlayer int
	numMarbles := 1
	marble := &Marble{0, nil, nil}
	marble.next = marble
	marble.prev = marble

	scores := make(map[int]int)

	for numMarbles < lastMarble {
		currPlayer = currPlayer % numPlayers

		if numMarbles%23 != 0 {
			marble = insertMarble(marble, numMarbles)
		} else {
			scores[currPlayer] += numMarbles
			marble = removeMarble(marble.prev.prev.prev.prev.prev.prev.prev)
			scores[currPlayer] += marble.value
			marble = marble.next
		}

		numMarbles += 1
		currPlayer += 1
	}

	maxVal := 0
	for _, v := range scores {
		if v > maxVal {
			maxVal = v
		}
	}

	return maxVal
}

func main() {
	// same as part one, except multiply lastMarble by 100
	// therefore we need to use a linked list-like data structure

	var data string = helpers.GetStringInput()

	//data = "9 players; last marble is worth 25 points" // 32
	//data = "10 players; last marble is worth 1618 points" // 8317

	numPlayers, lastMarble := getInts(data)
	lastMarble *= 100

	playerScores := placeMarbles(numPlayers, lastMarble)

	fmt.Print("\n", playerScores)

}
