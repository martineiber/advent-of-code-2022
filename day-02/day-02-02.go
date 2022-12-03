package main

import (
	"os"
	"path/filepath"
	"strings"
)

func main() {
	absPath, _ := filepath.Abs("day-02/input2.txt")
	data, _ := os.ReadFile(absPath)
	content := string(data)
	rounds := strings.Split(content, "\r\n")
	var points int
	for _, value := range rounds {
		s := strings.Fields(value)
		var opponentMove string = s[0]
		var myStrategy string = s[1]

		if myStrategy == "X" { // lose
			points += pointOfFigure(howToLose(opponentMove))
		} else if myStrategy == "Y" { // draw
			points += pointOfFigure(opponentMove)
			points += 3
		} else if myStrategy == "Z" { // win
			points += pointOfFigure(howToWin(opponentMove))
			points += 6
		}
	}

	println(points)
}

func pointOfFigure(figure string) int {
	if figure == "A" { // Rock
		return 1
	}

	if figure == "B" { // Paper
		return 2
	}

	if figure == "C" { // Scissors
		return 3
	}

	return 0
}

func howToLose(opponentMove string) string {
	if opponentMove == "A" { // Rock
		return "C" // Scissors
	}

	if opponentMove == "B" { // Paper
		return "A" // Rock
	}

	if opponentMove == "C" { // Scissors
		return "B" // Paper
	}

	return "" // should not happen
}

func howToWin(opponentMove string) string {
	if opponentMove == "A" { // Rock
		return "B" // Paper
	}

	if opponentMove == "B" { // Paper
		return "C" // Scissors
	}

	if opponentMove == "C" { // Scissors
		return "A" // Rock
	}

	return "" // should not happen
}
