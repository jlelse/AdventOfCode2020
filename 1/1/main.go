package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numbers := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, n)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, i := range numbers {
		for _, j := range numbers {
			if i+j == 2020 {
				fmt.Println(i * j)
				return
			}
		}
	}

}
