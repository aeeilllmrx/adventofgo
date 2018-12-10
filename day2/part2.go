package main

import (
	"bufio"
	"fmt"
	"os"
)

func compare(x string, y string) bool {
	n := len(x)

	diffs := 0
	for i := 0; i < n; i++ {
		if x[i] != y[i] {
			diffs = diffs + 1
		}
	}

	if diffs == 1 {
		return true
	} else {
		return false
	}
}

func get_same(x string, y string) []string {
	n := len(x)

	same := make([]string, n)
	for i := 0; i < n; i++ {
		if x[i] == y[i] {
			same = append(same, string(x[i]))
		}
	}

	return same
}

func main() {
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

	// Find two strings that differ by only one character
	for _, x := range lines {
		for _, y := range lines {
			check := compare(x, y)
			if check == true {
				same := get_same(x, y)
				fmt.Println(same)
			} else {
				continue
			}
		}
	}
}
