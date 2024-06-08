package main

import (
	"fmt"

	"github.com/ayushgpt01/advent-of-code-golang/input"
	"github.com/ayushgpt01/advent-of-code-golang/pkg/day1"
)

func main() {
	input, err := input.Read("day1.txt")
	if err != nil {
		panic(err)
	}
	result := day1.Solve(input)
	fmt.Println(result)

	result2 := day1.SolvePart2(input)
	fmt.Println(result2)
}
