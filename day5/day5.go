package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	lower int
	upper int
}

func main() {

	part1Count := part1()
	part2Count := part2()

	fmt.Printf("Part 1: %v\n", part1Count)
	fmt.Printf("Part 2: %v\n", part2Count)

}

func part1() int {
	file, _ := os.Open("day5/day5.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ranges := []Range{}
	totalFresh := 0
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "-") {
			ranges = addRange(line, ranges)
		} else {
			isFresh := checkFresh(line, ranges)

			if isFresh {
				totalFresh++
			}
		}
	}

	return totalFresh
}

func addRange(line string, ranges []Range) []Range {
	splitLine := strings.Split(line, "-")

	lower, _ := strconv.Atoi(splitLine[0])
	upper, _ := strconv.Atoi(splitLine[1])

	r := Range{lower: lower, upper: upper}

	return append(ranges, r)
}

func checkFresh(line string, ranges []Range) bool {
	val, err := strconv.Atoi(line)

	if err != nil {
		return false
	}

	for _, r := range ranges {
		if val >= r.lower && val <= r.upper {
			return true
		}
	}

	return false
}

func part2() any {
	return 0
}
