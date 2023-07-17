package utils

type CustomError struct {
	HTTPStatus int
	Message    interface{}
	Err        error
}

func NewCustomError(httpStatus int, message interface{}, err error) *CustomError {
	return &CustomError{
		HTTPStatus: httpStatus,
		Message:    message,
		Err:        err,
	}
}

func (err *CustomError) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	if message, messageIsString := err.Message.(string); messageIsString {
		return message
	}

	return "Unknown error"
}
