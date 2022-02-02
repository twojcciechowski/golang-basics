package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func updateIntA(number *int) {
	*number = *number * 2
}

func updateIntB(number *int) {

	println("address of number: ", &number)
	b := 10
	number = &b

}

func updateIntC(number *int) *int {
	b := 10
	return &b
}

func checkTypeA() {
	var a byte
	fmt.Printf("original value : %d \n", a)

	a = 'a'
	b := "s"
	fmt.Printf("type of a: %s \n", reflect.TypeOf(a))
	fmt.Printf("type of b: %s \n", reflect.TypeOf(b))

}

func checkTypeB() {
	var a int
	a = 1

	fmt.Printf("original value : %d \n", a)

	s := strconv.Itoa(a)
	fmt.Printf("modified value : %s \n", s)
	fmt.Printf("modified type : %s \n", reflect.TypeOf(s))

}

func checkTypeC() {
	val, _ := strconv.Atoi("9999")

	fmt.Printf("modified value : %d \n", val)
	fmt.Printf("modified value : %s \n", reflect.TypeOf(val))
}

func main() {

	checkTypeA()
	checkTypeB()
	checkTypeC()

	a := 1

	fmt.Printf("original value : %d \n", a)

	updateIntA(&a)
	fmt.Printf("return value: %d \n", a)

	pa := &a
	println("address of a", &pa)

	updateIntB(&a) //does not work why??
	fmt.Printf("return value: %d", a)

	fmt.Printf("return value: %d", *updateIntC(&a))

}
