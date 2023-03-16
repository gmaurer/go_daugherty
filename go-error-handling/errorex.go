package main

import (
	"errors"
	"fmt"
	"log"
)

type InterestingError struct {
	Count int
}

func (e *InterestingError) Error() string {
	return fmt.Sprintf("It happened %d times.", e.Count)
}

func NewInterestingError(count int) error {
	return &InterestingError{Count: count}
}

func returnsError() error {
	return NewInterestingError(10)
}

func ReceivesError() {
	err := returnsError()
	if err != nil {
		if v, ok := err.(*InterestingError); ok {
			log.Println(v.Count)
		} else {
			log.Println(v.Error())
		}
	}
}

func WrappedCount(count int) error {
	err := CheckCount(count)
	if err != nil {
		wrappedError := fmt.Errorf("there was an error: %w", err)
		return wrappedError
	}
	return err
}

func CheckCount(count int) error {
	if count > 100 {
		return errors.New("count is too high")
	}

	return nil
}

func GetResponse(count int) (string, error) {
	if count > 100 {
		return "", errors.New("count is too high")
	}

	return fmt.Sprintf("The count was: %d", count), nil
}

func startToPanic() {
	panic("I am freaking out!")
}

func SimplePanicAndRecoverExample() {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Printf("I have recovered from the panic Error: %v\n", rec)
		}
	}()

	startToPanic()

	//This will be unreachable due to the panic in startToPanic()
	fmt.Println("After call to Panic")
}
