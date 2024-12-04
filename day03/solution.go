package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part1() int {
	// Import file
	file, err := os.Open("day03/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var grandTotal int
	var toMultiplySlice []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		r, _ := regexp.Compile(`mul\(\d+,\d+\)`)
		toMultiplySlice = append(toMultiplySlice, r.FindAllString(line, -1)...)
	}

	var cleanedToMultiplySlice []string

	for i := range toMultiplySlice {
		toMultiply := toMultiplySlice[i]
		toMultiply = strings.ReplaceAll(toMultiply, "mul(", "")
		toMultiply = strings.ReplaceAll(toMultiply, ")", "")
		cleanedToMultiplySlice = append(cleanedToMultiplySlice, toMultiply)
	}

	for i := range cleanedToMultiplySlice {
		nums := strings.Split(cleanedToMultiplySlice[i], ",")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		multipliedNum := num1 * num2
		grandTotal += multipliedNum
	}

	return grandTotal
}

func Part2() int {
	// Import file
	file, err := os.Open("day03/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var toMultiplySlice []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		r, _ := regexp.Compile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
		toMultiplySlice = append(toMultiplySlice, r.FindAllString(line, -1)...)
	}

	var parsedToMultiplySlice []string
	addMuls := true

	for _, mult := range toMultiplySlice {
		if mult == "don't()" {
			addMuls = false
		} else if mult == "do()" {
			addMuls = true
		} else if addMuls && len(mult) >= 3 && mult[:3] == "mul" {
			parsedToMultiplySlice = append(parsedToMultiplySlice, mult)
		}
	}

	var grandTotal int
	var cleanedToMultiplySlice []string

	for i := range parsedToMultiplySlice {
		toMultiply := parsedToMultiplySlice[i]
		toMultiply = strings.ReplaceAll(toMultiply, "mul(", "")
		toMultiply = strings.ReplaceAll(toMultiply, ")", "")
		cleanedToMultiplySlice = append(cleanedToMultiplySlice, toMultiply)
	}

	for i := range cleanedToMultiplySlice {
		nums := strings.Split(cleanedToMultiplySlice[i], ",")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		multipliedNum := num1 * num2
		grandTotal += multipliedNum
	}

	return grandTotal
}
