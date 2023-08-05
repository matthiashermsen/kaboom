package server

import "fmt"

type PortInvalidError struct{}

func (portInvalidError PortInvalidError) Error() string {
	return fmt.Sprintf("Invalid port. Port number should be between %d and %d (inclusive).", minimumPort, maximumPort)
}
