package main

import "fmt"

type Person interface {
	GetName() string
}

type Employee struct {
	ID   int
	Name string
}

type BusinessPerson struct {
	ID   int
	Name string
}

func (b BusinessPerson) GetName() string {
	return b.Name
}

func (e Employee) GetName() string {
	return e.Name
}

func displayPerson(p Person) {
	fmt.Println(p.GetName())
}

func main() {

	kevin := Employee{
		ID:   1,
		Name: "Kevin",
	}

	jane := BusinessPerson{
		ID:   1,
		Name: "Jane",
	}

	displayPerson(kevin)
	displayPerson(jane)
}
