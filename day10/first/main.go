package first

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

func Solve() (int, error) {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		return 0, errors.New("Cannot open file")
	}
	defer file.Close()

	sum := 1
	result := 0
	bound := 20

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for currLineIndex := 1; fileScanner.Scan(); currLineIndex++ {
		if currLineIndex == bound {
			result += sum * bound
			bound += 40
		}

		line := fileScanner.Text()
		if line[:4] != "addx" {
			continue
		}

		number, err := strconv.Atoi(line[5:])
		if err != nil {
			return 0, errors.New("Cannot convert to int")
		}

		if currLineIndex == bound-1 {
			result += sum * bound
			bound += 40
		}

		sum += number
		currLineIndex++
	}

	return result, nil
}
