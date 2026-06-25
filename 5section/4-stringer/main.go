package main

import "fmt"

type Person interface {
	GetName() string
}

type BusinessPerson struct {
	ID   int
	Name string
}

func (b BusinessPerson) GetName() string {
	return b.Name
}

func (b BusinessPerson) String() string {
	return fmt.Sprintf("Person[ID:%d, Name:%s]", b.ID, b.Name)
}

type ID int

func (idx ID) String() string {
	return fmt.Sprintf("COMING FROM HERE ID[%d]", idx)
}

func main() {

	jane := BusinessPerson{
		ID:   1,
		Name: "Jane",
	}

	fmt.Println(jane)

	var myId ID
	myId = 715
	fmt.Println(myId)

}
