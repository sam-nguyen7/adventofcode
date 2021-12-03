package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// get file & split by newline
	file, _ := os.ReadFile("input.txt")
	splitString := strings.Split(string(file), "\n")

	previous := -1
	count := 0

	for _, stringData := range splitString {

		// convert string to int
		number, _ := strconv.Atoi(stringData)
		// save prev and cur num
		// compare & count
		if previous != -1 {
			if number > previous {
				count++
			}
		}
		previous = number
	}

	fmt.Println(count)
}
