package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Room struct {
	name     string
	rawName  string
	sectorID int
	checksum string
}

type Pair struct {
	letter string
	count  int
}

func (r *Room) MostCommonLetters() map[rune]int {
	counter := make(map[rune]int)
	for _, r := range r.name {
		if unicode.IsLetter(r) {
			counter[r]++
		}
	}
	return counter
}

func parseRoom(line string) Room {
	parts := strings.Split(line, "[")
	nameID := strings.Split(parts[0], "-")
	n := len(nameID)
	id := nameID[n-1]
	name := strings.Join(nameID[:n-1], "")
	sectorID, _ := strconv.Atoi(id)
	checksum := strings.TrimRight(parts[1], "]")
	return Room{name, parts[0], sectorID, checksum}
}

func decryptName(rawName string) string {
	// Extract sector ID
	parts := strings.Split(rawName, "-")
	n := len(parts)
	id := parts[n-1]
	sectorID, _ := strconv.Atoi(id)
	// Extract Name
	name := parts[:n-1]
	// Count number of bytes
	total := 0
	for i := 0; i < len(name); i++ {
		total += len(name[i])
	}
	sb := strings.Builder{}
	sb.Grow(total)
	for _, word := range name {
		for _, r := range word {
			sb.WriteRune('a' + rune((int(r)-'a'+sectorID)%26))
		}
		sb.WriteRune(' ')
	}
	return strings.TrimSpace(sb.String())
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		room := parseRoom(line)
		counter := room.MostCommonLetters()
		// Create slice of pairs to sort by count and then letter
		pairs := make([]Pair, len(counter))
		for r, c := range counter {
			pairs = append(pairs, Pair{string(r), c})
		}
		sort.Slice(pairs, func(i, j int) bool {
			if pairs[i].count == pairs[j].count {
				return pairs[i].letter < pairs[j].letter
			}
			return pairs[i].count > pairs[j].count
		})
		// Build checksum
		sb := strings.Builder{}
		sb.Grow(5)
		count := 0
		for _, p := range pairs {
			if p.count > 0 && count < 5 {
				sb.WriteString(p.letter)
				count++
			}
		}
		// Part 1: Add sector ID to answer if checksum matches
		if sb.String() == room.checksum {
			ans += room.sectorID
		}

		// Part 2: Decrypt name
		decrypted := decryptName(room.rawName)
		if strings.Contains(decrypted, "north") {
			fmt.Printf("Part 2: %s %d\n", decrypted, room.sectorID)
		}
	}
	fmt.Printf("Part 1: %d\n", ans)
}
