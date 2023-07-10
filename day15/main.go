package main

import (
	"fmt"

	"github.com/kitanovicd/Advent-Of-Code/Advent-Of-Code/day15/first"
)

func main() {
	result, err := first.Solve()
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
