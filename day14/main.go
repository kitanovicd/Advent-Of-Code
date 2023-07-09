package main

import "github.com/kitanovicd/Advent-Of-Code/Advent-Of-Code/day14/first"

func main() {
	result, err := first.Solve()
	if err != nil {
		panic(err)
	}

	println(result)
}
