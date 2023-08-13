package response

type ApiResponse[T any] struct {
	Status string           `json:"status"`
	Data   T                `json:"data"`
	Error  ApiResponseError `json:"error"`
}

type ApiResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewSuccessApiResponse[T any](data T) ApiResponse[T] {
	return ApiResponse[T]{
		Status: "success",
		Data:   data,
		Error:  ApiResponseError{},
	}
}

func NewFailureApiResponse(errorCode string, errorMessage string) ApiResponse[interface{}] {
	return ApiResponse[interface{}]{
		Status: "failure",
		Data:   nil,
		Error: ApiResponseError{
			Code:    errorCode,
			Message: errorMessage,
		},
	}
}

func NewErrorApiResponse() ApiResponse[interface{}] {
	return ApiResponse[interface{}]{
		Status: "error",
		Data:   nil,
		Error: ApiResponseError{
			Code:    InternalError,
			Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
		},
	}
}
