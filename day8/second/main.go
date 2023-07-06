package second

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

func parseFile(file *os.File) ([][]int, error) {
	data := make([][]int, 1)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanBytes)

	for fileScanner.Scan() {
		if fileScanner.Text() == "\n" {
			data = append(data, []int{})
			continue
		}

		height, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			return nil, err
		}

		data[len(data)-1] = append(data[len(data)-1], height)
	}

	return data, nil
}

func processUp(data [][]int, row, column int) int {
	for i := row - 1; i >= 0; i-- {
		if data[i][column] >= data[row][column] {
			return row - i
		}
	}

	return row
}

func processDown(data [][]int, row, column int) int {
	for i := row + 1; i < len(data); i++ {
		if data[i][column] >= data[row][column] {
			return i - row
		}
	}

	return len(data) - row - 1
}

func processLeft(data [][]int, row, column int) int {
	for i := column - 1; i >= 0; i-- {
		if data[row][i] >= data[row][column] {
			return column - i
		}
	}

	return column
}

func processRight(data [][]int, row, column int) int {
	for i := column + 1; i < len(data[row]); i++ {
		if data[row][i] >= data[row][column] {
			return i - column
		}
	}

	return len(data[row]) - column - 1
}

func Solve() (int, error) {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		return 0, errors.New("Can't open file")
	}
	defer file.Close()

	data, err := parseFile(file)
	if err != nil {
		return 0, errors.New("Can't parse file")
	}

	max := 0
	for i := 1; i < len(data)-1; i++ {
		for j := 1; j < len(data[i])-1; j++ {
			up := processUp(data, i, j)
			down := processDown(data, i, j)
			left := processLeft(data, i, j)
			right := processRight(data, i, j)

			sum := up * down * left * right
			if sum > max {
				max = sum
			}
		}
	}

	return max, nil
}
