package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextId = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("Contact already exists: %+v\n", name)
		return
	}

	newContact := Contact{
		ID:    nextId,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	nextId++
	contactList = append(contactList, newContact)
	//contactIndexByName[name] = newContact.ID
	// The index is used in the custom type slice, so it should start from 0
	contactIndexByName[name] = len(contactList) - 1
	fmt.Printf("Contact added: %+v\n", name)
}

func findContact(name string) *Contact {
	// Why we don't use the map index by key,
	// because we need to get the nil when not found.
	// And we need to obtain the reference of item in the slice.

	// The index is used in the custom type slice, so it should start from 0
	index, exists := contactIndexByName[name]
	if exists {
		return &contactList[index] // used as a slice subscript
	}
	return nil
}

func listContacts() {
	fmt.Println("--- Listing contacts ---")
	if len(contactList) == 0 {
		fmt.Println("No contacts found")
		return
	}
	for i, contact := range contactList {
		fmt.Printf("%d. ID: %d, Name: %s, Email: %s, Phone: %s\n", i+1, contact.ID, contact.Name, contact.Email, contact.Phone)
	}

	fmt.Println("")
}

func main() {

	addContact("Alice Wonderland", "alice@example.com", "111-2222")
	addContact("Bob The Builder", "bob@example.com", "333-4444")
	addContact("Charlie Brown", "charlie@example.com", "555-6666")
	addContact("Alice Wonderland", "alice.new@example.com", "777-8888") // Attempt to duplicate

	listContacts()

	fmt.Printf("%+v\n", contactIndexByName)

	searchName := "Bob The Builder"
	searchIndex := findContact(searchName)
	if searchIndex == nil {
		fmt.Printf("%s not found\n", searchName)
	} else {
		fmt.Printf("%s found\n"+
			"ID: %d, Name: %s, Email: %s, Phone: %s\n", searchName, searchIndex.ID, searchIndex.Name, searchIndex.Email, searchIndex.Phone)

	}

	//searchIndex := contactIndexByName[searchName]
	//fmt.Printf("%+v\n", searchIndex)
}
