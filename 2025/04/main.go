package main

import (
	"fmt"
	"strings"

	"github.com/rodrigoaraujo46/aoc/2025/readers"
)

const todaysPath = "./2025/04"

func isAcessible(grid [][]bool, x, y int) bool {
	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	count := 0
	for _, dir := range dirs {
		newY, newX := y+dir[0], x+dir[1]
		if 0 <= newY && newY < len(grid) && 0 <= newX && newX < len(grid[0]) {
			if grid[newY][newX] {
				count++
			}
		}
	}

	return count < 4
}

func main() {
	lines := strings.Split(readers.ReadInput(todaysPath), "\n")

	grid := make([][]bool, len(lines)-1)
	for i := range grid {
		row := make([]bool, len(lines[0]))
		for j, r := range lines[i] {
			if r == '@' {
				row[j] = true
			}
		}
		grid[i] = row
	}

	ans := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] && isAcessible(grid, x, y) {
				grid[y][x] = false
				ans++
				y = 0
				x = 0
			}
		}
	}

	fmt.Println(ans)
}
