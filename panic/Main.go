package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

func main() {

	x, s := f1()
	fmt.Println("x=", x)
	fmt.Println("s=", s)
}

func f1() (int, string) {

	w, err := add(1, 2)
	if err != nil {
		println(err.Error())
	}
	t := "txt"
	return w, t
}

func add(x, y int) (z int, e error) {
	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered from : ", r)
			b := debug.Stack()
			fmt.Println(string(b))
			z = 0
			e = errors.New("index out of bound")
		}
	}()

	a := []int{1, 2, 3}
	println(a[4])

	return a[4], nil
}
