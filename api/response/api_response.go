package response

type APIResponse[T any] struct {
	Status string            `json:"status"`
	Data   T                 `json:"data"`
	Error  *APIResponseError `json:"error"`
}

type APIResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewSuccessAPIResponse[T any](data T) APIResponse[T] {
	return APIResponse[T]{
		Status: "success",
		Data:   data,
		Error:  nil,
	}
}

func NewFailureAPIResponse(errorCode string, errorMessage string) APIResponse[any] {
	return APIResponse[any]{
		Status: "failure",
		Data:   nil,
		Error: &APIResponseError{
			Code:    errorCode,
			Message: errorMessage,
		},
	}
}

func NewErrorAPIResponse() APIResponse[any] {
	return APIResponse[any]{
		Status: "error",
		Data:   nil,
		Error: &APIResponseError{
			Code:    InternalError,
			Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
		},
	}
}
