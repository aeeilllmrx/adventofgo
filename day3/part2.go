package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Square struct {
    x int
    y int
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

    // Map squares to number of claims
    m := make(map[Square]int)

    // Lines look like #1 @ 872,519: 18x18
    x_max := 0
    y_max := 0

    for _, curr := range lines {
        line := strings.Split(curr, " ")
        coords := line[2]
        n := len(coords)
        coords = coords[0 : n-1]

        parts := strings.Split(coords, ",")
        x, _ := strconv.Atoi(parts[0])
        y, _ := strconv.Atoi(parts[1])

        dims := strings.Split(line[3], "x")
        i, _ := strconv.Atoi(dims[0])
        j, _ := strconv.Atoi(dims[1])

        for row := 0; row < i; row++ {
            for col := 0; col < j; col++ {
                x_val := x + row
                y_val := y + col

                if x_val > x_max {
                    x_max = x_val
                }
                if y_val > y_max {
                    y_max = y_val
                }

                square := Square{x_val, y_val}
                m[square] = m[square] + 1
            }
        }
    }          

    // Now find the only claim that doesn't overlap
    for _, curr := range lines {
        line := strings.Split(curr, " ")
        coords := line[2]
        n := len(coords)
        coords = coords[0 : n-1]

        parts := strings.Split(coords, ",")
        x, _ := strconv.Atoi(parts[0])
        y, _ := strconv.Atoi(parts[1])

        dims := strings.Split(line[3], "x")
        i, _ := strconv.Atoi(dims[0])
        j, _ := strconv.Atoi(dims[1])

        alone := true

        for row := 0; row < i; row++ {
            for col := 0; col < j; col++ {
                x_val := x + row
                y_val := y + col

                if m[Square{x_val, y_val}] > 1 {
                    alone = false
                }
            }
        }

        if alone == true {
            fmt.Println(curr)
        }
    }
}







