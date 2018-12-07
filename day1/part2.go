package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	m := make(map[int]bool)
	m[0] = true

	str := string(b)
	total := 0
	lines := strings.Split(str, "\n")

	for i := 0; i < 10000; {
		element := lines[i]

		if len(element) > 0 {
			sign := element[0]
			num, _ := strconv.Atoi(element[1:])
			if sign == '+' {
				total = total + num
			}
			if sign == '-' {
				total = total - num
			}
			_, ok := m[total]
			if ok {
				print(total)
				return
			} else {
				m[total] = true
			}
			i = i + 1
		} else {
			i = 0
		}
	}
}
