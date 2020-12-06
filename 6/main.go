package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type void struct{}
	yes := void{}
	group := map[rune]void{}
	count := 0

	newGroup := func() {
		count += len(group)
		group = map[rune]void{}
	}

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			// Empty row, new group
			newGroup()
		} else {
			for _, r := range strings.ReplaceAll(text, " ", "") {
				group[r] = yes
			}
		}
	}
	// Last group
	newGroup()

	fmt.Println("Part 1:", count)

	// Part 2
	count = 0
	isNewGroup := true

	file2, _ := os.Open("input.txt")
	defer file2.Close()
	scanner = bufio.NewScanner(file2)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			// Empty row, new group
			newGroup()
			isNewGroup = true
		} else {
			if isNewGroup {
				// First person
				for _, r := range strings.ReplaceAll(text, " ", "") {
					group[r] = yes
				}
				isNewGroup = false
			} else {
				for r := range group {
					if !strings.ContainsRune(text, r) {
						delete(group, r)
					}
				}
			}
		}
	}
	newGroup()

	fmt.Println("Part 2:", count)
}
