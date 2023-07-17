package first

import (
	"bufio"
	"errors"
	"fmt"
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

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	isPresent := map[string]bool{}
	area := 0

	for fileScanner.Scan() {
		var x, y, z int
		line := fileScanner.Text()
		count, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if err != nil || count != 3 {
			return 0, errors.New("Cannot parse line")
		}

		if isPresent[line] {
			continue
		}

		isPresent[line] = true
		borderCount := 0
		for _, direction := range []struct {
			x, y, z int
		}{
			{x - 1, y, z},
			{x + 1, y, z},
			{x, y - 1, z},
			{x, y + 1, z},
			{x, y, z - 1},
			{x, y, z + 1},
		} {
			key := strconv.Itoa(direction.x) + "," + strconv.Itoa(direction.y) + "," + strconv.Itoa(direction.z)
			if isPresent[key] {
				borderCount++
			}

		}

		area += 6 - 2*borderCount
	}

	return area, nil
}
