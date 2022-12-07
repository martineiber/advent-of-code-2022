package main

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// input File should be converted from CRLF to LF (\n Chr(10))
func main() {
	absPath, _ := filepath.Abs("day-04/input1.txt")
	data, _ := os.ReadFile(absPath)
	content := string(data)
	pairs := strings.Split(content, "\n")
	var sum int
	for _, value := range pairs {
		pair := strings.Split(value, ",")

		part1 := strings.Split(pair[0], "-")
		part2 := strings.Split(pair[1], "-")
		if overlapp(part1, part2) {
			sum++
		}
	}
	println(sum)
}

func overlapp(part1 []string, part2 []string) bool {
	value_pair_1_1, _ := strconv.ParseInt(part1[0], 10, 32)
	value_pair_1_2, _ := strconv.ParseInt(part1[1], 10, 32)
	value_pair_2_1, _ := strconv.ParseInt(part2[0], 10, 32)
	value_pair_2_2, _ := strconv.ParseInt(part2[1], 10, 32)

	if value_pair_1_1 >= value_pair_2_1 && value_pair_1_2 <= value_pair_2_2 {
		return true

	}

	if value_pair_2_1 >= value_pair_1_1 && value_pair_2_2 <= value_pair_1_2 {
		return true
	}

	return false
}
