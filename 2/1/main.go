package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Your flight departs in a few days from the coastal airport; the easiest way down to the coast from here is via toboggan.

// The shopkeeper at the North Pole Toboggan Rental Shop is having a bad day. "Something's wrong with our computers; we can't log in!" You ask if you can take a look.

// Their password database seems to be a little corrupted: some of the passwords wouldn't have been allowed by the Official Toboggan Corporate Policy that was in effect when they were chosen.

// To try to debug the problem, they have created a list (your puzzle input) of passwords (according to the corrupted database) and the corporate policy when that password was set.

// For example, suppose you have the following list:

// 1-3 a: abcde
// 1-3 b: cdefg
// 2-9 c: ccccccccc

// Each line gives the password policy and then the password. The password policy indicates the lowest and highest number of times a given letter must appear for the password to be valid. For example, 1-3 a means that the password must contain a at least 1 time and at most 3 times.

// In the above example, 2 passwords are valid. The middle password, cdefg, is not; it contains no instances of b, but needs at least 1. The first and third passwords are valid: they contain one a or nine c, both within the limits of their respective policies.

// How many passwords are valid according to their policies?

type password struct {
	min, max         int
	letter, password string
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	passwords := []*password{}

	inputRegex := regexp.MustCompile(`(?m)^(\d+)-(\d+) (\w): (\w+)$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		minString := inputRegex.ReplaceAllString(text, "$1")
		min, _ := strconv.Atoi(minString)
		maxString := inputRegex.ReplaceAllString(text, "$2")
		max, _ := strconv.Atoi(maxString)
		pw := &password{
			min:      min,
			max:      max,
			letter:   inputRegex.ReplaceAllString(text, "$3"),
			password: inputRegex.ReplaceAllString(text, "$4"),
		}
		passwords = append(passwords, pw)
	}

	correct := 0
	for _, pw := range passwords {
		if count := strings.Count(pw.password, pw.letter); count >= pw.min && count <= pw.max {
			correct++
		}
	}

	fmt.Println(correct)
	return
}