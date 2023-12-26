package storage

import "github.com/matthiashermsen/kaboom/board"

type Store interface {
	GetBoard() ([][]board.Cell, error)
	CreateNewBoard(amountOfRows uint, amountOfColumns uint, amountOfMines uint) error
}
