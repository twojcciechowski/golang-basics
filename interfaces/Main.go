package main

import "fmt"

type MyTextMessage interface {
	String() string
}
type Email struct {
	Text string
}

func (e *Email) String() string {
	return e.Text
}
func PrintMessage(m MyTextMessage) {
	fmt.Printf("%s \n", m.String())
}
func main() {
	e := Email{Text: "hello"}
	PrintMessage(&e)
}
