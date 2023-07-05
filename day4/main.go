package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		panic("Cannot open file")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	count := 0
	for fileScanner.Scan() {
		var from1, to1, from2, to2 int
		num, err := fmt.Sscanf(fileScanner.Text(), "%d-%d,%d-%d", &from1, &to1, &from2, &to2)

		if err != nil || num != 4 {
			panic("Cannot parse line")
		}

		if (from1 >= from2 && from1 <= to2) || (from2 >= from1 && from2 <= to1) {
			count++
		}
	}

	fmt.Println(count)
}
