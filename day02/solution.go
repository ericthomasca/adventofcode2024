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
	reports := loadReports("day02/input.txt")
	safeReportCount := 0

	for _, report := range reports {
		if isReportSafe(report) {
			safeReportCount++
		}
	}

	return safeReportCount
}

func Part2() int {
	reports := loadReports("day02/input.txt")
	safeDroppedReportCount := 0

	for _, report := range reports {
		if isDroppedReportSafe(report) {
			safeDroppedReportCount++
		}
	}

	return safeDroppedReportCount
}

// loadReports reads the input file and converts each line to a slice of integers.
func loadReports(filepath string) [][]int {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	var reports [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rawReport := scanner.Text()
		reports = append(reports, parseReport(rawReport))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return reports
}

// parseReport converts a space-separated string of numbers into a slice of integers.
func parseReport(rawReport string) []int {
	stringValues := strings.Fields(rawReport)
	var report []int

	for _, strValue := range stringValues {
		value, err := strconv.Atoi(strValue)
		if err != nil {
			panic(err)
		}
		report = append(report, value)
	}

	return report
}

// isReportSafe checks if a report is either fully ascending or descending with adjacent differences in range 1–3.
func isReportSafe(report []int) bool {
	if isOrdered(report, true) || isOrdered(report, false) {
		return hasValidDifferences(report)
	}
	return false
}

// isOrdered checks if the report is sorted in ascending or descending order.
func isOrdered(report []int, ascending bool) bool {
	copied := append([]int{}, report...) // Create a copy of the report

	if ascending {
		slices.Sort(copied)
	} else {
		slices.Sort(copied)
		slices.Reverse(copied)
	}

	return slices.Equal(report, copied)
}

// hasValidDifferences checks if the differences between adjacent levels are within the range 1–3.
func hasValidDifferences(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		difference := math.Abs(float64(report[i]) - float64(report[i+1]))
		if difference < 1 || difference > 3 {
			return false
		}
	}
	return true
}

// isDroppedReportSafe checks if a report can be made safe by removing one level.
func isDroppedReportSafe(report []int) bool {
	for i := range report {
		modifiedReport := append([]int{}, report[:i]...)    
		modifiedReport = append(modifiedReport, report[i+1:]...) 

		if isReportSafe(modifiedReport) {
			return true
		}
	}
	return false
}
