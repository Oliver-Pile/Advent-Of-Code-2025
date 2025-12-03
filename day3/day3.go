package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("day3/day3.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1Count := part1(scanner)
	part2Count := part2()

	fmt.Printf("Part 1: %v\n", part1Count)
	fmt.Printf("Part 2: %v\n", part2Count)

}

func part1(scanner *bufio.Scanner) int {
	fmt.Println("In func")

	var total int
	for scanner.Scan() {
		fmt.Println("In Loop")
		line := scanner.Text()
		length := len(line)
		largest := findLargest(line[:length-1])
		fmt.Printf("Largest: %s\n", largest)
		largestIndex := strings.Index(line, largest)
		secondLargest := findLargest(line[largestIndex+1:])
		fmt.Printf("Second Largest: %s\n", secondLargest)

		strNum := fmt.Sprintf("%s%s", largest, secondLargest)

		num, _ := strconv.Atoi(strNum)
		total += num

	}

	return total
}

func findLargest(line string) string {
	var largest int
	for _, char := range line {
		num, _ := strconv.Atoi(string(char))
		if num > largest {
			largest = num
		}
	}
	return strconv.Itoa(largest)
}

func part2() int {
	return 0
}
