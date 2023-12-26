package inmemory_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/storage"
	"github.com/matthiashermsen/kaboom/storage/inmemory"
)

func TestNew(testing *testing.T) {
	store := inmemory.New()

	assert.Nil(testing, store.Board, "Expected board to be nil")
}

func TestGetBoard(suite *testing.T) {
	suite.Run("Returns the current board", func(testing *testing.T) {
		store := inmemory.New()

		amountOfRows := uint(2)
		amountOfColumns := uint(2)
		err := store.CreateNewBoard(amountOfRows, amountOfColumns, 2)

		assert.NoError(testing, err, "Expected no error when creating board")
		assert.NotNil(testing, store.Board, "Expected board not to be nil")

		cells, err := store.GetBoard()

		assert.NoError(testing, err, "Expected no error when returning board")

		assert.Equal(testing, int(amountOfRows), len(cells), fmt.Sprintf("Expected cell rows to have a length %d but got %d", amountOfRows, len(cells)))

		for _, row := range cells {
			assert.Equal(testing, int(amountOfColumns), len(row), fmt.Sprintf("Expected row columns to have a length %d but got %d", amountOfColumns, len(row)))
		}
	})

	suite.Run("Returns error if board is nil", func(testing *testing.T) {
		store := inmemory.New()

		_, err := store.GetBoard()

		assert.ErrorIs(testing, err, &storage.BoardMissingError{}, "Expected BoardMissingError but got %v", err)
	})

	suite.Run("Concurrent access", func(testing *testing.T) {
		store := inmemory.New()

		amountOfRows := uint(2)
		amountOfColumns := uint(2)
		err := store.CreateNewBoard(amountOfRows, amountOfColumns, 2)
		assert.NoError(testing, err, "Expected no error when creating new board")

		var waitGroup sync.WaitGroup
		amountOfConcurrentCalls := 10

		waitGroup.Add(amountOfConcurrentCalls)

		for i := 0; i < amountOfConcurrentCalls; i++ {
			go func() {
				defer waitGroup.Done()
				cells, err := store.GetBoard()

				assert.NoError(testing, err, "Expected no error when returning board")

				assert.Equal(testing, int(amountOfRows), len(cells), fmt.Sprintf("Expected cell rows to have a length %d but got %d", amountOfRows, len(cells)))

				for _, row := range cells {
					assert.Equal(testing, int(amountOfColumns), len(row), fmt.Sprintf("Expected row columns to have a length %d but got %d", amountOfColumns, len(row)))
				}
			}()
		}

		waitGroup.Wait()
	})
}

func TestCreateNewBoard(suite *testing.T) {
	suite.Run("Creates new board", func(testing *testing.T) {
		store := inmemory.New()

		err := store.CreateNewBoard(2, 2, 2)

		assert.NoError(testing, err, "Expected no error for valid settings")
		assert.NotNil(testing, store.Board, "Expected board not to be nil")
	})

	suite.Run("Invalid settings", func(testing *testing.T) {
		store := inmemory.New()

		err := store.CreateNewBoard(2, 2, 4)

		assert.Error(testing, err, "Expected error for invalid settings")
		assert.Nil(testing, store.Board, "Expected board to be nil")
	})

	suite.Run("Concurrent access", func(testing *testing.T) {
		store := inmemory.New()

		var waitGroup sync.WaitGroup
		amountOfConcurrentCalls := 10

		waitGroup.Add(amountOfConcurrentCalls)

		for i := 0; i < amountOfConcurrentCalls; i++ {
			go func() {
				defer waitGroup.Done()
				err := store.CreateNewBoard(2, 2, 2)

				assert.NoError(testing, err, "Expected no error when creating new board")
				assert.NotNil(testing, store.Board, "Expected board not to be nil during concurrent access")
			}()
		}

		waitGroup.Wait()
	})
}
