package main

import (
	"bufio"
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

func getMin(letters map[string]bool) string {
	minLetter := "Z"
	minChar := int([]byte(minLetter)[0])

	for letter, _ := range letters {
		char := int([]byte(letter)[0])
		if char < minChar {
			minChar = char
			minLetter = letter
		}
	}
	return minLetter
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

func toposort(first []string, unlocks map[string][]string, dependencies map[string]int) int {

	// maintain at most 5 keys whose values are the expiration time
	// add the alpha-earliest available letter whenever a letter finishes

	numWorkers := 5
	alphabetSize := 26
	taskTime := 60
	currentTime := 0

	available := make(map[string]bool)
	inProgress := make(map[string]int)
	completed := make(map[string]bool)

	for _, letter := range first {
		available[letter] = true
		char := int([]byte(letter)[0])
		inProgress[letter] = taskTime + char - 64
	}

	for len(completed) < alphabetSize {
		for len(inProgress) < numWorkers {
			if len(available) == 0 {
				break
			}
			letter := getMin(available)
			delete(available, letter)
			char := []byte(letter)
			inProgress[letter] = currentTime - 1 + taskTime + int(char[0]-64)
		}

		justCompleted := []string{}
		for letter, doneTime := range inProgress {
			if doneTime <= currentTime {
				if !completed[letter] {
					justCompleted = append(justCompleted, letter)
				}
			}
		}

		for _, letter := range justCompleted {
			delete(inProgress, letter)
			completed[letter] = true

			nextLetters := unlocks[letter]
			for _, next := range nextLetters {
				if dependencies[next] > 0 {
					dependencies[next] -= 1
					if dependencies[next] == 0 {
						available[next] = true
						delete(dependencies, next)
					}
				}
			}
		}
		currentTime += 1
	}

	return currentTime
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

	fmt.Print("\n", result)
}
