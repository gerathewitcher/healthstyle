package errors

type recordNotFound struct {
	message string
}

func (e *recordNotFound) Error() string {
	return e.message
}

var RecordNotFound = &recordNotFound{
	message: "record not found",
}
