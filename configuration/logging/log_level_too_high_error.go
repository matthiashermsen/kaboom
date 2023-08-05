package logging

import (
	"fmt"

	"github.com/rs/zerolog"
)

type LogLevelTooHighError struct{}

func (logLevelTooHighError LogLevelTooHighError) Error() string {
	return fmt.Sprintf("Log level is too high. Maximum log level is %d.", zerolog.PanicLevel)
}
