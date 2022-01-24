package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

func sum(x, y int) (z int) {
	z = x + y
	return
}

func sqrt(x int) (float64, error) {
	if x <= 0 {
		return 0, fmt.Errorf("incorrect number")
	}
	r := math.Sqrt(float64(x))
	return r, nil
}

func main() {

	//add
	//r := add(2, 3)
	//fmt.Printf("result = %d", r)

	//sum
	r := sum(2, 3)
	fmt.Printf("result = %d", r)

	//sqrt
	//r, err := sqrt(4)
	//
	//if err != nil {
	//	fmt.Printf("%s", err)
	//}
	//
	//fmt.Printf("result = %0.3f", r)
}
