package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var WORRY_DIVISION int = 1
var NUMBER_OF_ROUNDS int = 10000

type Monkey struct {
	Items               []int
	Operation           byte
	Operator            int
	Divisor             int
	ThrowToDivisible    int
	ThrowToNotDivisible int
	CountInspected      int
}

func findItems(line string) ([]int, error) {
	itemsStr := strings.Split(line[18:], ", ")
	items := []int{}

	for _, itemStr := range itemsStr {
		item, err := strconv.Atoi(itemStr)
		if err != nil {
			return nil, errors.New("Failed to convert number")
		}
		items = append(items, int(item))
	}

	return items, nil
}

func findOperationAndOperator(line string) (byte, int) {
	operation := line[23]
	operator, err := strconv.Atoi(line[25:])
	if err != nil {
		return operation, 0
	}
	return operation, int(operator)
}

func findDivisor(line string) (int, error) {
	divisor, err := strconv.Atoi(line[21:])
	if err != nil {
		return 0, errors.New("Failed to convert divisor")
	}
	return int(divisor), nil
}

func findThrowToDivisible(line string) (int, error) {
	throwToDivisible, err := strconv.Atoi(line[29:])
	if err != nil {
		return -1, errors.New("Failed to convert index of monkey to throw to if divisible")
	}
	return throwToDivisible, nil
}

func findThrowToNotDivisible(line string) (int, error) {
	throwToNotDivisible, err := strconv.Atoi(line[30:])
	if err != nil {
		return -1, errors.New("Failed to convert index of monkey to throw to if not divisible")
	}
	return throwToNotDivisible, nil
}

func parseBlock(file *os.File, fileScanner *bufio.Scanner) (Monkey, error) {
	fileScanner.Scan()

	items, err := findItems(fileScanner.Text())
	if err != nil {
		return Monkey{}, err
	}

	fileScanner.Scan()
	operation, operator := findOperationAndOperator(fileScanner.Text())

	fileScanner.Scan()
	divisor, err := findDivisor(fileScanner.Text())

	fileScanner.Scan()
	throwToDivisible, err := findThrowToDivisible(fileScanner.Text())
	if err != nil {
		return Monkey{}, err
	}

	fileScanner.Scan()
	throwToNotDivisible, err := findThrowToNotDivisible(fileScanner.Text())
	if err != nil {
		return Monkey{}, err
	}

	fileScanner.Scan()
	return Monkey{
		Items:               items,
		Operation:           operation,
		Operator:            operator,
		Divisor:             divisor,
		ThrowToDivisible:    throwToDivisible,
		ThrowToNotDivisible: throwToNotDivisible,
		CountInspected:      0,
	}, nil
}

func executeOperation(operation byte, operator1 int, operator2 int, modBy int) (int, error) {
	if operator2 == 0 {
		operator2 = operator1
	}

	var result int
	switch operation {
	case '+':
		result = operator1 + operator2
	case '-':
		result = operator1 - operator2
	case '*':
		result = operator1 * operator2
	case '/':
		result = operator1 / operator2
	default:
		return 0, errors.New("Invalid operation")
	}

	return (result / WORRY_DIVISION) % modBy, nil
}

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		panic("Cannot open file")
	}
	defer file.Close()

	monkeys := []Monkey{}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	modBy := 1
	for fileScanner.Scan() {
		monkey, err := parseBlock(file, fileScanner)

		if err != nil {
			panic(err)
		}

		monkeys = append(monkeys, monkey)
		modBy *= monkey.Divisor
	}

	for step := 1; step <= NUMBER_OF_ROUNDS; step++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.Items {
				result, err := executeOperation(monkey.Operation, item, monkey.Operator, modBy)
				if err != nil {
					panic(err)
				}

				var throwTo int
				if (result % monkey.Divisor) == 0 {
					throwTo = monkey.ThrowToDivisible
				} else {
					throwTo = monkey.ThrowToNotDivisible
				}

				monkeys[throwTo].Items = append(monkeys[throwTo].Items, result)
			}

			monkeys[i].CountInspected += len(monkeys[i].Items)
			monkeys[i].Items = []int{}

		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].CountInspected > monkeys[j].CountInspected
	})

	fmt.Println("Result is:", monkeys[0].CountInspected*monkeys[1].CountInspected)
}
