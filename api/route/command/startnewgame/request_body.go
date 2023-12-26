package startnewgame

type RequestBody struct {
	AmountOfRows    uint `json:"amountOfRows"`
	AmountOfColumns uint `json:"amountOfColumns"`
	AmountOfMines   uint `json:"amountOfMines"`
}
