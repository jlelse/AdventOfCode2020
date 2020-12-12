package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	type command struct {
		action rune
		value  int
	}
	var commands []command

	file, _ := os.Open("input.txt")
	defer file.Close()
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		text := scanner.Text()
		value, _ := strconv.Atoi(text[1:])
		commands = append(commands, command{
			action: rune(text[0]),
			value:  value,
		})
	}

	shipNorth, shipEast, shipDirection := 0, 0, 90
	for _, cmd := range commands {
		switch cmd.action {
		case 'N':
			shipNorth += cmd.value
		case 'S':
			shipNorth -= cmd.value
		case 'E':
			shipEast += cmd.value
		case 'W':
			shipEast -= cmd.value
		case 'L':
			shipDirection -= cmd.value
			shipDirection %= 360
		case 'R':
			shipDirection += cmd.value
			shipDirection %= 360
		case 'F':
			switch shipDirection {
			case 0:
				shipNorth += cmd.value
			case 90, -270:
				shipEast += cmd.value
			case 180, -180:
				shipNorth -= cmd.value
			case 270, -90:
				shipEast -= cmd.value
			}
		}
	}

	fmt.Println("Part 1:", math.Abs(float64(shipNorth))+math.Abs(float64(shipEast)))

	wpNorth, wpEast := 1, 10
	shipNorth = 0
	shipEast = 0
	var rotateWaypoint func(int)
	rotateWaypoint = func(degree int) {
		switch degree {
		case 90, -270:
			newWpNorth, newWpEast := -wpEast, wpNorth
			wpNorth = newWpNorth
			wpEast = newWpEast
		case 180, -180:
			rotateWaypoint(90)
			rotateWaypoint(90)
		case 270, -90:
			rotateWaypoint(90)
			rotateWaypoint(90)
			rotateWaypoint(90)
		}
	}
	for _, cmd := range commands {
		switch cmd.action {
		case 'N':
			wpNorth += cmd.value
		case 'S':
			wpNorth -= cmd.value
		case 'E':
			wpEast += cmd.value
		case 'W':
			wpEast -= cmd.value
		case 'L':
			rotateWaypoint(-cmd.value)
		case 'R':
			rotateWaypoint(cmd.value)
		case 'F':
			shipNorth += wpNorth * cmd.value
			shipEast += wpEast * cmd.value
		}
	}

	fmt.Println("Part 2:", math.Abs(float64(shipNorth))+math.Abs(float64(shipEast)))
}
