package main

import (
	"errors"
	"fmt"
)

var ERR_WRONG error = errors.New("something wrong")

func main() {

	if err := func1(); errors.Is(err, ERR_WRONG) {
		fmt.Printf("func1() nie działa bo: %s \n", err)
	}

	//if err := func3(); errors.Is(err, ERR_WRONG) {
	//	fmt.Printf("%v \n", err.Error())
	//
	//	a := errors.Unwrap(err)
	//	fmt.Printf("%v \n", a.Error())
	//
	//	b := errors.Unwrap(a)
	//	fmt.Printf("%v \n", b.Error())
	//}

}

func func3() error {

	if err := func2(); err != nil {
		return fmt.Errorf("func2() does not work: %w", err)
	}

	return nil
}

func func2() error {

	if err := func1(); err != nil {
		return fmt.Errorf("func1() does not work: %w", err)
	}

	return nil
}

func func1() error {
	return ERR_WRONG
}
