package errors

func IsNotFound(err error) bool {
	switch typedErr := err.(type) {
	case ErrorWithCode:
		return typedErr.ErrorCode() == int(ErrorCodeNotFound)
	}
	return false
}
