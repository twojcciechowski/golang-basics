package structs

import "fmt"

type Person struct {
	age       int
	Firstname string
	Lastname  string
	pesel     string
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.Firstname, p.Lastname)
}

func NewPerson(age int, firstname, lastname, pesel string) Person {
	return Person{
		age:       age,
		Firstname: firstname,
		Lastname:  lastname,
		pesel:     pesel,
	}
}
