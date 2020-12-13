package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var earliest int
	var buslines []int

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	earliest, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	for _, s := range strings.Split(scanner.Text(), ",") {
		if s != "x" {
			line, _ := strconv.Atoi(s)
			buslines = append(buslines, line)
		} else {
			buslines = append(buslines, -1)
		}
	}

	favLine := 0
	waitTime := math.MaxInt64
	for _, l := range buslines {
		if l != -1 {
			if t := l - earliest%l; t < waitTime {
				favLine = l
				waitTime = t
			}
		}
	}

	fmt.Println("Part 1:", favLine*waitTime)

	t, step := 0, 1
	for i, l := range buslines {
		if l == -1 {
			continue
		}
		for (t+i)%l != 0 {
			t += step
		}
		step *= l
	}

	fmt.Println("Part 2:", t)
}
