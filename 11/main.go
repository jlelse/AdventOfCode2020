package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	var scannedSeats, currentSeats [][]rune

	file, _ := os.Open("input.txt")
	defer file.Close()
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		row := []rune{}
		for _, r := range scanner.Text() {
			row = append(row, r)
		}
		scannedSeats = append(scannedSeats, row)
	}

	// Part 1

	currentSeats = scannedSeats
rounds:
	for {
		newSeats := [][]rune{}
		for i := range currentSeats {
			row := []rune{}
			for j, r := range currentSeats[i] {
				newRune := r
				if os := occupiedAdjacent(currentSeats, position{i, j}); r == 'L' && os == 0 {
					newRune = '#'
				} else if r == '#' && os >= 4 {
					newRune = 'L'
				}
				row = append(row, newRune)
			}
			newSeats = append(newSeats, row)
		}
		if reflect.DeepEqual(currentSeats, newSeats) {
			break rounds
		}
		currentSeats = newSeats
	}

	fmt.Println("Part 1:", countOccupied(currentSeats))

	// Part 2

	currentSeats = scannedSeats
rounds2:
	for {
		newSeats := [][]rune{}
		for i := range currentSeats {
			row := []rune{}
			for j, r := range currentSeats[i] {
				newRune := r
				if cs := canSeeAdjacent(currentSeats, position{i, j}); r == 'L' && cs == 0 {
					newRune = '#'
				} else if r == '#' && cs >= 5 {
					newRune = 'L'
				}
				row = append(row, newRune)
			}
			newSeats = append(newSeats, row)
		}
		if reflect.DeepEqual(currentSeats, newSeats) {
			break rounds2
		}
		currentSeats = newSeats
	}

	fmt.Println("Part 2:", countOccupied(currentSeats))
}

type position struct {
	row, pos int
}

var dirs = []position{
	{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
}

func direction(pos, dir position) position {
	return position{pos.row + dir.row, pos.pos + dir.pos}
}

func isSafe(seats [][]rune, p position) bool {
	return !(p.row < 0 || p.row >= len(seats) || p.pos < 0 || p.pos >= len(seats[p.row]))
}

func occupied(seats [][]rune, p position) bool {
	return isSafe(seats, p) && seats[p.row][p.pos] == '#'
}

func occupiedAdjacent(seats [][]rune, pos position) (o int) {
	for _, dir := range dirs {
		if occupied(seats, direction(pos, dir)) {
			o++
		}
	}
	return
}

func canSee(seats [][]rune, p, dir position) bool {
	return isSafe(seats, p) && seats[p.row][p.pos] != 'L' && (seats[p.row][p.pos] == '#' || canSee(seats, direction(p, dir), dir))
}

func canSeeAdjacent(seats [][]rune, pos position) (o int) {
	for _, dir := range dirs {
		if canSee(seats, direction(pos, dir), dir) {
			o++
		}
	}
	return
}

func countOccupied(seats [][]rune) (o int) {
	for _, ro := range seats {
		o += strings.Count(string(ro), "#")
	}
	return
}
