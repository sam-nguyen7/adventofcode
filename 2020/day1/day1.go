package day1

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// open file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read input.txt then convert to int slice
	rd := bufio.NewScanner(file)
	var numbers []int
	for rd.Scan() {
		num, _ := strconv.Atoi(rd.Text())
		numbers = append(numbers, num)
	}

	part1, _ := part1(numbers)
	fmt.Println("Product of nums:", part1)

}

func part1(numbers []int) (int, error) {
	// loop though each int in slice
	// nested for loop to compare ints
	// break if same int
	for a, b := range numbers {
		for i, j := range numbers {
			if i == a {
				break
			}
			if b+j == 2020 {
				return a * b, nil
			}
		}
	}
	err := errors.New("No entries found")
	return 0, err
}
