package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"strconv"
)

// input format: [1518-06-23 00:43] wakes up
// which guard/minute combination is most frequent

func get_input() []string {
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

    sort.Strings(lines)
	
    return lines
}

func get_most_common(guard_minutes map[string]int) string {
	most_common := ""
	max_count := 0

	for key, count := range guard_minutes { 
    	if count > max_count {
    		most_common = key
    		max_count = count
    	}
	}
	
	return most_common
}

func get_guard_minutes(lines []string) map[string]int {
	guard_minute_map := make(map[string]int)
    begin_asleep := 0
    guard := ""
    tmp := ""

	for _, line := range lines {

		date := line[1:17]
		time := strings.Split(date, " ")[1]
		minute, _ := strconv.Atoi(time[3:5])
		text := line[19:len(line)]
		
		if strings.HasPrefix(text, "Guard") {
            guard = strings.Split(text, " ")[1]
			begin_asleep = 0
		} else if text == "falls asleep" {
			begin_asleep = minute
		} else {
			for i := begin_asleep; i < minute; i++ {
                tmp = strconv.Itoa(i)
				guard_minute_map[guard + " " + tmp] += 1
			}
		}
	}

	return guard_minute_map
}

func main() {
	sorted_lines := get_input()

	guard_minutes := get_guard_minutes(sorted_lines)

	most_common := get_most_common(guard_minutes)

    parts := strings.Split(most_common, " ")
    guard, _ := strconv.Atoi(parts[0][1:])
    minute, _ := strconv.Atoi(parts[1])

	fmt.Print(most_common, "\n", guard * minute)

}