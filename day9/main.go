package main

import "github.com/kitanovicd/Advent-Of-Code/Advent-Of-Code/day9/second"

func main() {
	result, err := second.Solve()
	if err != nil {
		panic(err)
	}

	println(result)
}
