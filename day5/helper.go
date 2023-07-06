package main

import (
	"bufio"
	"container/list"
	"regexp"
	"strconv"
)

func InitCrates(fileScanner *bufio.Scanner) [9]*list.List {
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

func GetCountFromAndTo(line string) (int, int, int, error) {
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
