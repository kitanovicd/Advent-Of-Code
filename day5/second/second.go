package second

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func initCrates(fileScanner *bufio.Scanner) [9][]byte {
	var crates [9][]byte

	for fileScanner.Scan() && fileScanner.Text()[1] != '1' {
		line := fileScanner.Text()

		for i := 1; i < len(line); i += 4 {
			if line[i] != ' ' {
				index := i / 4
				crates[index] = append([]byte(line[i:i+1]), crates[index]...)
			}
		}
	}

	fileScanner.Scan()
	return crates
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

	crates := initCrates(fileScanner)

	for fileScanner.Scan() {
		count, from, to, err := parseLine(fileScanner.Text())

		if err != nil {
			panic("Cannot parse line")
		}

		crates[to] = append(crates[to], crates[from][len(crates[from])-count:]...)
		crates[from] = crates[from][:len(crates[from])-count]
	}

	result := ""
	for _, crate := range crates {
		if len(crate) == 0 {
			continue
		}

		result += string(crate[len(crate)-1])
	}

	fmt.Println(result)

}
