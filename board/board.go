package board

import "github.com/go-playground/validator/v10"

type Board struct {
	AmountOfRows    uint `validate:"required,gte=2"`
	AmountOfColumns uint `validate:"required,gte=2"`
	AmountOfMines   uint `validate:"required,gte=1,isAmountOfMinesTwoLessThanBoardSize"`
	Cells           [][]Cell
}

func New(amountOfRows uint, amountOfColumns uint, amountOfMines uint) (*Board, error) {
	board := Board{
		AmountOfRows:    amountOfRows,
		AmountOfColumns: amountOfColumns,
		AmountOfMines:   amountOfMines,
	}

	err := board.Validate()

	if err != nil {
		return nil, err
	}

	board.populate()

	return &board, nil
}

func (board *Board) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("isAmountOfMinesTwoLessThanBoardSize", isAmountOfMinesTwoLessThanBoardSize)

	return validate.Struct(board)
}

func isAmountOfMinesTwoLessThanBoardSize(fieldValidator validator.FieldLevel) bool {
	amountOfMines := fieldValidator.Field().Uint()

	settingsValue := fieldValidator.Parent()
	amountOfColumns := settingsValue.FieldByName("AmountOfColumns").Uint()
	amountOfRows := settingsValue.FieldByName("AmountOfRows").Uint()

	validMaximumAmountOfMines := amountOfColumns*amountOfRows - 2

	return amountOfMines <= validMaximumAmountOfMines
}

func (board *Board) populate() {
	cells := make([][]Cell, board.AmountOfRows)

	for rowIndex := range cells {
		cells[rowIndex] = make([]Cell, board.AmountOfColumns)

		for columnIndex := range cells[rowIndex] {
			cells[rowIndex][columnIndex] = Cell{}
		}
	}

	board.Cells = cells
}
