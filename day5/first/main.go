package first

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func initCrates(fileScanner *bufio.Scanner) [9]*list.List {
	var lists [9]*list.List

	for i := 0; i < 9; i++ {
		lists[i] = list.New()
	}

	for fileScanner.Scan() && fileScanner.Text()[1] != '1' {
		line := fileScanner.Text()

		for i := 1; i < len(line); i += 4 {
			if line[i] != ' ' {
				index := i / 4
				lists[index].PushFront(line[i])
			}
		}
	}

	fileScanner.Scan()
	return lists
}

func parseLine(line string) (int, int, int, error) {
	pattern := `move (\d+) from (\d+) to (\d+)`
	reg := regexp.MustCompile(pattern)
	numbers := reg.FindStringSubmatch(line)[1:]

	count, err := strconv.Atoi(numbers[0])
	if err != nil {
		return 0, 0, 0, err
	}

	from, err := strconv.Atoi(numbers[1])
	if err != nil {
		return 0, 0, 0, err
	}

	to, err := strconv.Atoi(numbers[2])
	if err != nil {
		return 0, 0, 0, err
	}

	return count, from - 1, to - 1, nil
}

func Solve() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		panic("Cannot open file")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	lists := initCrates(fileScanner)

	for fileScanner.Scan() {
		count, from, to, err := parseLine(fileScanner.Text())
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
