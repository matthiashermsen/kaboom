package server

func ValidateConfiguration(configuration Configuration) error {
	if configuration.Port < minimumPort || configuration.Port > maximumPort {
		return PortInvalidError{}
	}

	return nil
}
