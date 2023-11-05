package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Counter map[string]int
type Pair struct {
	letter string
	count  int
}

func (c Counter) Sort() []Pair {
	pairs := make([]Pair, len(c))
	for k, v := range c {
		pairs = append(pairs, Pair{string(k), v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count == pairs[j].count {
			return pairs[i].letter < pairs[j].letter
		}
		return pairs[i].count > pairs[j].count
	})
	return pairs
}

func (c Counter) RSort() []Pair {
	pairs := make([]Pair, 0)
	for k, v := range c {
		pairs = append(pairs, Pair{string(k), v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count == pairs[j].count {
			return pairs[i].letter < pairs[j].letter
		}
		return pairs[i].count < pairs[j].count
	})
	return pairs
}

func main() {
	//filename := "test.txt"
	filename := "input.txt"
	var n int
	if filename == "test.txt" {
		n = 6
	} else if filename == "input.txt" {
		n = 8
	} else {
		log.Fatal("Unknown filename")
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	counters := make([]Counter, n)
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			if counters[i] == nil {
				counters[i] = make(Counter)
			}
			counters[i][string(c)]++
		}
	}
	// Part 1 we choose the most common letter
	sb := strings.Builder{}
	sb.Grow(n)
	for _, counter := range counters {
		for _, pair := range counter.Sort() {
			sb.WriteString(pair.letter)
			break
		}
	}
	fmt.Printf("Part 1: %s\n", sb.String())
	// Part 2 we choose the least common letter
	sb.Reset()
	for _, counter := range counters {
		for _, pair := range counter.RSort() {
			sb.WriteString(pair.letter)
			break
		}
	}
	fmt.Printf("Part 2: %s\n", sb.String())
}
