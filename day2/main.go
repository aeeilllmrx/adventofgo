package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	two_total := 0
	three_total := 0
	lines := strings.Split(str, "\n")

	for _, element := range lines {
		if len(element) > 0 {
			distinct_letters := []string{}
			for _, letter := range element {
				new_letter := string(letter)
				flag := 0
				for _, uniq := range distinct_letters {
					if new_letter == uniq {
						flag = 1
					}
				}
				if flag == 0 {
					distinct_letters = append(distinct_letters, new_letter)
				}
			}

			for _, letter := range distinct_letters {
				num_letter := strings.Count(element, letter)
				if num_letter == 2 {
					two_total = two_total + 1
					break
				}
			}
			for _, letter := range distinct_letters {
				num_letter := strings.Count(element, letter)
				if num_letter == 3 {
					three_total = three_total + 1
					break
				}
			}
		}
	}
	total := two_total * three_total
	fmt.Println(total)
}
