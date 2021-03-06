package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// The Elves in accounting are thankful for your help; one of them even offers you a starfish coin they had left over from a past vacation. They offer you a second one if you can find three numbers in your expense report that meet the same criteria.

// Using the above example again, the three entries that sum to 2020 are 979, 366, and 675. Multiplying them together produces the answer, 241861950.

// In your expense report, what is the product of the three entries that sum to 2020?

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	numbers := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, n)
	}

	for _, i := range numbers {
		for _, j := range numbers {
			for _, k := range numbers {
				if i+j+k == 2020 {
					fmt.Println(i * j * k)
					return
				}
			}
		}
	}

}
