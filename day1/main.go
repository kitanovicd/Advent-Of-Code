package main

import (
	"bufio"
	"os"
	"strconv"
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

	first := 0
	second := 0
	third := 0
	currSum := 0

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			switch {
			case currSum > first:
				third = second
				second = first
				first = currSum
			case currSum > second:
				third = second
				second = currSum
			case currSum > third:
				third = currSum
			}
			currSum = 0
			continue
		}

		currentValue, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			panic("Cannot convert string to int")
		}

		currSum += currentValue
	}

	switch {
	case currSum > first:
		third = second
		second = first
		first = currSum
	case currSum > second:
		third = second
		second = currSum
	case currSum > third:
		third = currSum
	}

	println("Result is", first+second+third)
}
