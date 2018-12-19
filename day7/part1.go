package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
	"adventofgo/helpers"
)

// input: ["X must come before Y", ...]

func extractInfo(lines []string) ([]string, map[string][]string, map[string]int) {
	lettersMap := make(map[string]bool)
	var firstLetters []string
	unlocks := make(map[string][]string)
	dependencies := make(map[string]int)

	var a, b string
	for _, line := range lines {
		words := strings.Split(line, " ")
		a, b = words[1], words[7]
		lettersMap[a] = true
		lettersMap[b] = true
		unlocks[a] = append(unlocks[a], b)
	}

	for _, v := range unlocks {
		for _, next := range v {
			dependencies[next] += 1
		}
	}

	for letter, _ := range lettersMap {
		if dependencies[letter] == 0 {
			firstLetters = append(firstLetters, letter)
		}
	}

	sort.Strings(firstLetters)

	return firstLetters, unlocks, dependencies
}

func toposort(first []string, unlocks map[string][]string, dependencies map[string]int) []string {

	var order []string

	// make a heap starting with all first letters
	// priority is simply the letter itself
	var pq = helpers.NewPQ(len(first))

	i := 0
	for _, letter := range first {
		char := []byte(letter)
		pq[i] = helpers.NewItem(letter, int(char[0]))
		i++
	}

	heap.Init(&pq)

	// keep popping the lowest priority letter, appending to list
	// add new letters to heap once they have no more prereqs
	for len(pq) > 0 {
		item := heap.Pop(&pq).(*helpers.Item)
		order = append(order, item.Value)
		nextLetters := unlocks[item.Value]

		toAdd := []string{}
		for _, next := range nextLetters {
			dependencies[next] -= 1
			if dependencies[next] == 0 {
				toAdd = append(toAdd, next)
				delete(dependencies, next)
			}
		}
		sort.Strings(toAdd)
		for _, newLetter := range toAdd {
			if !helpers.StringInSlice(newLetter, order) {
				char := []byte(newLetter)
				item := helpers.NewItem(newLetter, int(char[0]))
				heap.Push(&pq, item)
			}
		}
	}

	return order
}

func main() {
	var data []string = helpers.GetLines()

	/*
	   data := []string{"Step C must be finished before step A can begin.",
	       "Step C must be finished before step F can begin.",
	       "Step A must be finished before step B can begin.",
	       "Step A must be finished before step D can begin.",
	       "Step B must be finished before step E can begin.",
	       "Step D must be finished before step E can begin.",
	       "Step F must be finished before step E can begin."}
	*/

	first, unlocks, dependencies := extractInfo(data)

	result := toposort(first, unlocks, dependencies)

	fmt.Print("\n", strings.Join(result, ""))
}
