package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	weight    int
	childrens []string
}
type Graph map[string]Data

func ExtractNumber(nstr string) int {
	cleaned := strings.TrimFunc(nstr, func(r rune) bool { return r == '(' || r == ')' })
	number, _ := strconv.Atoi(cleaned)
	return number
}

func main() {
	data, _ := os.ReadFile("test.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	graph := make(Graph, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		nParts := len(parts)
		if nParts > 2 {
			name := parts[0]
			weight := ExtractNumber(parts[1])
			childrens := parts[3:nParts]
			graph[name] = Data{weight, childrens}
		} else {
			name := parts[0]
			weight := ExtractNumber(parts[1])
			graph[name] = Data{weight, nil}
		}
	}
	fmt.Printf("%v", graph)
}
