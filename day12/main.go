package main

import (
	"bufio"
	"os"
)

type Coordinate struct {
	x int
	y int
}

func canIMoveUp(current Coordinate, table [][]byte, visited map[Coordinate]bool) (Coordinate, bool) {
	nextPos := Coordinate{current.x - 1, current.y}
	if current.x == 0 || visited[nextPos] {
		return Coordinate{}, false
	}

	if int(table[current.x][current.y])-int(table[current.x-1][current.y]) >= -1 {
		return nextPos, true
	}

	return Coordinate{}, false
}

func canIMoveDown(current Coordinate, table [][]byte, visited map[Coordinate]bool) (Coordinate, bool) {
	nextPos := Coordinate{current.x + 1, current.y}
	if current.x == len(table)-1 || visited[nextPos] {
		return Coordinate{}, false
	}

	if int(table[current.x][current.y])-int(table[current.x+1][current.y]) >= -1 {
		return nextPos, true
	}

	return Coordinate{}, false
}

func canIMoveLeft(current Coordinate, table [][]byte, visited map[Coordinate]bool) (Coordinate, bool) {
	nextPos := Coordinate{current.x, current.y - 1}
	if current.y == 0 || visited[nextPos] {
		return Coordinate{}, false
	}

	if int(table[current.x][current.y])-int(table[current.x][current.y-1]) >= -1 {
		return nextPos, true
	}

	return Coordinate{}, false
}

func canIMoveRight(current Coordinate, table [][]byte, visited map[Coordinate]bool) (Coordinate, bool) {
	nextPos := Coordinate{current.x, current.y + 1}
	if current.y == len(table[0])-1 || visited[nextPos] {
		return Coordinate{}, false
	}

	if int(table[current.x][current.y])-int(table[current.x][current.y+1]) >= -1 {
		return nextPos, true
	}

	return Coordinate{}, false
}

func parseFile(file *os.File) (table [][]byte, beginings []Coordinate, destination Coordinate) {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanRunes)

	table = [][]byte{{}}
	beginings = []Coordinate{}
	for fileScanner.Scan() {
		char := fileScanner.Text()[0]
		switch char {
		case '\n':
			table = append(table, []byte{})
			continue
		case 'a':
			beginings = append(beginings, Coordinate{len(table) - 1, len(table[len(table)-1])})
		case 'S':
			char = 'a'
			beginings = append(beginings, Coordinate{len(table) - 1, len(table[len(table)-1])})
		case 'E':
			char = 'z'
			destination = Coordinate{len(table) - 1, len(table[len(table)-1])}
		}

		table[len(table)-1] = append(table[len(table)-1], char)
	}

	return table, beginings, destination
}

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		panic("Cannot open file")
	}
	defer file.Close()

	table, beginings, destination := parseFile(file)

	type Elem struct {
		Coordinate Coordinate
		Steps      int
	}
	queue := []Elem{}
	steps := 0
	visited := map[Coordinate]bool{}
	for _, begining := range beginings {
		queue = append(queue, Elem{begining, steps})
		visited[begining] = true
	}

	curr := queue[0].Coordinate
	steps = queue[0].Steps

	for curr.x != destination.x || curr.y != destination.y {
		nextPos, canIMove := canIMoveUp(curr, table, visited)
		if canIMove {
			queue = append(queue, Elem{nextPos, steps + 1})
			visited[nextPos] = true
		}

		nextPos, canIMove = canIMoveDown(curr, table, visited)
		if canIMove {
			queue = append(queue, Elem{nextPos, steps + 1})
			visited[nextPos] = true
		}

		nextPos, canIMove = canIMoveLeft(curr, table, visited)
		if canIMove {
			queue = append(queue, Elem{nextPos, steps + 1})
			visited[nextPos] = true
		}

		nextPos, canIMove = canIMoveRight(curr, table, visited)
		if canIMove {
			queue = append(queue, Elem{nextPos, steps + 1})
			visited[nextPos] = true
		}

		e := queue[0]
		queue = queue[1:]

		curr = e.Coordinate
		steps = e.Steps
		visited[curr] = true
	}

	println(steps)
}
