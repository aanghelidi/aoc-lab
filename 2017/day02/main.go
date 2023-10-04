package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var ans, ans2, n int
	for s.Scan() {
		line := s.Text()
		elms := strings.Fields(line)
		n = len(elms)
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			num, _ := strconv.Atoi(elms[i])
			nums[i] = num
			for j := 1; j < n; j++ {
				step := (i + j) % n
				num_step, _ := strconv.Atoi(elms[step])
				if num%num_step == 0 {
					ans2 += num / num_step
				}
			}
		}
		min_r, max_r := slices.Min(nums), slices.Max(nums)
		ans += max_r - min_r
	}
	fmt.Println(ans)
	fmt.Println(ans2)
}
