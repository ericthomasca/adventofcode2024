package day02

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1() int {
	// Import file
	file, err := os.Open("day02/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var safeReportCount int
	var safeReportDroppedCount int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Each line is a rawReport
		rawReport := scanner.Text()

		// Convert report string into array of levels
		// Numbers are strings at this point
		reportStringSlice := strings.Fields(rawReport)

		// Convert from []string to []int
		var report []int
		for _, i := range reportStringSlice {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			report = append(report, j)
		}

		// If report pass safety checks
		// increment safe report count
		if isReportSafe(report) {
			safeReportCount++
		}

		if isDroppedReportSafe(report) {
			safeReportDroppedCount++
		}
	}

	return safeReportCount
}

func isReportSafe(report []int) bool {
	// Create ascending report and compare to
	// original and save if same as bool
	// Checks if in ascending order
	var ascendingReport []int
	ascendingReport = append(ascendingReport, report...)
	slices.Sort(ascendingReport)
	isAscending := slices.Equal(report, ascendingReport)

	// Create descending report and compare to
	// original and save if same as bool
	// Checks if in descending order
	var descendingReport []int
	descendingReport = append(descendingReport, report...)
	slices.Sort(descendingReport)
	slices.Reverse(descendingReport)
	isDescending := slices.Equal(report, descendingReport)

	// If not ascending or descending, return false
	if !isAscending && !isDescending {
		return false
	}

	// Check if differences between adjacent levels are within the range 1-3
	for i := 0; i < len(report)-1; i++ {
		difference := math.Abs(float64(report[i]) - float64(report[i+1]))
		if difference < 1 || difference > 3 {
			return false
		}
	}

	return true
}

func isDroppedReportSafe(report []int) bool {
	// If the report is already safe, return true
	if isReportSafe(report) {
		return true
	}

	// Check if removing one element (from any position) makes the report safe
	for i := 0; i < len(report); i++ {
		// Create a modified report by removing the element at index `i`
		modifiedReport := append(report[:i], report[i+1:]...)
		if isReportSafe(modifiedReport) {
			return true
		}
	}

	// If it's still not safe, return false
	return false
}

func Part2() string {
	return "Part 2 NOT FOUND"
}
