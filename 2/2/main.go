package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// While it appears you validated the passwords correctly, they don't seem to be what the Official Toboggan Corporate Authentication System is expecting.

// The shopkeeper suddenly realizes that he just accidentally explained the password policy rules from his old job at the sled rental place down the street! The Official Toboggan Corporate Policy actually works a little differently.

// Each policy actually describes two positions in the password, where 1 means the first character, 2 means the second character, and so on. (Be careful; Toboggan Corporate Policies have no concept of "index zero"!) Exactly one of these positions must contain the given letter. Other occurrences of the letter are irrelevant for the purposes of policy enforcement.

// Given the same example list from above:

//     1-3 a: abcde is valid: position 1 contains a and position 3 does not.
//     1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
//     2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.

// How many passwords are valid according to the new interpretation of the policies?

type password struct {
	first, second    int
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
		first, _ := strconv.Atoi(inputRegex.ReplaceAllString(text, "$1"))
		second, _ := strconv.Atoi(inputRegex.ReplaceAllString(text, "$2"))
		passwords = append(passwords, &password{
			first:    first,
			second:   second,
			letter:   inputRegex.ReplaceAllString(text, "$3"),
			password: inputRegex.ReplaceAllString(text, "$4"),
		})
	}

	correct := 0
	for _, pw := range passwords {
		if (string(pw.password[pw.first-1]) == pw.letter) != (string(pw.password[pw.second-1]) == pw.letter) {
			correct++
		}
	}

	fmt.Println(correct)
	return
}
