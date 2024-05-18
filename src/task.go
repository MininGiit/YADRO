package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n int
	fmt.Fscan(in, &n)
	containerSize := make([]int, n)
	ballsCount := make([]int, n)
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			var temp int
			fmt.Fscan(in, &temp)
			ballsCount[j] += temp
			count += temp
		}
		containerSize[i] = count
	}
	sort.Slice(containerSize, func(i, j int) bool {
		return containerSize[i] < containerSize[j]
	})
	sort.Slice(ballsCount, func(i, j int) bool {
		return ballsCount[i] < ballsCount[j]
	})
	result := true
	for i := 0; i < n; i++ {
		if ballsCount[i] != containerSize[i] {
			result = false
			break
		}
	}
	if result {
		fmt.Fprintln(out, "yes")
	} else {
		fmt.Fprintln(out, "no")
	}
}