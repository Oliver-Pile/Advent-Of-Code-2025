package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
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

func part2() int {
	file, _ := os.Open("day5/day5.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	unique := make(map[int]struct{})

	var mu sync.Mutex
	var wg sync.WaitGroup
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "-") {
			r := calculateRange(line)

			wg.Add(1)
			go addFresh(r, unique, &mu, &wg)

		}
	}

	wg.Wait()

	return len(unique)
}

func addFresh(r Range, unique map[int]struct{}, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := r.lower; i <= r.upper; i++ {
		mu.Lock()
		unique[i] = struct{}{}
		mu.Unlock()

	}
}
