package main

import (
	"fmt"
	game "go-minesweeper/cmd/domain/game/models"
)

func main() {
	fmt.Println("Hello World")

	X_SIZE := 15
	Y_SIZE := 10
	DIFFICULTY := game.Easy

	gameBoard := game.NewBoard(X_SIZE, Y_SIZE, DIFFICULTY)
}
