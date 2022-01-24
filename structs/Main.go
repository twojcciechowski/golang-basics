package main

import (
	structs "GoTraining/structs/xxxaaaa"
	"fmt"
	"reflect"
)

func main() {

	p := structs.NewPerson(10, "tom", "", "111")

	fmt.Println(p)

	fmt.Println("type:", reflect.TypeOf(p))

	fmt.Printf("%v \n", p)  //print struct
	fmt.Printf("%+v \n", p) //print struct with field names

	//fmt.Printf("age: %d \n", p.age)
	fmt.Printf("name: %s \n", p.Firstname)

	fmt.Println("-------")

	p.Lastname = "Cos"
	fmt.Printf("name: %s \n", p.FullName())

	fmt.Println("-------")

	//k := new(structs.Person) //<-- creates pointer to the struct
	//k.age = 1
	//k.Firstname = "Adam"
	//
	//fmt.Println(k)
	//fmt.Println("type:", reflect.TypeOf(k))
	//fmt.Printf("%v \n", k)  //print struct
	//fmt.Printf("%+v \n", k) //print struct with field names
	//
	//fmt.Printf("age: %d \n", k.age)
	//fmt.Printf("name: %s \n", k.Firstname)
	//
	//fmt.Println("-------")
	//k.Lastname = "Cos"
	//fmt.Printf("name: %s \n", k.FullName())
	//
	//structs.CheckPesel(k, "aaa")

}
