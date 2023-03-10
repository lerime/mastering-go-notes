package main

import (
	"errors"
	"fmt"
)

func main() {
	//defining new error
	err := defineNewError()
	if err != nil {
		fmt.Println(err)
	}

	//using Errorf
	err = getErrorWithErrorf("ne hatasi bu!")
	if err != nil {
		fmt.Println(err)
	}
}

func defineNewError() error {
	return errors.New("new error")
}

func getErrorWithErrorf(message string) error {
	return fmt.Errorf(message)
}
