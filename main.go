package main

import (
	"fmt"

	"github.com/ericthomasca/adventofcode2024/day01"
	"github.com/ericthomasca/adventofcode2024/day02"
	"github.com/ericthomasca/adventofcode2024/day03"
	"github.com/ericthomasca/adventofcode2024/day04"
)

func main() {
	fmt.Println("Advent of Code 2024 Solutions")
	fmt.Println("-----------------------------")
	fmt.Println() 

	// Day 1
	fmt.Println("Day 1")
	fmt.Println("Part 1:", day01.Part1())
	fmt.Println("Part 2:", day01.Part2())
	fmt.Println()
	
	// Day 2
	fmt.Println("Day 2")
	fmt.Println("Part 1:", day02.Part1())
	fmt.Println("Part 2:", day02.Part2())
	fmt.Println()

	// Day 3
	fmt.Println("Day 3")
	fmt.Println("Part 1:", day03.Part1())
	fmt.Println("Part 2:", day03.Part2())
	fmt.Println()

	// Day 4
	fmt.Println("Day 4")
	fmt.Println("Part 1:", day04.Part1())
	fmt.Println("Part 2:", day04.Part2())
	fmt.Println()
}
