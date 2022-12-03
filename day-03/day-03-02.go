package main

import (
	"os"
	"path/filepath"
	"strings"
)

// input File should be converted from CRLF to LF (\n Chr(10))
func main() {
	absPath, _ := filepath.Abs("day-03/input2.txt")
	data, _ := os.ReadFile(absPath)
	content := string(data)
	backpacks := strings.Split(content, "\n")
	var sum int

	for i := 0; i < len(backpacks); i = i + 3 {
		byteArray1 := []byte(backpacks[i])
		byteArray2 := []byte(backpacks[i+1])
		byteArray3 := []byte(backpacks[i+2])
		duplicate := findDuplicate(byteArray1, byteArray2, byteArray3)
		sum += getCharValue(duplicate[0])
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

func findDuplicate(part1 []byte, part2 []byte, part3 []byte) []byte {
	var duplicates []byte
	for _, value := range part1 {
		if contains(value, part2) && contains(value, part3) {
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
