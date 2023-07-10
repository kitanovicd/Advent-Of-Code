package first

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var BORDER = 2000000

type coordinate struct {
	X, Y int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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

	count := 0
	flaged := map[int]byte{}

	for fileScanner.Scan() {
		var sensorCoord, beaconCoord coordinate
		_, err := fmt.Sscanf(fileScanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorCoord.Y, &sensorCoord.X, &beaconCoord.Y, &beaconCoord.X)
		if err != nil {
			return 0, errors.New("Cannot parse line")
		}

		distance := abs(sensorCoord.X-beaconCoord.X) + abs(sensorCoord.Y-beaconCoord.Y)
		if sensorCoord.X == BORDER {
			if flaged[sensorCoord.Y] == '#' {
				count--
			}
			flaged[sensorCoord.Y] = 'S'
		}
		if beaconCoord.X == BORDER {
			if flaged[beaconCoord.Y] == '#' {
				count--
			}
			flaged[beaconCoord.Y] = 'B'
		}

		if abs(sensorCoord.X-BORDER) <= distance {
			for i := 0; i <= distance-abs(sensorCoord.X-BORDER); i++ {
				if flaged[sensorCoord.Y-i] == 0 {
					flaged[sensorCoord.Y-i] = '#'
					count++
				}
				if flaged[sensorCoord.Y+i] == 0 {
					flaged[sensorCoord.Y+i] = '#'
					count++
				}
			}
		}
	}

	return count, nil
}
