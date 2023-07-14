package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	area := 0
	isPresent := map[string]bool{}

	for fileScanner.Scan() {
		var x, y, z int
		line := fileScanner.Text()
		count, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if err != nil || count != 3 {
			panic("Cannot parse line")
		}

		if isPresent[line] {
			fmt.Println(fileScanner.Text())
			fmt.Println(line)
			continue
		}

		isPresent[line] = true

		borderCound := 0
		key1 := strconv.Itoa(x-1) + "," + strconv.Itoa(y) + "," + strconv.Itoa(z)
		if isPresent[key1] {
			borderCound++
		}

		key2 := strconv.Itoa(x+1) + "," + strconv.Itoa(y) + "," + strconv.Itoa(z)
		if isPresent[key2] {
			borderCound++
		}

		key3 := strconv.Itoa(x) + "," + strconv.Itoa(y-1) + "," + strconv.Itoa(z)
		if isPresent[key3] {
			borderCound++
		}

		key4 := strconv.Itoa(x) + "," + strconv.Itoa(y+1) + "," + strconv.Itoa(z)
		if isPresent[key4] {
			borderCound++
		}

		key5 := strconv.Itoa(x) + "," + strconv.Itoa(y) + "," + strconv.Itoa(z-1)
		if isPresent[key5] {
			borderCound++
		}

		key6 := strconv.Itoa(x) + "," + strconv.Itoa(y) + "," + strconv.Itoa(z+1)
		if isPresent[key6] {
			borderCound++
		}

		area += 6 - 2*borderCound

	}

	fmt.Println(area)
}
