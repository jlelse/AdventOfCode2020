package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	maxSeat := 0
	seats := [127*8 + 7]bool{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seat := seat(scanner.Text())
		if seat > maxSeat {
			maxSeat = seat
		}
		seats[seat] = true
	}

	fmt.Println("Part 1:", maxSeat)

outerloop:
	for i := 1; i < 127; i++ {
		for j := 0; j < 8; j++ {
			if seat := seatID(i, j); !seats[seat] && seats[seat-1] && seats[seat+1] {
				fmt.Println("Part 2:", seat)
				break outerloop
			}
		}
	}
}

func seat(input string) int {
	row := partition(input, 0, 127, 'F', 'B')
	column := partition(input, 0, 7, 'L', 'R')
	return seatID(row, column)
}

func seatID(row, col int) int {
	return row*8 + col
}

func partition(input string, min, max int, lower, upper rune) int {
	for _, c := range input {
		if c == lower {
			max -= (max - min + 1) / 2
		} else if c == upper {
			min += (max - min + 1) / 2
		}
	}
	return min
}
