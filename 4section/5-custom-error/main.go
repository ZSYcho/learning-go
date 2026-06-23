package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrDivisionByZero = errors.New("division by zero")
var ErrNumTooLarge = errors.New("number too large")

type OpError struct {
	Op      string
	Code    int
	Message string
	Time    time.Time
}

func (op OpError) Error() string {
	return op.Message
}

func NewOpError(op string, code int, message string, t time.Time) *OpError {
	return &OpError{op, code, message, t}
}

func DoSomething() error {
	return NewOpError("doSomething", 100, "do something failed!", time.Now())
}
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}

	if a > 1000 {
		return 0, ErrNumTooLarge
	}
	return a / b, nil
}

func main() {

	value, err := divide(10001, 1)
	if err != nil {
		if errors.Is(err, ErrDivisionByZero) {
			fmt.Println("divide by zero")
		} else if errors.Is(err, ErrNumTooLarge) {
			fmt.Println("number too big")
			fmt.Println(err)
		}
		return
	}

	fmt.Println(value)
}
