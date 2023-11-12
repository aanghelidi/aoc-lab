package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func supportTLS(line string) bool {
	abba := false
	for i := 0; i < len(line)-3; i++ {
		if line[i] == line[i+3] && line[i+1] == line[i+2] && line[i] != line[i+1] {
			abba = true
		}
	}
	hypernet := false
	for i := 0; i < len(line)-3; i++ {
		if line[i] == ']' {
			hypernet = false
		}
		if line[i] == '[' {
			hypernet = true
		}
		if line[i] == line[i+3] && line[i+1] == line[i+2] && line[i] != line[i+1] && hypernet {
			return false
		}
	}
	return abba
}

func supportSSL(line string) bool {
	// Split by [ or ]
	hypernets := regexp.MustCompile(`\[(.*?)\]`).FindAllStringSubmatch(line, -1)
	supernets := regexp.MustCompile(`^(.*?)\[|\](.*?)\[|\](.*?)$`).FindAllStringSubmatch(line, -1)
	hyperenetsclean := make([]string, 0)
	for _, matches := range hypernets {
		for _, hypernet := range matches {
			if hypernet != "" {
				if strings.Contains(hypernet, "[") || strings.Contains(hypernet, "]") {
					continue
				}
				hyperenetsclean = append(hyperenetsclean, hypernet)
			}
		}
	}
	supernetsclean := make([]string, 0)
	for _, matches := range supernets {
		for _, supernet := range matches {
			if supernet != "" {
				if strings.Contains(supernet, "[") || strings.Contains(supernet, "]") {
					continue
				}
				supernetsclean = append(supernetsclean, supernet)
			}
		}
	}
	for _, supernet := range supernetsclean {
		for i := 0; i < len(supernet)-2; i++ {
			if supernet[i] == supernet[i+2] && supernet[i] != supernet[i+1] {
				babb := strings.Builder{}
				babb.Grow(3)
				babb.WriteByte(supernet[i+1])
				babb.WriteByte(supernet[i])
				babb.WriteByte(supernet[i+1])
				for _, hypernet := range hyperenetsclean {
					if strings.Contains(hypernet, babb.String()) {
						return true
					}
				}
			}
		}
	}

	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	ans := 0
	ans2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		if supportTLS(line) {
			ans++
		}
		if supportSSL(line) {
			ans2++
		}
	}
	fmt.Printf("Part 1: %d\n", ans)
	fmt.Printf("Part 2: %d\n", ans2)
}
