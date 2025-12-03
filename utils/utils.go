package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(path string) *bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error %v", err)
	}
	defer file.Close()

	return bufio.NewScanner(file)
}
