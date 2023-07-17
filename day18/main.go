package main

import "github.com/kitanovicd/Advent-Of-Code/Advent-Of-Code/day18/first"

func main() {
	result, err := first.Solve()
	if err != nil {
		panic(err)
	}

	println(result)
}
