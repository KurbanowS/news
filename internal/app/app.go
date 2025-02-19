package app

import "strings"

type AppError struct {
	code    string
	key     string
	comment string
}

func NewAppError(code, key, comment string) *AppError {
	return &AppError{code, key, comment}
}

func (err AppError) Error() string {
	return strings.Trim(err.code+":"+err.key+" - "+err.comment, " -:")
}

func (err AppError) Code() string {
	return err.code
}

func (err AppError) Key() string {
	return err.key
}

func (err AppError) Comment() string {
	return err.comment
}

func (err *AppError) SetKey(key string) *AppError {
	err.key = key
	return err
}

func (err *AppError) SetComment(comment string) *AppError {
	err.comment = comment
	return err
}

var (
	ErrNotFound = NewAppError("not found", "", "")
	ErrRequired = NewAppError("required", "", "please fill required field")
	ErrInvalid  = NewAppError("invalid", "", "please fill with valid options")
)
