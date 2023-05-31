package main

import (
	"errors"
	"fmt"
)

func main() {
	err := function()
	if err != nil {
		fmt.Printf("not finished")
	} else {
		fmt.Printf("finished")
	}
}

func function() error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("some error occurred")
		}
	}()
	panic("panic")
	return err
}
