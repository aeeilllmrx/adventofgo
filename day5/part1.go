package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// input: a really long string
// repeatedly search for lowercase and uppercase of the same letter and remove them

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

func main() {
	data := getInput()

	canReact := true
	successfulRound := false

	for canReact == true {
		successfulRound = false
		for i := 0; i < len(data)-1; i++ {
			if react(string(data[i]), string(data[i+1])) {
				data = append(data[:i], data[i+2:]...)
				successfulRound = true
			}
		}
		if successfulRound != true || len(data) <= 1 {
			canReact = false
		}
	}

	fmt.Print(len(data))
}
