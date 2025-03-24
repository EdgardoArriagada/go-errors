package nerr

type temporaryError struct {
	message string
}

func (e *temporaryError) Error() string {
	return e.message
}

var TemporaryError = &temporaryError{message: "[Temporary Error]"}
