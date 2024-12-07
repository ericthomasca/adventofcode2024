package day06

import (
	"bufio"
	"os"
)

func Part1() int {
	grid, startX, startY, startDir := parseMap("day06/input.txt")
	visited := simulateGuard(grid, startX, startY, startDir)
	distinctVisited := len(visited)
	return distinctVisited
}

func Part2() int {
	grid, startX, startY, startDir := parseMap("day06/input.txt")
	loopPositions := findLoopPositions(grid, startX, startY, startDir)
	distinctLoopPositions := len(loopPositions)
	return distinctLoopPositions
}

func findLoopPositions(grid [][]rune, startX, startY int, startDir string) [][2]int {
	possiblePositions := [][2]int{}

	// Iterate over all positions in the grid
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			// Skip obstacles and the guard's starting position
			if grid[y][x] != '.' || (x == startX && y == startY) {
				continue
			}

			// Temporarily place an obstruction
			grid[y][x] = '#'

			// Check if this causes a loop
			if causesLoop(grid, startX, startY, startDir) {
				possiblePositions = append(possiblePositions, [2]int{x, y})
			}

			// Remove the obstruction
			grid[y][x] = '.'
		}
	}

	return possiblePositions
}

func causesLoop(grid [][]rune, startX, startY int, startDir string) bool {
	// Direction vectors: [x, y]
	directions := map[string][2]int{
		"^": {0, -1},
		">": {1, 0},
		"v": {0, 1},
		"<": {-1, 0},
	}
	// Right turn mapping
	rightTurn := map[string]string{
		"^": ">",
		">": "v",
		"v": "<",
		"<": "^",
	}

	// Set to track visited states (x, y, direction)
	visited := make(map[[3]int]bool)
	x, y := startX, startY
	dir := startDir

	for {
		state := [3]int{x, y, dirToIndex(dir)}
		if visited[state] {
			// Loop detected
			return true
		}
		visited[state] = true

		// Compute the next position
		dx, dy := directions[dir][0], directions[dir][1]
		nextX, nextY := x+dx, y+dy

		// Check if the guard is out of bounds
		if nextY < 0 || nextY >= len(grid) || nextX < 0 || nextX >= len(grid[0]) {
			break
		}

		// Check if there is an obstacle
		if grid[nextY][nextX] == '#' {
			// Turn right
			dir = rightTurn[dir]
		} else {
			// Move forward
			x, y = nextX, nextY
		}
	}

	// No loop detected
	return false
}

func dirToIndex(dir string) int {
	switch dir {
	case "^":
		return 0
	case ">":
		return 1
	case "v":
		return 2
	case "<":
		return 3
	}
	return -1
}

func simulateGuard(grid [][]rune, startX, startY int, startDir string) map[[2]int]bool {
	// Direction vectors: [x, y]
	directions := map[string][2]int{
		"^": {0, -1},
		">": {1, 0},
		"v": {0, 1},
		"<": {-1, 0},
	}
	// Right turn mapping
	rightTurn := map[string]string{
		"^": ">",
		">": "v",
		"v": "<",
		"<": "^",
	}

	visited := make(map[[2]int]bool)
	x, y := startX, startY
	dir := startDir

	for {
		// Mark the current position as visited
		visited[[2]int{x, y}] = true

		// Compute the next position
		dx, dy := directions[dir][0], directions[dir][1]
		nextX, nextY := x+dx, y+dy

		// Check if the guard is out of bounds
		if nextY < 0 || nextY >= len(grid) || nextX < 0 || nextX >= len(grid[0]) {
			break
		}

		// Check if there is an obstacle
		if grid[nextY][nextX] == '#' {
			// Turn right
			dir = rightTurn[dir]
		} else {
			// Move forward
			x, y = nextX, nextY
		}
	}

	return visited
}

func parseMap(filename string) ([][]rune, int, int, string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	var startX, startY int
	var startDir string

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		for x, char := range row {
			if char == '^' || char == 'v' || char == '>' || char == '<' {
				startX, startY = x, y
				startDir = string(char)
				row[x] = '.' // Replace the guard's symbol with open space
			}
		}
		grid = append(grid, row)
		y++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return grid, startX, startY, startDir
}
