package main

import (
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	absPath, _ := filepath.Abs("day-01/input.txt")
	data, _ := os.ReadFile(absPath)
	content := string(data)

	elfs := strings.Split(content, "\r\n\r\n")

	var list []int

	for _, value := range elfs {
		calories := strings.Split(value, "\r\n")
		var sum int
		for _, calorie := range calories {
			intCal, _ := strconv.ParseInt(calorie, 10, 32)
			sum = int(intCal) + sum
		}
		list = append(list, sum)
	}
	sort.Ints(list)

	var l1 int = list[len(list)-1]
	var l2 int = list[len(list)-2]
	var l3 int = list[len(list)-3]

	var sum_top3 = l1 + l2 + l3
	println(sum_top3)
	println(l1)
}
