package day01

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func Part1() int {
	list1, list2 := createLists()

	// Sort Lists
	slices.Sort(list1)
	slices.Sort(list2)

	// Create array of differences in
	// corresponding items of both lists
	var differenceList []int
	for i := range len(list1) {
		difference := list1[i] - list2[i]
		if difference < 0 {
			difference = difference / -1
		}
		differenceList = append(differenceList, difference)
	}

	// Add differences to total
	var differenceTotal int
	for i := range differenceList {
		differenceTotal += differenceList[i]
	}

	return differenceTotal

}

func Part2() int {
	list1, list2 := createLists()
	var similarityScore int

	// For every item in list 1
	for i := range len(list1) {
		// If the item for list 1 is in list 2
		if slices.Contains(list2, list1[i]) {
			// Look through list 2 and count how many times the
			// item from list 1 is in list 2
			var matches int
			for _, num := range list2 {
				if list1[i] == num {
					matches++
				}
			}

			// Multiply the item from list 1 to the match
			// number and add to total
			itemSimilarity := list1[i] * matches
			similarityScore += itemSimilarity

		}
	}

	return similarityScore
}

// Create 2 lists from the input file
// Each line is 5 digits, 3 spaces, then 5 digits
// representing list1[i] and list2[i]
func createLists() ([]int, []int) {
	// Import file
	file, err := os.Open("day01/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Instantiate lists
	var list1 []int
	var list2 []int

	// Go line by line on the input file
	// Add left ints to list1
	// Add right ints to list2
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		unsortedListLine := scanner.Text()

		// First 5 characters is first ID
		firstUnsortedIdString := unsortedListLine[:5]
		firstUnsortedId, _ := strconv.Atoi(firstUnsortedIdString)

		// Characters 8 to end(13) is second ID
		secondUnsortedIdString := unsortedListLine[8:]
		secondUnsortedId, _ := strconv.Atoi(secondUnsortedIdString)

		// Add IDs to lists
		list1 = append(list1, firstUnsortedId)
		list2 = append(list2, secondUnsortedId)
	}

	return list1, list2
}
