package main

import "fmt"

func main() {

	//Arrays
	a := map[string]int{"a": 1, "b": 2, "c": 3}

	fmt.Println(a)
	fmt.Println(a["a"])

	a["a"] = 10
	fmt.Println(a)
	fmt.Println("len: ", len(a)) //<--length

}
