package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

var BEGINNING = Coordinate{0, 35}

func getCoordinates(number string) (int, int, error) {
	currCoordStr := strings.Split(number, ",")

	y, err := strconv.Atoi(currCoordStr[0])
	if err != nil {
		return 0, 0, errors.New("Cannot convert string to int for X coordinate")
	}

	x, err := strconv.Atoi(currCoordStr[1])
	if err != nil {
		return 0, 0, errors.New("Cannot convert string to int for Y coordinate")
	}
	return x, y, nil
}

func findBorders() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		panic("Cannot open file")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	toXMax := 0
	toYMax := 0
	fromYMin := 1000000000

	for fileScanner.Scan() {
		values := strings.Split(fileScanner.Text(), " -> ")

		for i := 0; i < len(values)-1; i++ {
			fromX, fromY, err := getCoordinates(values[i])
			if err != nil {
				panic(err)
			}

			toX, toY, err := getCoordinates(values[i+1])
			if err != nil {
				panic(err)
			}

			if fromX > toX {
				fromX, toX = toX, fromX
			}

			if fromY > toY {
				fromY, toY = toY, fromY
			}

			if toX > toXMax {
				toXMax = toX
			}

			if toY > toYMax {
				toYMax = toY
			}

			if fromY < fromYMin {
				fromYMin = fromY
			}
		}
	}

	fmt.Println("To X max:", toXMax)
	fmt.Println("From Y min:", fromYMin)
	fmt.Println("To Y max:", toYMax)

}

func main() {
	findBorders()
}
