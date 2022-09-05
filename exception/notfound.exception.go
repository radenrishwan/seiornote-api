package exception

type NotFoundException struct {
	Message string
}

func NewNotFoundException(message string) ValidateException {
	return ValidateException{Message: message}
}

func (exception NotFoundException) Error() string {
	return exception.Message
}
