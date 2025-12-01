package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	part1Count := part1()

	fmt.Printf("%v\n", part1Count)

}

func part1() int {
	file, err := os.Open("day1/day1.txt")
	if err != nil {
		log.Printf("Error %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	position := 50
	countZero := 0

	for scanner.Scan() {
		line := scanner.Text()

		operation := string(line[0])
		number, _ := strconv.Atoi(line[1:])

		position = calculatePosition(position, number, operation)
		if position == 0 {
			countZero++
		}

	}

	return countZero
}

func calculatePosition(position, number int, operation string) int {
	moddedNumber := number % 100

	if operation == "L" {
		newPosition := position - moddedNumber
		return handleNegative(newPosition)
	} else {
		newPosition := position + moddedNumber
		return handleOver99(newPosition)
	}
}

func handleNegative(number int) int {
	if number < 0 {
		return number + 100
	}

	return number
}
func handleOver99(number int) int {
	if number > 99 {
		return number - 100
	}

	return number
}
