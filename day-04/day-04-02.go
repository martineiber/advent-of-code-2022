package main

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// input File should be converted from CRLF to LF (\n Chr(10))
func main() {
	absPath, _ := filepath.Abs("day-04/input2.txt")
	data, _ := os.ReadFile(absPath)
	content := string(data)
	pairs := strings.Split(content, "\n")
	var sum int
	for _, value := range pairs {
		pair := strings.Split(value, ",")

		part1 := strings.Split(pair[0], "-")
		part2 := strings.Split(pair[1], "-")
		if overlappPartial(part1, part2) {
			sum++
		}
	}
	println(sum)
}

func overlappPartial(part1 []string, part2 []string) bool {
	value_pair_1_1, _ := strconv.ParseInt(part1[0], 10, 32)
	value_pair_1_2, _ := strconv.ParseInt(part1[1], 10, 32)
	value_pair_2_1, _ := strconv.ParseInt(part2[0], 10, 32)
	value_pair_2_2, _ := strconv.ParseInt(part2[1], 10, 32)

	range1 := getFullArray(value_pair_1_1, value_pair_1_2)
	range2 := getFullArray(value_pair_2_1, value_pair_2_2)

	for _, value := range range1 {
		if contains(value, range2) {
			return true
		}
	}
	return false
}

func getFullArray(from int64, to int64) []int64 {
	var val []int64
	for i := from; i <= to; i++ {
		val = append(val, i)
	}

	return val
}

func contains(needle int64, haystack []int64) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}
	return false
}
