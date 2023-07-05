package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	roundPoints := map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}

	filePath := "index.txt"

	file, err := os.Open(filePath)
	if err != nil {
		panic("Cannot open file")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	totalPoints := 0
	for fileScanner.Scan() {
		totalPoints += roundPoints[fileScanner.Text()]
	}

	fmt.Println("Total points:", totalPoints)
}
