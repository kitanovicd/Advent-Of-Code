package second

import (
	"bufio"
	"os"

	"github.com/kitanovicd/Advent-Of-Code/Advent-Of-Code/day9/first"
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
		return 0, err
	}
	defer file.Close()

	result := 1
	head := coordinate{0, 0}
	tails := [9]coordinate{}
	visited := make(map[coordinate]bool)
	visited[head] = true

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		xDir, yDir, numberOfSteps, err := first.ParseLine(fileScanner.Text())
		if err != nil {
			return 0, err
		}

		for i := 0; i < numberOfSteps; i++ {
			head.x += xDir
			head.y += yDir

			curr := head
			for j := 0; j < 9; j++ {
				tails[j] = moveTail(curr, tails[j])
				curr = tails[j]
			}

			if !visited[curr] {
				visited[curr] = true
				result++
			}
		}
	}

	return result, nil
}
