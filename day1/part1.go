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

	str := string(b)

	total := 0
	lines := strings.Split(str, "\n")

	for _, element := range lines {
		if len(element) > 0 {
			sign := element[0]
			num, _ := strconv.Atoi(element[1:])
			if sign == '+' {
				total = total + num
			}
			if sign == '-' {
				total = total - num
			}
		}
	}

	fmt.Println(total)
}
