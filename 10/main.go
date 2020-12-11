package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Part 1

	joltages := []int{}

	file, _ := os.Open("input.txt")
	defer file.Close()
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		joltage, _ := strconv.Atoi(scanner.Text())
		joltages = append(joltages, joltage)
	}

	sort.Ints(joltages)

	jolts := map[int]int{}
	currentJolt := 0

	for _, j := range joltages {
		jolts[j-currentJolt]++
		currentJolt = j
	}

	// Device itself
	jolts[3]++

	fmt.Println("Part 1:", jolts[1]*jolts[3])

	// Part 2

	fmt.Println("Part 2:", countPaths(0, joltages, map[int]int{}))
}

func countPaths(current int, joltages []int, memory map[int]int) (total int) {
	if len(joltages) == 0 {
		return 1
	}
	for _, next := range joltages {
		if next-current > 3 {
			break
		}
		if _, exists := memory[next]; !exists {
			memory[next] = countPaths(next, joltages[1:], memory)
		}
		total += memory[next]
	}
	return
}
