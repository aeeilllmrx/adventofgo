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
// find guard that spends most time asleep, and the minute most reliably asleep

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
	return lines
}

func sort_lines(lines []string) []string {
	sort.Strings(lines)
	return lines
}

func get_sleep_times(sorted_lines []string) map[string]int{
	time_asleep := make(map[string]int)

	begin_asleep := 0
	guard := ""

	for _, line := range sorted_lines {

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
			time_asleep[guard] += (minute - begin_asleep)
		}
	}

	return time_asleep
}

func get_guard(time_asleep map[string]int) string {
	sleepiest_guard := ""
	max_time := 0

	for guard, time := range time_asleep { 
    	if time > max_time {
    		sleepiest_guard = guard
    		max_time = time
    	}
	}
	
	return sleepiest_guard[1:]
}

func get_guard_lines(guard string, lines []string) []string {
	var guard_lines []string

	is_guard := false

	for _, line := range lines {
		if strings.Contains(line, guard) {
			is_guard = true
			guard_lines = append(guard_lines, line)
		} else if strings.Contains(line, "#") {
			is_guard = false
		} else if is_guard == true {
			guard_lines = append(guard_lines, line)
		}
	}

	return guard_lines
}

func get_minute_count(lines []string) map[int]int {
	minute_count := make(map[int]int)
    begin_asleep := 0

	for _, line := range lines {

		date := line[1:17]
		time := strings.Split(date, " ")[1]
		minute, _ := strconv.Atoi(time[3:5])
		text := line[19:len(line)]
		
		if strings.HasPrefix(text, "Guard") {
			begin_asleep = 0
		} else if text == "falls asleep" {
			begin_asleep = minute
		} else {
			for i := begin_asleep; i < minute; i++ {
				minute_count[i] += 1
			}
		}
	}

	return minute_count
}

func get_minute(minute_count map[int]int) int {
	sleepiest_minute := 0
	max_count := 0

	for minute, count := range minute_count { 
    	if count > max_count {
    		sleepiest_minute = minute
    		max_count = count
    	}
	}
	
	return sleepiest_minute
}

func main() {
	lines := get_input()

	sorted_lines := sort_lines(lines)

	time_asleep := get_sleep_times(sorted_lines)

	sleepiest_guard := get_guard(time_asleep)

	guard_lines := get_guard_lines(sleepiest_guard, sorted_lines)

	minute_count := get_minute_count(guard_lines)

	sleepiest_minute := get_minute(minute_count)

	guard_id, _ := strconv.Atoi(sleepiest_guard)
	fmt.Print(guard_id * sleepiest_minute, "\n")

}