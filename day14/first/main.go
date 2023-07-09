package first

import (
	"bufio"
	"errors"
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

func dropDown(coord Coordinate, matrix *[][]byte) Coordinate {
	data := *matrix
	for coord.X+1 < len(data) && data[coord.X+1][coord.Y] == 0 {
		coord.X++
	}

	if coord.X == len(data) {
		panic("Ball is dropped in void!!!!!")
	}

	return coord
}

func dropLeftOrRight(coord Coordinate, matrix *[][]byte) (Coordinate, error) {
	x := coord.X + 1
	y := coord.Y - 1
	data := *matrix

	if x >= len(data) || y < 0 {
		return Coordinate{0, 0}, errors.New("Ball is dropped in void on left side")
	}
	if data[x][y] == 0 {
		return Coordinate{x, y}, nil
	}

	y = coord.Y + 1
	if y >= len(data[0]) {
		return Coordinate{0, 0}, errors.New("Ball is dropped in void on right side")
	}
	if data[x][y] == 0 {
		return Coordinate{x, y}, nil
	}

	return coord, nil
}

func drop(coord Coordinate, data *[][]byte) Coordinate {
	var err error

	nextCoord := dropDown(coord, data)
	nextCoord, err = dropLeftOrRight(nextCoord, data)
	if err != nil {
		return BEGINNING
	}

	return nextCoord
}

func playMove(data *[][]byte) bool {
	currCoord := BEGINNING
	nextCoord := drop(currCoord, data)
	for ; nextCoord != currCoord && nextCoord != BEGINNING; nextCoord = drop(currCoord, data) {
		currCoord = nextCoord
	}

	if nextCoord != BEGINNING {
		(*data)[currCoord.X][currCoord.Y] = 'O'
		return true
	}

	return false
}

func printData(data *[][]byte) {
	for index, row := range *data {
		if index < 10 {
			print(index, "  ")
		} else if index < 100 {
			print(index, " ")
		} else {
			print(index, "")
		}

		for _, col := range row {
			if col == 0 {
				print(".")
				continue
			}
			print(string(col))
		}
		println()
	}
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

func Solve() (int, error) {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		return 0, errors.New("Cannot open file")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var data [][]byte
	for i := 0; i < 167; i++ {
		data = append(data, make([]byte, 64))
	}

	for fileScanner.Scan() {
		values := strings.Split(fileScanner.Text(), " -> ")

		for i := 0; i < len(values)-1; i++ {
			from, to, err := findCoordinates(values[i], values[i+1])
			if err != nil {
				return 0, err
			}

			for x := from.X; x <= to.X; x++ {
				for y := from.Y; y <= to.Y; y++ {
					data[x][y-465] = '#'
				}
			}
		}
	}

	count := 0
	for playMove(&data) {
		count++
	}

	return count, nil
}
