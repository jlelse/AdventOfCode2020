package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	numbers := []int{}

	file, _ := os.Open("input.txt")
	defer file.Close()
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}

	preambleLength := 25
	preambleStart := 0
	invalidNum := 0
	for i, number := range numbers {
		if i >= preambleLength {
			valid := validNumber(preamble(numbers, preambleStart, preambleLength), number)
			if !valid {
				invalidNum = number
				break
			}
			preambleStart++
		}
	}

	fmt.Println("Part 1:", invalidNum)

	finalSet := []int{}
outer:
	for i := 0; i < len(numbers)-1; i++ {
		for j := len(numbers) - 1; j > i; j-- {
			if testSet := numbers[i:j]; sum(testSet) == invalidNum {
				finalSet = testSet
				break outer
			}
		}
	}
	sort.Ints(finalSet)
	weakness := finalSet[0] + finalSet[len(finalSet)-1]

	fmt.Println("Part 2:", weakness)

}

func preamble(n []int, start, length int) []int {
	return n[start : start+length]
}

func validNumber(preamble []int, num int) bool {
	for _, i := range preamble {
		for _, j := range preamble {
			if i+j == num && i != j {
				return true
			}
		}
	}
	return false
}

func sum(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}
