package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	treeMap := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		treeRow := []int{}
		for _, char := range scanner.Text() {
			if char == '#' {
				treeRow = append(treeRow, 1)
			} else if char == '.' {
				treeRow = append(treeRow, 0)
			}
		}
		treeMap = append(treeMap, treeRow)
	}

	rowLength := len(treeMap[0])

	row := 0
	pos := 0
	trees := 0

	// Part 1

	for row < len(treeMap) {
		trees += treeMap[row][pos]
		row++
		pos += 3
		pos %= rowLength
	}

	fmt.Println("Part 1:", trees)

	// Part 2

	trees = 1

	type slope struct {
		right, down int
	}

	for _, s := range []slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}} {
		slopeTrees := 0
		row = 0
		pos = 0
		for row < len(treeMap) {
			slopeTrees += treeMap[row][pos]
			row += s.down
			pos += s.right
			pos %= rowLength
		}
		trees *= slopeTrees
	}

	fmt.Println("Part 2:", trees)
}
