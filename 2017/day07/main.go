package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ExtractNumber(nstr string) int {
	cleaned := strings.TrimFunc(nstr, func(r rune) bool { return r == '(' || r == ')' })
	number, _ := strconv.Atoi(cleaned)
	return number
}

type Node struct {
	name      string
	childrens []*Node
}

type Tree []*Node

func (t Tree) FindRoot() *Node {
	nodeToParent := make(map[string]string, len(t))
	for _, node := range t {
		for _, child := range node.childrens {
			nodeToParent[child.name] = node.name
		}
	}
	for _, node := range t {
		// if not is set -> no parent -> root
		if _, isParent := nodeToParent[node.name]; !isParent {
			return node
		}
	}
	return nil
}

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	tree := make(Tree, len(lines))
	weights := make(map[string]int64, len(lines))
	// Build Tree (graph)
	for i, line := range lines {
		parts := strings.Split(line, " ")
		nParts := len(parts)
		if nParts > 2 {
			name := parts[0]
			weight := ExtractNumber(parts[1])
			weights[name] = int64(weight)
			childrens := make([]*Node, nParts-3)
			for j := 0; j < nParts-3; j++ {
				child := strings.TrimSuffix(strings.TrimSpace(parts[j+3]), ",")
				childNode := Node{name: child}
				childrens[j] = &childNode
			}
			node := Node{name: name, childrens: childrens}
			tree[i] = &node
		} else {
			name := parts[0]
			weight := ExtractNumber(parts[1])
			weights[name] = int64(weight)
			node := Node{name: name}
			tree[i] = &node
		}
	}
	// Find Root
	root := tree.FindRoot()
	fmt.Printf("Part 1: %s\n", root.name)
	// Part 2 - Balance the Tree from root
}
