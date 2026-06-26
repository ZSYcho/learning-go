package main

import "fmt"

// In OOP language like Java Python JS PHP, we use inherent
// In Go we use composition instead
// They are two different concepts

// Composition --> Has-A relationship
// Inheritance --> Is-A relationship
// Car -> is composed of many parts (Engine, Doors)

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

func (a Address) FullAddress() string {
	if a.Street == "" && a.City == "" { // Handle zero value
		return "No address provided"
	}

	return fmt.Sprintf("%s, %s, %s %s", a.Street, a.City, a.State, a.ZipCode)
}

type Customer struct {
	CustomerID      int
	Name            string
	Email           string
	BillingAddress  Address // embedded
	ShippingAddress Address // embedded
}

func (c Customer) PrintDetails() {
	fmt.Printf("CustomerID: %d\n", c.CustomerID)
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Email: %s\n", c.Email)
	fmt.Printf("BillingAddress: %s\n", c.BillingAddress.FullAddress())   // Accessing method of composed type
	fmt.Printf("ShippingAddress: %s\n", c.ShippingAddress.FullAddress()) // Accessing method of composed type
}

func main() {

	fmt.Println("----- Composition -----")

	cust1 := Customer{
		CustomerID: 1001,
		Name:       "John Smith",
		Email:      "sales@johnsmith.com",
		BillingAddress: Address{
			Street:  "123 street",
			City:    "Innovateville",
			State:   "CA",
			ZipCode: "90210",
		},
		ShippingAddress: Address{
			Street:  "456 Factory Lane",
			City:    "Manufacturicity",
			State:   "NV",
			ZipCode: "89101",
		},
	}

	cust1.PrintDetails()

	fmt.Println("------ customer with same billing and shipping address -----")
	mainAddress := Address{
		Street:  "789 Main St",
		City:    "Anytown",
		State:   "TX",
		ZipCode: "75001",
	}
	cust2 := Customer{
		CustomerID:      1002,
		Name:            "Ray Dalio",
		Email:           "ray.dalio@email.com",
		BillingAddress:  mainAddress,
		ShippingAddress: mainAddress, // Reusing the existing address
	}

	cust2.PrintDetails()
}
