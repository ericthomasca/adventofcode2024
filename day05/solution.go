package day05

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1() int {
	correctOrders, err := getCorrectOrders("day05/input.txt")
	if err != nil {
		panic(err)
	}

	var middleNumberTotal int

	for _, correctOrder := range correctOrders {
		middleNumberIndex := (len(correctOrder) - 1) / 2
		middleNumber := correctOrder[middleNumberIndex]
		middleNumberTotal += middleNumber
	}

	return middleNumberTotal
}

func Part2() int {
	pageOrderingRules, pageOrdersToCheck, err := splitInput("day05/input.txt")
	if err != nil {
		panic(err)
	}

	var middleNumberTotal int

	for _, order := range pageOrdersToCheck {
		if isOrderCorrect(order, pageOrderingRules) {
			continue
		}

		// Correct the order
		correctedOrder := reorderPages(order, pageOrderingRules)

		// Find the middle number
		middleNumberIndex := (len(correctedOrder) - 1) / 2
		middleNumber := correctedOrder[middleNumberIndex]
		middleNumberTotal += middleNumber
	}

	return middleNumberTotal
}

func getCorrectOrders(filename string) ([][]int, error) {
	pageOrderingRules, pageOrdersToCheck, err := splitInput(filename)
	if err != nil {
		return nil, err
	}

	var correctOrders [][]int

	for _, order := range pageOrdersToCheck {
		passingOrder := true
		for _, rules := range pageOrderingRules {
			beforeRule := rules[0]
			afterRule := rules[1]

			beforeRuleLocation := slices.Index(order, beforeRule)
			afterRuleLocation := slices.Index(order, afterRule)

			// Skip this rule if either location is -1
			if beforeRuleLocation == -1 || afterRuleLocation == -1 {
				continue
			}

			// If the rule is violated, mark the order as failing
			if beforeRuleLocation > afterRuleLocation {
				passingOrder = false
				break
			}
		}

		// Add to correctOrders if it passed all applicable rules
		if passingOrder {
			correctOrders = append(correctOrders, order)
		}
	}

	return correctOrders, nil
}

func splitInput(filename string) ([][]int, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var pageOrderingRules [][]int
	var pageOrdersToCheck [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		if strings.Contains(line, "|") {
			pageOrderingRules, err = stringToSlices(line, pageOrderingRules, "|")
			if err != nil {
				return nil, nil, err
			}
		} else if strings.Contains(line, ",") {
			pageOrdersToCheck, err = stringToSlices(line, pageOrdersToCheck, ",")
			if err != nil {
				return nil, nil, err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}

	return pageOrderingRules, pageOrdersToCheck, nil
}

func isOrderCorrect(order []int, rules [][]int) bool {
	for _, rule := range rules {
		beforeRule := rule[0]
		afterRule := rule[1]

		beforeRuleLocation := slices.Index(order, beforeRule)
		afterRuleLocation := slices.Index(order, afterRule)

		// Skip this rule if either location is -1
		if beforeRuleLocation == -1 || afterRuleLocation == -1 {
			continue
		}

		// If the rule is violated, the order is incorrect
		if beforeRuleLocation > afterRuleLocation {
			return false
		}
	}
	return true
}

func reorderPages(order []int, rules [][]int) []int {
	// Start with the original order
	remainingPages := make([]int, len(order))
	copy(remainingPages, order)
	sortedOrder := []int{}

	for len(remainingPages) > 0 {
		newRemainingPages := []int{}

		for _, page := range remainingPages {
			isValid := true

			// Check if this page can be placed next in the sorted order
			for _, rule := range rules {
				beforeRule := rule[0]
				afterRule := rule[1]

				// If this page must come after another page still in remainingPages, skip it
				if page == afterRule && slices.Contains(remainingPages, beforeRule) {
					isValid = false
					break
				}
			}

			if isValid {
				// Add the page to the sorted order
				sortedOrder = append(sortedOrder, page)
			} else {
				// Keep the page in the remainingPages for the next iteration
				newRemainingPages = append(newRemainingPages, page)
			}
		}

		// Update remainingPages for the next iteration
		remainingPages = newRemainingPages
	}

	return sortedOrder
}

func stringToSlices(line string, appendingSlice [][]int, splitChar string) ([][]int, error) {
	sliceStrings := strings.Split(line, splitChar)
	var slices []int
	for _, elem := range sliceStrings {
		j, err := strconv.Atoi(strings.TrimSpace(elem))
		if err != nil {
			return nil, fmt.Errorf("Error converting to integer: %v\n", elem)
		}
		slices = append(slices, j)
	}
	appendingSlice = append(appendingSlice, slices)
	return appendingSlice, nil
}
