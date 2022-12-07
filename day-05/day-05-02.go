package main

import (
	"container/list"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// input File should be converted from CRLF to LF (\n Chr(10))
func main() {
	absPath, _ := filepath.Abs("day-05/input2.txt")
	data, _ := os.ReadFile(absPath)
	content := string(data)
	parts := strings.Split(content, "\n\n")
	shipStack := generateShitStack(parts[0])
	moveList := generateMoveList(parts[1])
	reorder(&shipStack, moveList)

	for _, stack := range shipStack.stack {
		value := stack.items.Back().Value
		fmt.Print(value)
	}
}

func reorder(stack *ShipStack, moveList []Move) {
	for _, move := range moveList {
		moveStacks(stack, move)
	}
}

func moveStacks(stack *ShipStack, move Move) {
	list := list.New()
	for i := 0; i < move.amount; i++ {
		item := stack.stack[move.from-1].items.Back()
		list.PushFront(item.Value)
		stack.stack[move.from-1].items.Remove(item)
	}

	for e := list.Front(); e != nil; e = e.Next() {
		stack.stack[move.to-1].items.PushBack(e.Value)
	}

}

func generateMoveList(input string) []Move {
	var moves []Move
	lines := strings.Split(input, "\n")
	for _, value := range lines {
		parts := strings.Split(value, " ")
		amount, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])
		moves = append(moves, Move{amount, from, to})
	}

	return moves
}

func generateShitStack(input string) ShipStack {
	parts := strings.Split(input, "\n")
	var shipstask = initStacks(parts)
	parts = parts[:len(parts)-1]
	for _, part := range parts {
		var stackNumber = 0
		for i := 1; i < len(part); i = i + 4 {
			if part[i] != 32 {
				shipstask.stack[stackNumber].items.PushFront(string(part[i]))
			}
			stackNumber++
		}
	}

	return shipstask
}

func initStacks(allLines []string) ShipStack {
	var size int
	for _, line := range allLines {
		if len(line)/4 > size {
			size = len(line) / 4
		}
	}
	var stack []Stack
	for i := 0; i <= size; i++ {
		stack = append(stack, Stack{stackNumber: i + 1})
	}
	return ShipStack{stack: stack}
}

type Move struct {
	amount int
	from   int
	to     int
}

type ShipStack struct {
	stack []Stack
}

type Stack struct {
	items       list.List
	stackNumber int
}
