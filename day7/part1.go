package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strings"
)

// input: ["X must come before Y", ...]

// helper for `if a in list`
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getLines() []string {
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

	return lines
}

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

// have to copy all this crap just to use a priority queue

type Item struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func toposort(first []string, unlocks map[string][]string, dependencies map[string]int) []string {

	var order []string

	// make a heap starting with all first letters
	// priority is simply the letter itself
	pq := make(PriorityQueue, len(first))

	i := 0
	for _, letter := range first {
		char := []byte(letter)
		pq[i] = &Item{
			value:    letter,
			priority: int(char[0]),
		}
		i++
	}

	heap.Init(&pq)

	// keep popping the lowest priority letter, appending to list
	// add new letters to heap once they have no more prereqs
	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
		order = append(order, item.value)
		nextLetters := unlocks[item.value]

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
			if !stringInSlice(newLetter, order) {
				char := []byte(newLetter)
				item := &Item{
					value:    newLetter,
					priority: int(char[0]),
				}
				heap.Push(&pq, item)
			}
		}
	}

	return order
}

func main() {
	var data []string = getLines()

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
