package main

import (
	"os"
	"path/filepath"
	"strings"
)

// input File should be converted from CRLF to LF (\n Chr(10))
func main() {
	absPath, _ := filepath.Abs("day-03/input1.txt")
	data, _ := os.ReadFile(absPath)
	content := string(data)
	backpacks := strings.Split(content, "\n")
	var sum int
	for _, value := range backpacks {
		byteArray := []byte(value)
		length := len(byteArray)
		part1 := byteArray[0 : length/2]
		part2 := byteArray[length/2 : length]
		duplicates := findDuplicate(part1, part2)
		for _, byteValue := range duplicates {
			sum += getCharValue(byteValue)
		}
	}
	println(sum)
}

func getCharValue(value byte) int {
	if value >= 65 && value <= 90 {
		return int(value - 38)
	}

	if value >= 97 && value <= 122 { // lowecase
		return int(value - 96)
	}

	return 0
}

func findDuplicate(part1 []byte, part2 []byte) []byte {
	var duplicates []byte
	for _, value := range part1 {
		if contains(value, part2) {
			if !contains(value, duplicates) {
				duplicates = append(duplicates, value)
			}
		}
	}
	return duplicates
}

func contains(char uint8, string []byte) bool {
	for _, value := range string {
		if value == char {
			return true
		}
	}
	return false
}
