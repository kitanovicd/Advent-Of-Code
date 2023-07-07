package first

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

type coordinate struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func ParseLine(line string) (int, int, int, error) {
	direction := line[0]
	numberOfSteps, err := strconv.Atoi(line[2:])
	if err != nil {
		return 0, 0, 0, err
	}

	switch direction {
	case 'U':
		return 0, 1, numberOfSteps, nil
	case 'D':
		return 0, -1, numberOfSteps, nil
	case 'L':
		return -1, 0, numberOfSteps, nil
	case 'R':
		return 1, 0, numberOfSteps, nil
	}

	return 0, 0, 0, errors.New("Invalid direction")
}

func areHeadAndTailConnected(head, tail coordinate) bool {
	xDistance := head.x - tail.x
	yDistance := head.y - tail.y
	return abs(xDistance) <= 1 && abs(yDistance) <= 1
}

func moveTail(head, tail coordinate) coordinate {
	if areHeadAndTailConnected(head, tail) {
		return tail
	}

	if head.y > tail.y {
		tail.y++
	} else if head.y < tail.y {
		tail.y--
	}

	if head.x > tail.x {
		tail.x++
	} else if head.x < tail.x {
		tail.x--
	}

	return tail
}

func Solve() (int, error) {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		return 0, errors.New("Cannot open file")
	}
	defer file.Close()

	result := 1
	head := coordinate{0, 0}
	tail := coordinate{0, 0}
	visited := make(map[coordinate]bool)
	visited[tail] = true

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		xDir, yDir, numberOfSteps, err := ParseLine(fileScanner.Text())
		if err != nil {
			return 0, errors.New("Cannot parse line")
		}

		for i := 0; i < numberOfSteps; i++ {
			head.x += xDir
			head.y += yDir

			tail = moveTail(head, tail)
			if !visited[tail] {
				visited[tail] = true
				result++
			}
		}
	}

	return result, nil
}
