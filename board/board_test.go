package board_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/board"
)

func TestNew(suite *testing.T) {
	suite.Run("Creates new board", func(testing *testing.T) {
		amountOfRows := uint(4)
		amountOfColumns := uint(8)
		amountOfMines := uint(3)

		newBoard, err := board.New(amountOfRows, amountOfColumns, amountOfMines)

		assert.NoError(testing, err, "Expected no error for valid settings")
		assert.NotNil(testing, newBoard, "Expected board not to be nil")
		assert.Equal(testing, amountOfRows, newBoard.AmountOfRows, fmt.Sprintf("Expected rows to be %d but got %d", amountOfRows, newBoard.AmountOfRows))
		assert.Equal(testing, amountOfColumns, newBoard.AmountOfColumns, fmt.Sprintf("Expected columns to be %d but got %d", amountOfColumns, newBoard.AmountOfColumns))
		assert.Equal(testing, amountOfMines, newBoard.AmountOfMines, fmt.Sprintf("Expected mines to be %d but got %d", amountOfMines, newBoard.AmountOfMines))

		assert.Equal(testing, int(amountOfRows), len(newBoard.Cells), fmt.Sprintf("Expected cell rows to have a length %d but got %d", amountOfRows, len(newBoard.Cells)))

		for _, row := range newBoard.Cells {
			assert.Equal(testing, int(amountOfColumns), len(row), fmt.Sprintf("Expected row columns to have a length %d but got %d", amountOfColumns, len(row)))

			for _, cell := range row {
				assert.False(testing, cell.IsRevealed, "Expected IsRevealed to be false")
				assert.False(testing, cell.HasFlag, "Expected HasFlag to be false")
				assert.False(testing, cell.HasMine, "Expected HasMine to be false")
				assert.Equal(testing, 0, int(cell.AmountOfNeighboursWithMine), fmt.Sprintf("Expected AmountOfNeighboursWithMine to be 0 but got %d", cell.AmountOfNeighboursWithMine))
			}
		}
	})

	suite.Run("Invalid settings", func(suite *testing.T) {
		suite.Run("Single row", func(testing *testing.T) {
			newBoard, err := board.New(1, 2, 2)

			assert.Error(testing, err, "Expected error for board with single row")
			assert.Nil(testing, newBoard, "Expected board to be nil")
		})

		suite.Run("Single column", func(testing *testing.T) {
			newBoard, err := board.New(2, 1, 2)

			assert.Error(testing, err, "Expected error for board with single column")
			assert.Nil(testing, newBoard, "Expected board to be nil")
		})

		suite.Run("No mines", func(testing *testing.T) {
			newBoard, err := board.New(2, 2, 0)

			assert.Error(testing, err, "Expected error for board with no mines")
			assert.Nil(testing, newBoard, "Expected board to be nil")
		})

		suite.Run("Too many mines", func(testing *testing.T) {
			newBoard, err := board.New(2, 2, 4)

			assert.Error(testing, err, "Expected error for board with too many mines")
			assert.Nil(testing, newBoard, "Expected board to be nil")
		})
	})
}
