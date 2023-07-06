package main

import (
	"fmt"
	"os"
	"strconv"
)

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	filePath := "input.txt"
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Cannot parse argument")
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		panic("Cannot read from file")
	}

	startPos := 0
	lastAppereance := make(map[byte]int)

	for i := 0; i < len(data); i++ {
		from := 0
		if i >= n {
			from = i - n + 1
		}

		if lastAppereance[data[i]] >= from {
			startPos = max(startPos, lastAppereance[data[i]]+1)
			lastAppereance[data[i]] = i

			continue
		}
		lastAppereance[data[i]] = i

		if (i - startPos) == n-1 {
			fmt.Println("Result:", i+1)
			break
		}
	}
}
