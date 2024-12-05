package day04

import (
	"bufio"
	"fmt"
	"os"
)

func Part1() int {
	xmasMatrix := createStringMatrix("day04/input.txt")

	var xmasCount int

	for i := range xmasMatrix {
		for j := range xmasMatrix[i] {
			// Stay in bounds on the rows
			if j+3 < len(xmasMatrix[i]) {
				leftToRight := string(xmasMatrix[i][j]) + string(xmasMatrix[i][j+1]) + string(xmasMatrix[i][j+2]) + string(xmasMatrix[i][j+3])
				rightToLeft := string(xmasMatrix[i][j+3]) + string(xmasMatrix[i][j+2]) + string(xmasMatrix[i][j+1]) + string(xmasMatrix[i][j])
				
				// Check for XMAS on rows reading left to right or right to left
				if leftToRight == "XMAS" || rightToLeft == "XMAS" {
					xmasCount++
				}

			}
			// Stay in bounds on the columns
			if i+3 < len(xmasMatrix) {
				topToBottom := string(xmasMatrix[i][j]) + string(xmasMatrix[i+1][j]) + string(xmasMatrix[i+2][j]) + string(xmasMatrix[i+3][j])
				bottomToTop := string(xmasMatrix[i+3][j]) + string(xmasMatrix[i+2][j]) + string(xmasMatrix[i+1][j]) + string(xmasMatrix[i][j])
				
				// Check for XMAS in columns reading top to bottom and bottom to top
				if topToBottom == "XMAS" || bottomToTop == "XMAS" {
					xmasCount++
				}
			}

			// Stay in bounds for diagonals top left to bottom right and reverse
			if i+3 < len(xmasMatrix) && j+3 < len(xmasMatrix[i]) {
				topLeftToBottomRight := string(xmasMatrix[i][j]) + string(xmasMatrix[i+1][j+1]) + string(xmasMatrix[i+2][j+2]) + string(xmasMatrix[i+3][j+3])
				bottomRightToTopLeft := string(xmasMatrix[i+3][j+3]) + string(xmasMatrix[i+2][j+2]) + string(xmasMatrix[i+1][j+1]) + string(xmasMatrix[i][j])
				
				// Check for XMAS in columns reading along diagonals
				if topLeftToBottomRight == "XMAS" || bottomRightToTopLeft == "XMAS" {
					xmasCount++
				}
			}

			// Stay in bounds for diagonals the other direction
			if i+3 < len(xmasMatrix) && j >= 3 {
				topRightToBottomLeft := string(xmasMatrix[i][j]) + string(xmasMatrix[i+1][j-1]) + string(xmasMatrix[i+2][j-2]) + string(xmasMatrix[i+3][j-3])
				bottomLeftToTopRight := string(xmasMatrix[i+3][j-3]) + string(xmasMatrix[i+2][j-2]) + string(xmasMatrix[i+1][j-1]) + string(xmasMatrix[i][j])

				// Check for XMAS on diagonals top-right to bottom-left and reverse
				if topRightToBottomLeft == "XMAS" || bottomLeftToTopRight == "XMAS" {
					xmasCount++
				}
			}
		}
	}

	return xmasCount
}

func Part2() int {
	xmasMatrix := createStringMatrix("day04/input.txt")

	var xmasCount int

	for i := range xmasMatrix {
		for j := range xmasMatrix[i] {
			// Check if there's enough space for the "X-MAS" pattern
			if i-1 >= 0 && i+1 < len(xmasMatrix) && j-1 >= 0 && j+1 < len(xmasMatrix[i]) {
				topLeftToBottomRight := string(xmasMatrix[i-1][j-1]) + string(xmasMatrix[i][j]) + string(xmasMatrix[i+1][j+1])
				bottomLeftToTopRight := string(xmasMatrix[i+1][j-1]) + string(xmasMatrix[i][j]) + string(xmasMatrix[i-1][j+1])
	
				// Check for XMAS pattern
				if (topLeftToBottomRight == "MAS" || topLeftToBottomRight == "SAM") &&
					(bottomLeftToTopRight == "MAS" || bottomLeftToTopRight == "SAM") {
					xmasCount++
				}
			}
		}
	}

	return xmasCount
}

func createStringMatrix(filename string) []string {
	// Import file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var xmasMatrix []string

	// Put each line into an array
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		xmasMatrix = append(xmasMatrix, line)
	}

	return xmasMatrix
}
