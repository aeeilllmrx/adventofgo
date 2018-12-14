package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// input: a really long string
// for each letter of the alphabet
// remove all instances and then check for result of reactions

func getInput() []byte {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Print(err)
	}

	return data
}

func react(i string, j string) bool {
	if strings.ToLower(i) == j && strings.ToUpper(j) == i {
		return true
	} else if strings.ToLower(j) == i && strings.ToUpper(i) == j {
		return true
	} else {
		return false
	}
}

func reactionResultLength(polymer []byte) int {
	canReact := true
	successfulRound := false

	for canReact == true {
		successfulRound = false
		for i := 0; i < len(polymer) - 1; i++ {
	        if react(string(polymer[i]), string(polymer[i + 1])) {
	        	polymer = append(polymer[:i], polymer[i + 2:]...)
	        	successfulRound = true
			}
	    }	
		if successfulRound != true || len(polymer) <= 1 {
			canReact = false
		}
	}

	return len(polymer)
}

func removeLetter(polymer []byte, letter byte) []byte {
	var newPolymer []byte

	for i := 0; i < len(polymer) - 1; i++ {
		if polymer[i] != letter && string(polymer[i]) != strings.ToUpper(string(letter)) {
			newPolymer = append(newPolymer, polymer[i])
		}
	}

	return []byte(newPolymer)
}

func main() {
	polymer := getInput()

	min_count := 10000
	newPolymer := []byte{}
	count := 0
	alphabet := []byte("abcdefghijklmnopqrstuvwxyz")

	for _, letter := range alphabet {
		newPolymer = removeLetter(polymer, letter)
		count = reactionResultLength(newPolymer)

		if count < min_count {
			min_count = count
		}
	}
	
	print(min_count)
}