package second

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Solve() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		panic("Cannot open file")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	area := 0
	isPresent := map[string]bool{}
	isSurounded := map[string]bool{}
	neighbourCount := map[string]int{}

	for fileScanner.Scan() {
		var x, y, z int
		line := fileScanner.Text()
		count, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if err != nil || count != 3 {
			panic("Cannot parse line")
		}

		if isSurounded[line] {
			fmt.Println(line)
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
			neighbourCount[key]++
			if isPresent[key] {
				borderCount++
			} else if neighbourCount[key] == 6 {
				isSurounded[key] = true
				area -= 12
			}
		}

		area += 6 - 2*borderCount
	}

}
