package main

import (
	"fmt"
	"math"
)

func findLayer(input int) (int, int) {
	layerSquares, layer := 0, 1
	for {
		layerSquares = int(math.Pow(2*float64(layer)-1, 2))
		if input <= layerSquares {
			return layer, 2*layer - 1
		}
		layer++
	}
}

func computeCorners(layer int) (int, int, int, int) {
	bottomRight := int(math.Pow(2*float64(layer)-1, 2))
	bottomLeft := bottomRight - (2*layer - 2)
	topLeft := bottomLeft - (2*layer - 2)
	topRight := topLeft - (2*layer - 2)
	return bottomRight, bottomLeft, topLeft, topRight
}

func RetrieveSidePosition(input int, br int, bl int, tl int, tr int) (string, int) {
	if input >= tr && input <= tl {
		return "top", tl - input
	} else if input >= tl && input <= bl {
		return "left", bl - input
	} else if input >= bl && input <= br {
		return "bottom", br - input
	} else {
		return "right", tr - input
	}
}

func Abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

type Point struct {
	x int
	y int
}

type Grid map[Point]int

func (p *Point) Neighbours() []Point {
	return []Point{
		{p.x + 1, p.y},
		{p.x, p.y + 1},
		{p.x - 1, p.y},
		{p.x, p.y - 1},
		{p.x + 1, p.y + 1},
		{p.x - 1, p.y + 1},
		{p.x - 1, p.y - 1},
		{p.x + 1, p.y - 1},
	}
}

func Move(p Point, direction string) Point {
	switch direction {
	case "top":
		return Point{p.x, p.y + 1}
	case "bottom":
		return Point{p.x, p.y - 1}
	case "left":
		return Point{p.x - 1, p.y}
	case "right":
		return Point{p.x + 1, p.y}
	}
	return p
}

func IsRightCorner(p Point, layer int) bool {
	if p.x == layer && p.y == -layer {
		return true
	}
	return false
}

func CheckChangeDirection(p Point, direction string, layer int) string {
	if IsRightCorner(p, layer) {
		return "right"
	}
	nextCurrent := Move(p, direction)
	if Abs(nextCurrent.x, 0) > layer || Abs(nextCurrent.y, 0) > layer {
		switch direction {
		case "top":
			return "left"
		case "bottom":
			return "right"
		case "left":
			return "bottom"
		case "right":
			return "top"
		}
	}
	return direction
}

func main() {
	input := 277678

	// Part 1
	layer, dimension := findLayer(input)
	bottomRight, bottomLeft, topLeft, topRight := computeCorners(layer)
	_, pos := RetrieveSidePosition(input, bottomRight, bottomLeft, topLeft, topRight)
	middleSide := (dimension / 2) + 1
	distanceToMid := Abs(middleSide, pos)
	fmt.Println("Part 1:", distanceToMid+layer-2)

	// Part 2
	grid := make(Grid)
	grid[Point{0, 0}] = 1
	start := Point{0, 0}
	current := start
	layer = 1
	direction := "right"
	for {
		for _, p := range current.Neighbours() {
			if _, ok := grid[p]; !ok {
				continue
			}
			grid[current] += grid[p]
		}
		if grid[current] > input {
			fmt.Println("Part 2:", grid[current])
			break
		}
		current = Move(current, direction)
		direction = CheckChangeDirection(current, direction, layer)
		if direction == "right" {
			if IsRightCorner(current, layer) {
				layer++
			}
		}
	}
}
