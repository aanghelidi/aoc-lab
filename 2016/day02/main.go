package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(file)
	// Part 1
	//grid := make(map[Point]int, 9)
	//grid[Point{-1, 1}] = 1
	//grid[Point{0, 1}] = 2
	//grid[Point{1, 1}] = 3
	//grid[Point{-1, 0}] = 4
	//grid[Point{0, 0}] = 5
	//grid[Point{1, 0}] = 6
	//grid[Point{-1, -1}] = 7
	//grid[Point{0, -1}] = 8
	//grid[Point{1, -1}] = 9

	// Part 2
	grid := make(map[Point]int, 25)

	grid[Point{-2, 2}] = 0
	grid[Point{-1, 2}] = 0
	grid[Point{0, 2}] = 1
	grid[Point{1, 2}] = 0
	grid[Point{2, 2}] = 0

	grid[Point{-2, 1}] = 0
	grid[Point{-1, 1}] = 2
	grid[Point{0, 1}] = 3
	grid[Point{1, 1}] = 4
	grid[Point{2, 1}] = 0

	grid[Point{-2, 0}] = 5
	grid[Point{-1, 0}] = 6
	grid[Point{0, 0}] = 7
	grid[Point{1, 0}] = 8
	grid[Point{2, 0}] = 9

	grid[Point{-2, -1}] = 0
	grid[Point{-1, -1}] = 'A'
	grid[Point{0, -1}] = 'B'
	grid[Point{1, -1}] = 'C'
	grid[Point{2, -1}] = 0

	grid[Point{-2, -2}] = 0
	grid[Point{-1, -2}] = 0
	grid[Point{0, -2}] = 'D'
	grid[Point{1, -2}] = 0
	grid[Point{2, -2}] = 0

	instructions := make([]string, 0)
	for s.Scan() {
		line := s.Text()
		instructions = append(instructions, line)
	}
	// Part 1
	// start := Point{0, 0}
	start := Point{-2, 0}
	current := start

	total := 0
	for i := 0; i < len(instructions); i++ {
		total += len(instructions[i])
	}
	sb := strings.Builder{}
	sb.Grow(total)

	for _, instruction := range instructions {
		for _, c := range instruction {
			switch c {
			case 'U':
				// if _, ok := grid[Point{current.x, current.y + 1}]; ok{
				// 	current.y++
				// }

				if value, ok := grid[Point{current.x, current.y + 1}]; ok && value != 0 {
					current.y++
				}
			case 'D':
				//if _, ok := grid[Point{current.x, current.y - 1}]; ok {
				//	current.y--
				//}
				if value, ok := grid[Point{current.x, current.y - 1}]; ok && value != 0 {
					current.y--
				}
			case 'L':
				//if _, ok := grid[Point{current.x - 1, current.y}]; ok {
				//	current.x--
				//}
				if value, ok := grid[Point{current.x - 1, current.y}]; ok && value != 0 {
					current.x--
				}
			case 'R':
				//if _, ok := grid[Point{current.x + 1, current.y}]; ok {
				//	current.x++
				//}
				if value, ok := grid[Point{current.x + 1, current.y}]; ok && value != 0 {
					current.x++
				}
			}
		}
		if grid[current] >= 'A' && grid[current] <= 'D' {
			_, _ = sb.WriteString(string(grid[current]))
		} else {
			_, _ = sb.WriteString(fmt.Sprintf("%d", grid[current]))
		}
	}
	fmt.Println(sb.String())
}
