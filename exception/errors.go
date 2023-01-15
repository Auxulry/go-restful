// Package exception is describe all exception errors
package exception

type Exception struct {
	message string
}

func NewException(message string) *Exception {
	return &Exception{message: message}
}

func (e *Exception) Error() string {
	return e.message
}
