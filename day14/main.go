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

var OFFSET = 465
var EXTENSION = 5000
var BEGINNING = Coordinate{0, 35 + EXTENSION/2}

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

func dropDown(coord Coordinate, field [][]byte) Coordinate {
	for coord.X+1 < len(field) && field[coord.X+1][coord.Y] == 0 {
		coord.X++
	}

	if coord.X == len(field) {
		panic("Ball is dropped in void!!!!!")
	}

	return coord
}

func dropLeftOrRight(coord Coordinate, field [][]byte) (Coordinate, error) {
	x := coord.X + 1
	y := coord.Y - 1

	if x >= len(field) || y < 0 {
		return Coordinate{0, 0}, errors.New("Ball is dropped in void on left side")
	}
	if field[x][y] == 0 {
		return Coordinate{x, y}, nil
	}

	y = coord.Y + 1
	if y >= len(field[0]) {
		return Coordinate{0, 0}, errors.New("Ball is dropped in void on right side")
	}
	if field[x][y] == 0 {
		return Coordinate{x, y}, nil
	}

	return coord, nil
}

func drop(coord Coordinate, field [][]byte) Coordinate {
	var err error

	nextCoord := dropDown(coord, field)
	nextCoord, err = dropLeftOrRight(nextCoord, field)
	if err != nil {
		return BEGINNING
	}

	return nextCoord
}

func playMove(field [][]byte) bool {
	currCoord := BEGINNING
	nextCoord := drop(currCoord, field)
	for ; nextCoord != currCoord && nextCoord != BEGINNING; nextCoord = drop(currCoord, field) {
		currCoord = nextCoord
	}

	if nextCoord != BEGINNING {
		field[currCoord.X][currCoord.Y] = 'O'
		return true
	}

	return false
}

func findCoordinates(line1, line2 string) (Coordinate, Coordinate, error) {
	fromX, fromY, err := getCoordinates(line1)
	if err != nil {
		return Coordinate{0, 0}, Coordinate{0, 0}, errors.New("Cannot get coordinates for line1")
	}

	toX, toY, err := getCoordinates(line2)
	if err != nil {
		return Coordinate{0, 0}, Coordinate{0, 0}, errors.New("Cannot get coordinates for line2")
	}

	if fromX > toX {
		fromX, toX = toX, fromX
	}
	if fromY > toY {
		fromY, toY = toY, fromY
	}

	return Coordinate{fromX, fromY}, Coordinate{toX, toY}, nil
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

	var field [][]byte
	for i := 0; i < 169; i++ {
		field = append(field, make([]byte, EXTENSION+64))
	}
	for i := 0; i < len(field[0]); i++ {
		field[len(field)-1][i] = '#'
	}

	for fileScanner.Scan() {
		values := strings.Split(fileScanner.Text(), " -> ")

		for i := 0; i < len(values)-1; i++ {
			from, to, err := findCoordinates(values[i], values[i+1])
			if err != nil {
				panic(err)
			}

			for x := from.X; x <= to.X; x++ {
				for y := from.Y; y <= to.Y; y++ {
					field[x][y-OFFSET+EXTENSION/2] = '#'
				}
			}
		}
	}

	count := 0
	for playMove(field) {
		count++
	}

	fmt.Println(count)
}
