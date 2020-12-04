package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	passports := []map[string]string{}

	passport := map[string]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			// Empty row, new passport
			passports = append(passports, passport)
			passport = map[string]string{}
		} else {
			for _, part := range strings.Split(text, " ") {
				keyValue := strings.Split(part, ":")
				passport[keyValue[0]] = keyValue[1]
			}
		}
	}
	// Last passport
	passports = append(passports, passport)

	count := 0
ppLoop:
	for _, pp := range passports {
		for _, field := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
			if _, fieldPresent := pp[field]; !fieldPresent {
				continue ppLoop
			}
		}
		count++
	}

	fmt.Println("Part 1:", count)

	// Part 2, stricter rules

	count = 0

	yrValid := func(byr string, min, max int) bool {
		byrInt, err := strconv.Atoi(byr)
		if err != nil || byrInt < min || byrInt > max {
			return false
		}
		return true
	}

	hgtValid := func(hgt string) bool {
		if strings.HasSuffix(hgt, "cm") {
			if no, err := strconv.Atoi(strings.TrimSuffix(hgt, "cm")); err != nil || no < 150 || no > 193 {
				return false
			}
		} else if strings.HasSuffix(hgt, "in") {
			if no, err := strconv.Atoi(strings.TrimSuffix(hgt, "in")); err != nil || no < 59 || no > 76 {
				return false
			}
		} else {
			return false
		}
		return true
	}

	hclValid := func(hcl string) bool {
		return regexp.MustCompile(`^#[a-f0-9]{6}$`).Match([]byte(hcl))
	}

	eclValid := func(ecl string) bool {
		return ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth"
	}

	pidValid := func(pid string) bool {
		if len(pid) != 9 {
			return false
		}
		if id, err := strconv.Atoi(pid); err != nil || id == 0 {
			return false
		}
		return true
	}

	for _, pp := range passports {
		if byr, byrPresent := pp["byr"]; !(byrPresent && yrValid(byr, 1920, 2002)) {
			continue
		}
		if iyr, iyrPresent := pp["iyr"]; !(iyrPresent && yrValid(iyr, 2010, 2020)) {
			continue
		}
		if eyr, eyrPresent := pp["eyr"]; !(eyrPresent && yrValid(eyr, 2020, 2030)) {
			continue
		}
		if hgt, hgtPresent := pp["hgt"]; !(hgtPresent && hgtValid(hgt)) {
			continue
		}
		if hcl, hclPresent := pp["hcl"]; !(hclPresent && hclValid(hcl)) {
			continue
		}
		if ecl, eclPresent := pp["ecl"]; !(eclPresent && eclValid(ecl)) {
			continue
		}
		if pid, pidPresent := pp["pid"]; !(pidPresent && pidValid(pid)) {
			continue
		}
		count++
	}

	fmt.Println("Part 2:", count)
}
