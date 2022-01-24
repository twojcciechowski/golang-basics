package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	x, s := f1()
	fmt.Println("x=", x)
	fmt.Println("s=", s)
}

func f1() (int, string) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from : ", r)
			b := debug.Stack()
			fmt.Println(string(b))
		}
	}()

	w := add(1, 2)
	t := "txt"
	return w, t
}

func add(x, y int) int {
	panic("panic error")
}
