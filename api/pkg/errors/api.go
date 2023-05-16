package errors

type ApiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewUnknownApiError(err error) *ApiError {
	return &ApiError{Code: "UNKNOWN", Message: err.Error()}
}
