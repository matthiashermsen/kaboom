package inmemory

import (
	"sync"

	"github.com/matthiashermsen/kaboom/board"
	"github.com/matthiashermsen/kaboom/storage"
)

type Store struct {
	mutex sync.RWMutex
	Board *board.Board
}

func New() *Store {
	return &Store{}
}

func (store *Store) GetBoard() ([][]board.Cell, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	if store.Board == nil {
		return [][]board.Cell{}, &storage.BoardMissingError{}
	}

	return store.Board.Cells, nil
}

func (store *Store) CreateNewBoard(amountOfRows uint, amountOfColumns uint, amountOfMines uint) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	newBoard, err := board.New(amountOfRows, amountOfColumns, amountOfMines)

	if err != nil {
		return err
	}

	store.Board = newBoard

	return nil
}
