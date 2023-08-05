package logging

import "github.com/rs/zerolog"

func ValidateConfiguration(configuration Configuration) error {
	if configuration.Level < zerolog.TraceLevel {
		return LogLevelTooLowError{}
	}

	if configuration.Level > zerolog.PanicLevel {
		return LogLevelTooHighError{}
	}

	return nil
}
