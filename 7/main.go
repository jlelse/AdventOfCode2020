package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	rules := map[string]map[string]int{}

	mainColorRe := regexp.MustCompile(`^(?P<color>\w+ \w+) bags contain (?P<contents>.*)\.$`)
	contentRe := regexp.MustCompile(`(?P<amount>\d+) (?P<color>\w+ \w+) bags?`)

	file, _ := os.Open("input.txt")
	defer file.Close()
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		text := scanner.Text()
		rule := map[string]int{}
		for _, c := range contentRe.FindAllString(mainColorRe.ReplaceAllString(text, "$contents"), -1) {
			rule[contentRe.ReplaceAllString(c, "$color")], _ = strconv.Atoi(contentRe.ReplaceAllString(c, "$amount"))
		}
		rules[mainColorRe.ReplaceAllString(text, "$color")] = rule
	}

	fmt.Println("Part 1:", len(unique(containedIn("shiny gold", rules))))
	fmt.Println("Part 2:", numberOfBags("shiny gold", rules)-1)
}

func containedIn(target string, rules map[string]map[string]int) (colors []string) {
	for k, r := range rules {
		if _, has := r[target]; has {
			colors = append(colors, append(containedIn(k, rules), k)...)
		}
	}
	return
}

func unique(s []string) (list []string) {
	k := map[string]bool{}
	for _, e := range s {
		if _, has := k[e]; !has {
			k[e] = true
			list = append(list, e)
		}
	}
	return
}

func numberOfBags(target string, rules map[string]map[string]int) int {
	bags := 1
	for color, amount := range rules[target] {
		bags += amount * numberOfBags(color, rules)
	}
	return bags
}
