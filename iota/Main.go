package main

import "fmt"

const (
	A1 = iota
	A2 = iota
	A3 = iota
)

const (
	B1 = iota + 10
	B2
	B3
)

func main() {
	fmt.Printf("A1 %d \n", A1)
	fmt.Printf("A2 %d \n", A2)
	fmt.Printf("A3 %d \n", A3)

	fmt.Printf("B1 %d \n", B1)
	fmt.Printf("B2 %d \n", B2)
	fmt.Printf("B3 %d \n", B3)
}
