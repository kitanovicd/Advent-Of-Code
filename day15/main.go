package main

import (
	"fmt"

	"github.com/kitanovicd/Advent-Of-Code/Advent-Of-Code/day15/second"
)

func main() {
	result, err := second.Solve()
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
