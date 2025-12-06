package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	r := calculateRange(line)
	return append(ranges, r)
}

func calculateRange(line string) Range {
	splitLine := strings.Split(line, "-")

	lower, _ := strconv.Atoi(splitLine[0])
	upper, _ := strconv.Atoi(splitLine[1])

	return Range{lower: lower, upper: upper}
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

// Optimised by AI (Otherwise it would have taken too long!)
func part2() int {
	file, _ := os.Open("day5/day5.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var ranges []Range
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "-") {
			r := calculateRange(line)
			ranges = append(ranges, r)
		}
	}

	// Step 1: sort by lower bound
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].lower < ranges[j].lower
	})

	// Step 2: merge overlapping ranges
	merged := make([]Range, 0, len(ranges))
	current := ranges[0]

	for _, r := range ranges[1:] {
		if r.lower <= current.upper+1 {
			// overlapping or adjacent → merge
			if r.upper > current.upper {
				current.upper = r.upper
			}
		} else {
			merged = append(merged, current)
			current = r
		}
	}
	merged = append(merged, current)

	// Step 3: count total unique ints
	total := 0
	for _, r := range merged {
		total += (r.upper - r.lower + 1)
	}

	return total
}
