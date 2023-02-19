package main

import "fmt"

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func main() {
	p := Person{Name: "Samiul"}
	p.talk()
	a := new(Android)
	a.Name = "Rocks"
	a.talk()
}

func (p *Person) talk() {
	fmt.Println("Hi, my name is", p.Name)
}
