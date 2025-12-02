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

	fmt.Printf("%v\n", part1Count)
	fmt.Printf("%v\n", part2Count)

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

	sum := 0
	for _, id := range invalidIDs {
		sum += id
	}
	return sum

}
func part2() int {

	return 0
}
