package main

import (
	"errors"
	"fmt"
)

func main() {

	var err = function3()

	fmt.Println(fmt.Sprintf("error %v", err))

}

func function3() (err error) {

	defer func() {
		if r := recover(); r != nil {
			err = errors.New("some error occurred")
		}
	}()

	panic("panic")

	return err
}
