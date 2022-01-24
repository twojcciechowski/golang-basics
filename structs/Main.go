package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Age       int
	firstname string
	lastname  string
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s % s", p.firstname, p.lastname)
}

func main() {

	p := Person{
		Age:       10,
		firstname: "Tom",
	}

	fmt.Println(p)

	fmt.Println("type:", reflect.TypeOf(p))

	fmt.Printf("%v \n", p)  //print struct
	fmt.Printf("%+v \n", p) //print struct with field names

	fmt.Printf("age: %d \n", p.Age)
	fmt.Printf("name: %s \n", p.firstname)

	fmt.Println("-------")
	k := new(Person) //<-- creates pointer to the struct
	k.Age = 1
	k.firstname = "Adam"

	fmt.Println(k)
	fmt.Println("type:", reflect.TypeOf(k))
	fmt.Printf("%v \n", k)  //print struct
	fmt.Printf("%+v \n", k) //print struct with field names

	fmt.Printf("age: %d \n", k.Age)
	fmt.Printf("name: %s \n", k.firstname)

	fmt.Println("-------")
	k.lastname = "Cos"
	fmt.Printf("name: %s \n", k.FullName())
}
