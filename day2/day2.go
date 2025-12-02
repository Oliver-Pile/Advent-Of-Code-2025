package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1Count := part1()
	part2Count := part2()

	fmt.Printf("Part1: %v\n", part1Count)
	fmt.Printf("Part2: %v\n", part2Count)

}

func part1() int {
	input, _ := os.ReadFile("day2/day2.txt")

	splitValues := strings.Split(string(input), ",")

	var invalidIDs []int
	for _, vals := range splitValues {
		pair := strings.Split(vals, "-")
		lower, _ := strconv.Atoi(pair[0])
		upper, _ := strconv.Atoi(pair[1])

		for i := lower; i <= upper; i++ {
			strVal := strconv.Itoa(i)
			length := len(strVal)
			if length%2 != 0 {
				continue
			}

			if strVal[:length/2] == strVal[length/2:] {
				invalidIDs = append(invalidIDs, i)
			}

		}
	}
	return sum(invalidIDs)

}
func part2() int {
	input, _ := os.ReadFile("day2/day2.txt")

	splitValues := strings.Split(string(input), ",")

	var invalidIDs []int
	for _, vals := range splitValues {
		pair := strings.Split(vals, "-")
		lower, _ := strconv.Atoi(pair[0])
		upper, _ := strconv.Atoi(pair[1])
		fmt.Printf("%v-%v\n", lower, upper)

		for i := lower; i <= upper; i++ {
			strVal := strconv.Itoa(i)
			strLen := len(strVal)
			for j := 2; j <= strLen; j++ {
				partial := strVal[:strLen/j]

				if strings.Repeat(partial, j) == strVal {
					fmt.Printf("Invalid: %v\n", i)
					invalidIDs = append(invalidIDs, i)
					break
				}

			}

		}
	}
	return sum(invalidIDs)
}

func sum(input []int) int {
	sum := 0
	for _, val := range input {
		sum += val
	}
	return sum
}
