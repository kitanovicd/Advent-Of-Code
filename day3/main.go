package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func calculatePriority(char byte) int {
	if unicode.IsLower(rune(char)) {
		return int(char-'a') + 1
	}
	return int(char-'A') + 27
}

func findCommonElements(str1, str2 string) string {
	isPresentInStr1 := make(map[rune]bool)
	for _, char := range str1 {
		isPresentInStr1[char] = true
	}

	common := ""
	for _, char := range str2 {
		if isPresentInStr1[char] {
			common += string(char)
		}
	}

	return common
}

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		panic("Cannot open file")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		str1 := fileScanner.Text()
		fileScanner.Scan()
		str2 := fileScanner.Text()
		fileScanner.Scan()
		str3 := fileScanner.Text()

		common := findCommonElements(str1, str2)
		common = findCommonElements(common, str3)
		if len(common) == 0 {
			panic("No common elements found")
		}

		sum += calculatePriority(common[0])
	}

	fmt.Println("Sum is", sum)
}
