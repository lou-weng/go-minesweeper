package models

import (
	enums "go-minesweeper/cmd/domain/game/enums"
)

type GameBoard struct {
	board *[][]enums.GameTile
	x_size int
	y_size int
	difficulty enums.Difficulty
}

func NewBoard(x_size int,  y_size int, difficulty enums.Difficulty) GameBoard {
	initialize_board := initialize_board(x_size, y_size)
	return GameBoard{
		board: initialize_board, 
		x_size: x_size, 
		y_size: y_size, 
		difficulty: difficulty}
}

func initialize_board(x_size int, y_size int) *[][]enums.GameTile {
	initialize_board := make([][]enums.GameTile, y_size)
	for i := range initialize_board {
		initialize_board[i] = make([]enums.GameTile, x_size)
	}
	return &initialize_board
}

func (b GameBoard) GetBoard() *[][]enums.GameTile {
	return b.board
}
