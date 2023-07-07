package second

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const HEIGHT = 6
const LENGTH = 40

func renderPosition(currLineIndex int, sum int) (int, int, byte) {
	row := (currLineIndex - 1) / LENGTH
	offset := (currLineIndex - 1) % LENGTH

	if offset >= sum-1 && offset <= sum+1 {
		return row, offset, '#'
	}
	return row, offset, '.'
}

func parseInstruction(line string) (bool, int, error) {
	if line[:4] != "addx" {
		return false, 0, nil
	}

	number, err := strconv.Atoi(line[5:])
	if err != nil {
		return false, 0, errors.New("Cannot convert to int")
	}

	return true, number, nil
}

func printResult(result [HEIGHT][LENGTH]byte) {
	for _, row := range result {
		for _, pixel := range row {
			fmt.Print(string(pixel))
		}
		fmt.Println()
	}
}

func Solve() error {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		return errors.New("Cannot open file")
	}
	defer file.Close()

	sum := 1
	pixels := [HEIGHT][LENGTH]byte{}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for currLineIndex := 1; fileScanner.Scan(); currLineIndex++ {
		row, offset, pixel := renderPosition(currLineIndex, sum)
		pixels[row][offset] = pixel

		isAddxInstruction, number, err := parseInstruction(fileScanner.Text())
		if err != nil {
			return err
		}
		if !isAddxInstruction {
			continue
		}

		currLineIndex++
		row, offset, pixel = renderPosition(currLineIndex, sum)
		pixels[row][offset] = pixel
		sum += number
	}

	printResult(pixels)
	return nil
}
