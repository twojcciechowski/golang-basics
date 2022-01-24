package main

import "fmt"

func main() {

	//Arrays
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //array

	//fmt.Println("----- example arrays 1 -----")
	//
	//fmt.Println(a)
	//fmt.Println("cap: ", cap(a)) //<--capacity
	//fmt.Println("len: ", len(a)) //<--length
	//
	//fmt.Println("----- example arrays 2 -----")
	//a = append(a, 10)
	//
	//fmt.Println(a)
	//fmt.Println("cap: ", cap(a)) //<--capacity
	//fmt.Println("len: ", len(a)) //<--length

	//Slices

	b := a[:2]

	fmt.Println("----- example slices 1 -----")
	fmt.Println(b)
	fmt.Println("cap: ", cap(b)) //<--capacity
	fmt.Println("len: ", len(b)) //<--length
	//fmt.Println("#3: ", b[3])

	fmt.Println("----- example slices 2 -----")

	c := a[4:]
	fmt.Println(c)
	fmt.Println("cap: ", cap(c)) //<--capacity
	fmt.Println("len: ", len(c)) //<--length

	fmt.Println("----- example slices 3 -----")
	c = a[4:6]
	fmt.Println(c)
	fmt.Println("cap: ", cap(c)) //<--capacity
	fmt.Println("len: ", len(c)) //<--length

	fmt.Println("----- example slices 4 -----")
	c = a[4:6:6]
	fmt.Println(c)
	fmt.Println("cap: ", cap(c)) //<--capacity
	fmt.Println("len: ", len(c)) //<--length

	for _, v := range b {

		fmt.Printf("index: %d, value: %v \n", 0, v)
	}
}
