package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		panic("Cannot open file")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	valveGraph := map[string][]string{}

	for fileScanner.Scan() {
		var rate int
		var fromValve string
		line := fileScanner.Text()

		_, err := fmt.Sscanf(line, "Valve %s has flow rate=%d;", &fromValve, &rate)
		if err != nil {
			panic("Cannot parse line")
		}

		startIndex := strings.Index(line, "valve") + 7
		endIndex := len(line)
		toValvesConnected := line[startIndex:endIndex]
		toValves := strings.Split(toValvesConnected, ", ")

		valveGraph[fromValve] = toValves
	}

	fmt.Println(valveGraph)
}
