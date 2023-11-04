package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func retrievePassword(puzzleInput string, part2 bool) string {
	index := 0
	sb := strings.Builder{}
	sb.Grow(8)
	count := 0
	password := make([]byte, 8)
	for {
		stringToHash := puzzleInput + strconv.Itoa(index)
		hash := md5.Sum([]byte(stringToHash))
		hex := hex.EncodeToString(hash[:])
		if hex[0:5] == "00000" {
			if part2 {
				if unicode.IsLetter(rune(hex[5])) {
					index++
					continue
				}
				pos, err := strconv.Atoi(string(hex[5]))
				if err != nil {
					log.Fatal(err)
				}
				passwordChar := hex[6]
				if pos > 7 || password[pos] != 0 {
					index++
					continue
				}
				password[pos] = passwordChar
				count++
				if count == 8 {
					break
				}
			} else {
				sb.WriteByte(hex[5])
				count++
				if count == 8 {
					break
				}
			}
		}
		index++
	}

	if part2 {
		return string(password)
	} else {
		return sb.String()
	}
}

func main() {
	puzzleInput := "reyedfim"
	ans := retrievePassword(puzzleInput, false)
	fmt.Printf("Part 1: %s\n", ans)
	ans = retrievePassword(puzzleInput, true)
	fmt.Printf("Part 2: %s\n", ans)
}
