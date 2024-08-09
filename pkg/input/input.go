package input

import (
	"bufio"
	"log"
	"os"
)

func GetInputMessages(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	// Print all lines (for verification purposes)
	for idx, line := range lines {
		log.Printf("Line %d: %s\n", idx+1, line)
	}

	return lines
}