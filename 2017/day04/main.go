package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func isValid(words []string, n int) bool {
	seen, correct := make(map[string]struct{}, n), 0
	for _, word := range words {
		if _, ok := seen[word]; ok {
			continue
		}
		correct++
		seen[word] = struct{}{}
	}
	return correct == n
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	ans, ans2 := 0, 0
	for s.Scan() {
		line := s.Text()
		words := strings.Fields(line)
		n := len(words)
		sWords := make([]string, n)
		for i, word := range words {
			rWord := []rune(word)
			sort.Slice(rWord, func(i, j int) bool { return rWord[i] < rWord[j] })
			sortedWord := string(rWord)
			sWords[i] = sortedWord
		}
		if isValid(words, n) {
			ans++
		}
		if isValid(sWords, n) {
			ans2++
		}
	}
	fmt.Println(ans)
	fmt.Println(ans2)
}
