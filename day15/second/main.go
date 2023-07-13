package second

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type coordinate struct {
	X, Y int
}

type sensor struct {
	coord coordinate
	dist  int
}

func (c coordinate) distance(dest coordinate) int {
	return abs(c.X-dest.X) + abs(c.Y-dest.Y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func lineIntersection(a, b, c, d coordinate) coordinate {
	a1 := b.Y - a.Y
	b1 := a.X - b.X
	c1 := a1*a.X + b1*a.Y

	a2 := d.Y - c.Y
	b2 := c.X - d.X
	c2 := a2*c.X + b2*c.Y

	det := a1*b2 - a2*b1
	x := (b2*c1 - b1*c2) / det
	y := (a1*c2 - a2*c1) / det

	return coordinate{x, y}
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

	sensors := []sensor{}

	for fileScanner.Scan() {
		var sensorCoord, beaconCoord coordinate
		_, err := fmt.Sscanf(fileScanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorCoord.Y, &sensorCoord.X, &beaconCoord.Y, &beaconCoord.X)
		if err != nil {
			return 0, errors.New("Cannot parse line")
		}

		distance := sensorCoord.distance(beaconCoord)
		sensors = append(sensors, sensor{sensorCoord, distance})
	}

	var a, b, c, d coordinate
	a = coordinate{0, 0}

	for i := 0; i < len(sensors); i++ {
		for j := i + 1; j < len(sensors); j++ {
			distance := sensors[i].coord.distance(sensors[j].coord)
			if distance == sensors[i].dist+sensors[j].dist+2 {
				if a.X == 0 {
					a = coordinate{
						sensors[i].coord.X + sensors[i].dist + 1,
						sensors[i].coord.Y,
					}
					b = coordinate{
						sensors[j].coord.X - sensors[j].dist - 1,
						sensors[j].coord.Y,
					}
				} else {
					c = coordinate{
						sensors[i].coord.X - sensors[i].dist - 1,
						sensors[i].coord.Y,
					}
					d = coordinate{
						sensors[j].coord.X + sensors[j].dist + 1,
						sensors[j].coord.Y,
					}
				}

			}

		}
	}

	intersection := lineIntersection(a, b, c, d)
	return intersection.Y*4000000 + intersection.X, nil
}
