package main

import "fmt"

// custom struct
type contactInfo struct {
	email   string
	zipCode int
}

// struct definition
type person struct {
	firstName string
	lastName  string
	age       uint8
	// embedded structs
	contact contactInfo
}

func main() {

	// one way of declaration
	me := person{
		firstName: "Hilmi",
		lastName:  "Arikmert",
		age:       24,
		// assigning values to embedded structs
		contact: contactInfo{
			email:   "mharikmert@gmail.com",
			zipCode: 12121,
		},
	}

	var alex person

	// updating structs
	alex.firstName = "alex"
	alex.lastName = "anderson"
	alex.age = 35
	alex.contact = contactInfo{
		email:   "alex@gmail.com",
		zipCode: 33333,
	}

	fmt.Println(me)
	// fmt.Println(alex)

	// usage of printf
	// fmt.Printf("%+v", alex)

	fmt.Println("\nBEFORE UPDATE")
	alex.print()

	alex.updateFirstName("New firstName")

	fmt.Println("\nAFTER UPDATE")
	alex.print() // it's not supposed be changed. -> pointer issues

	alex.updateFirstNameWithPassByReference("Changed")
	fmt.Println("\nAFTER PASS BY REFERENCE UPDATE")
	alex.print()

	updateFirstName(&alex, "Changed2")
	alex.print()

	alexPointer := &alex
	alexPointer.originalWayOfUpdatingFirstName("Changed with pointer")
	alex.print()

	firstNamePointer := &alex.firstName
	lastNamePointer := &alex.lastName

	*firstNamePointer = "something new" // changes the firstName of the alex
	*lastNamePointer = "new lastName"   // changes the lastName of the alex

	alex.print()

	fmt.Println(*&firstNamePointer)

	// fmt.Println("First name pointer", firstNamePointer)
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) updateFirstName(firstName string) {
	p.firstName = firstName // pass by value
}

// pointer shortcut way of defining a receiver function
func (p *person) updateFirstNameWithPassByReference(firstName string) {
	p.firstName = firstName // pass by reference
}

// original way of updating the firstName of the struct person
func (pointer *person) originalWayOfUpdatingFirstName(firstName string) {
	(*pointer).firstName = firstName
}

// standard function call and pass by reference
func updateFirstName(p *person, firstName string) {
	p.firstName = firstName // pass by reference
}
