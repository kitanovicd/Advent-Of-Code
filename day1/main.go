package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func updateTopThreeSums(first, second, third *int, value int) {
	switch {
	case value > *first:
		*third = *second
		*second = *first
		*first = value
	case value > *second:
		*third = *second
		*second = value
	case value > *third:
		*third = value
	}
}

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		panic("Cannot open file")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	first, second, third := 0, 0, 0
	currSum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			updateTopThreeSums(&first, &second, &third, currSum)
			currSum = 0
			continue
		}

		currentValue, err := strconv.Atoi(line)
		if err != nil {
			panic("Cannot convert string to int")
		}

		currSum += currentValue
	}

	updateTopThreeSums(&first, &second, &third, currSum)

	result := first + second + third
	fmt.Println("Result is", result)
}
