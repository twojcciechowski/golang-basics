package main

import "fmt"

func main() {

	x, s := f1()
	//x, s := f2()
	fmt.Println("x=", x)
	fmt.Println("s=", s)
}

func f1() (int, string) {

	defer func() (int, string) {
		return 100, "aa"
	}()

	w := 1
	t := "txt"
	return w, t
}

func f2() (w int, t string) {

	defer func() {
		w = 100
		t = "aa"
	}()

	w = 1
	t = "txt"
	return
}
