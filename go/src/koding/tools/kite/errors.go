package kite

import (
	"reflect"
)

type ErrorObject struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func CreateErrorObject(err error) *ErrorObject {
	return &ErrorObject{Name: reflect.TypeOf(err).Elem().Name(), Message: err.Error()}
}

type InternalKiteError struct{}

func (err *InternalKiteError) Error() string {
	return "An internal error occurred in the Kite."
}

type UnknownMethodError struct {
	Method string
}

func (err *UnknownMethodError) Error() string {
	return "Method '" + err.Method + "' not known."
}

type ArgumentError struct {
	Expected string
}

func (err *ArgumentError) Error() string {
	return "Invalid argument, " + err.Expected + " expected."
}

type PermissionError struct{}

func (err *PermissionError) Error() string {
	return "Permission denied."
}

type WrongChannelError struct{}

func (err *WrongChannelError) Error() string {
	return "Wrong channel."
}

type Error struct {
	Message string
}

func (err *Error) Error() string {
	return err.Message
}
