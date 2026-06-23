package main

import (
	"fmt"
	"time"
)

// Employee create the type
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

// NewEmployee the factory function
func NewEmployee(id int, firsName, lastName, position string, isActive bool) Employee {
	return Employee{
		ID:        id,
		FirstName: firsName,
		LastName:  lastName,
		Position:  position,
		IsActive:  isActive,
		JoinedAt:  time.Now(),
	}
}

func main() {

	jane := Employee{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
		Position:  "Night",
		Salary:    1000,
		IsActive:  true,
		JoinedAt:  time.Now(),
	}

	fmt.Printf("%+v\n", jane)
	fmt.Println(jane.ID)
	fmt.Println(jane.FirstName)
	fmt.Println(jane.LastName)
	fmt.Println(jane.Position)
	fmt.Println(jane.Salary)

	joe := NewEmployee(1, "Joe", "Lash", "Night", true)

	fmt.Println(joe.FirstName)
	fmt.Println(joe.LastName)
	fmt.Println(joe.Position)

	joe.Salary = 10000
	fmt.Println(joe.Salary)

	joePtr := &joe
	joePtr.IsActive = false // dereference automatically (recommended)
	(*joePtr).LastName = "Lash raj"
	fmt.Println(joe)
}
