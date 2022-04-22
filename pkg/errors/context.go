package errors

type errorContext struct {
	Field   string
	Message string
}

// GetType returns the error type
func GetType(err error) ErrorType {
	if customErr, ok := err.(CustomError); ok {
		return customErr.errType
	}

	return Error
}
