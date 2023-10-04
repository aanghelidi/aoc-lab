package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var part = flag.Int("part", 1, "Part of the problem")
	flag.Parse()
	var ans, step int
	for s.Scan() {
		line := s.Text()
		n := len(line)
		if *part == 1 {
			step = 1
		} else {
			step = n / 2
		}
		for i := 0; i < n; i++ {
			cur := line[i]
			next := line[(i+step)%n]
			if cur == next {
				n, _ := strconv.Atoi(string(cur))
				ans += n
			}
		}
	}
	fmt.Println(ans)
}
