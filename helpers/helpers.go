package helpers

import (
    "bufio"
    "fmt"
    "os"
    "io/ioutil"
)

// helper to check if string is in a slice
func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

// reads input file and turns into a slice of strings
func GetLines() []string {
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

// reads input file as a slice of bytes
func GetStringInput() string {
    data, err := ioutil.ReadFile("input.txt")

    if err != nil {
        fmt.Print(err)
    }

    return string(data)
}

// very hacky priority queue data structure
// use NewPQ to create the queue and NewItem to create new items
type Item struct {
    Value    string
    Priority int
    index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
    n := len(*pq)
    item := x.(*Item)
    item.index = n
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    item.index = -1 // for safety
    *pq = old[0 : n-1]
    return item
}

func NewPQ(size int) PriorityQueue {
    pq := make(PriorityQueue, size)
    return pq
}

func NewItem(val string, p int) *Item {
    item := Item {
        Value:    val,
        Priority: p,
    }
    
    return &item
}

