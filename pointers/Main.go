package main

import "fmt"

func updateIntA(number *int) {
	*number = *number * 2
}

func updateIntB(number *int) {
	b := 10
	number = &b
}

func updateIntC(number *int) *int {
	b := 10
	return &b
}

func main() {

	a := 1

	fmt.Printf("%d", a)

	//updateIntA(&a)
	//fmt.Printf("%d", a)

	//updateIntB(&a) //does not work why??
	//fmt.Printf("%d", a)

	//fmt.Printf("%d", *updateIntC(&a))

}
