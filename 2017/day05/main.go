package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var part1 = flag.Bool("part-1", false, "Solution for part 1")
	flag.Parse()
	data, _ := os.ReadFile("input.txt")
	sdata := strings.Split(strings.TrimSpace(string(data)), "\n")
	n := len(sdata)
	offsets := make([]int, n)
	for i, d := range sdata {
		offset, _ := strconv.Atoi(d)
		offsets[i] = offset
	}
	previousPosition, position, steps := 0, 0, 0
	for {
		coffset := offsets[position]
		position += coffset
		if *part1 {
			offsets[previousPosition]++
		} else {
			if coffset >= 3 {
				offsets[previousPosition]--
			} else {
				offsets[previousPosition]++
			}
		}
		previousPosition = position
		steps++
		if position >= n {
			break
		}
	}
	fmt.Println(steps)
}
