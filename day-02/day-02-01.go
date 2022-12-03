package main

import (
	"os"
	"path/filepath"
	"strings"
)

func main() {
	absPath, _ := filepath.Abs("day-02/input1.txt")
	data, _ := os.ReadFile(absPath)
	content := string(data)
	rounds := strings.Split(content, "\r\n")
	var points int
	for _, value := range rounds {
		s := strings.Fields(value)
		var opponentMove string = s[0]
		var myMove string = convertMove(s[1])

		if hasWon(opponentMove, myMove) {
			points += pointOfFigure(myMove)
			points += 6
		} else if isDraw(opponentMove, myMove) {
			points += pointOfFigure(myMove)
			points += 3
		} else {
			// lost
			points += pointOfFigure(myMove)
		}
	}

	println(points)
}

func hasWon(opponentMove string, myMove string) bool {
	if myMove == "A" && opponentMove == "C" { // Rock wins vs. Scissors
		return true
	}

	if myMove == "B" && opponentMove == "A" { // Paper wins vs. Rock
		return true
	}

	if myMove == "C" && opponentMove == "B" { // Scissors wins vs. Paper
		return true
	}

	return false
}

func isDraw(opponentMove string, myMove string) bool {
	if opponentMove == myMove {
		return true
	}
	return false
}

func convertMove(move string) string {
	if move == "Z" { // Scissors
		return "C"
	}

	if move == "Y" { // Paper
		return "B"
	}

	if move == "X" { // Rock
		return "A"
	}

	return move // should not happen ;)
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
