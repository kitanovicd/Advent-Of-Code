package first

import (
	"bufio"
	"os"
	"strconv"
)

type coordinate struct {
	row    int
	column int
}

func isBorder(row, column, height, width int) bool {
	return (row == 0) || (column == 0) || (row == height-1) || (column == width-1)
}

func count(counted *map[coordinate]bool, row, column int) int {
	if (!(*counted)[coordinate{row, column}]) {
		(*counted)[coordinate{row, column}] = true
		return 1
	}
	return 0
}

func parseFile(file *os.File) (*[][]int, *map[coordinate]bool, int, error) {
	row := 0
	column := 0
	result := 0
	data := make([][]int, 1)
	counted := make(map[coordinate]bool)
	maxHeightByRow := make(map[int]int)
	maxHeightByColumn := make(map[int]int)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanBytes)

	for fileScanner.Scan() {
		if fileScanner.Text() == "\n" {
			row++
			column = 0
			data = append(data, []int{})
			continue
		}

		height, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			return nil, nil, 0, err
		}

		data[row] = append(data[row], height)

		if height > maxHeightByRow[row] {
			maxHeightByRow[row] = height
			result += count(&counted, row, column)
		}

		if height > maxHeightByColumn[column] {
			maxHeightByColumn[column] = height
			result += count(&counted, row, column)
		}

		column++
	}

	return &data, &counted, result, nil
}

func processData(data [][]int, counted map[coordinate]bool) int {
	result := 0
	maxHeightByRow := make(map[int]int)
	maxHeightByColumn := make(map[int]int)

	for row := len(data) - 1; row >= 0; row-- {
		for column := len(data[row]) - 1; column >= 0; column-- {
			height := data[row][column]

			if !counted[coordinate{row, column}] && isBorder(row, column, len(data), len(data[row])) {
				result++
				counted[coordinate{row, column}] = true
			}

			if height > maxHeightByRow[row] {
				maxHeightByRow[row] = height
				result += count(&counted, row, column)
			}

			if height > maxHeightByColumn[column] {
				maxHeightByColumn[column] = height
				result += count(&counted, row, column)
			}
		}
	}

	return result
}

func Solve() (int, error) {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	data, counted, result, err := parseFile(file)
	if err != nil {
		return 0, err
	}

	result += processData(*data, *counted)
	return result, nil
}
