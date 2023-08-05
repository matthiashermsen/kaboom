package logging

import (
	"fmt"

	"github.com/rs/zerolog"
)

type LogLevelTooLowError struct{}

func (logLevelTooLowError LogLevelTooLowError) Error() string {
	return fmt.Sprintf("Log level is too low. Minimum log level is %d.", zerolog.TraceLevel)
}
