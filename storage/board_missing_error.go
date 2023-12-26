package storage

type BoardMissingError struct{}

func (err *BoardMissingError) Error() string {
	return "Board is missing."
}
