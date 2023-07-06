package main

import (
	"bufio"
	"fmt"
	"os"
)

func solveFirst() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		panic("Cannot open file")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	lists := InitCrates(fileScanner)

	for fileScanner.Scan() {
		count, from, to, err := GetCountFromAndTo(fileScanner.Text())
		if err != nil {
			panic("Cannot parse line")
		}

		for i := 0; i < int(count); i++ {
			e := lists[from].Back()
			lists[from].Remove(e)
			lists[to].PushBack(e.Value.(byte))
		}
	}

	result := ""
	for i := 0; i < 9; i++ {
		e := lists[i].Back()
		if e != nil {
			result += string(lists[i].Back().Value.(byte))
		}
	}

	fmt.Println(result)
}

func main() {
	solveFirst()
}
