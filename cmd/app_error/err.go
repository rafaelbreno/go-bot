package app_error

import (
	"fmt"
	"log"
)

type AppError struct {
	Err error
	Ctx string
}

func (e *AppError) Log() {
	log.Fatalf(`Err: %s - %s`, e.Message(), e.Context())
}

func (e *AppError) Context() string {
	return e.Ctx
}

func (e *AppError) Error() error {
	return e.Err
}

func (e *AppError) Message() string {
	return e.Err.Error()
}

func NewError(message, context string) {

	e := AppError{
		Err: fmt.Errorf(message),
		Ctx: context,
	}

	e.Log()
}
