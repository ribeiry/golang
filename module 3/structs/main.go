package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	var firstPerson person

	firstPerson.firstName = "Augusto"
	firstPerson.lastName = "Ribeiro"
	firstPerson.contact.email = "teste@teste.com.br"
	firstPerson.contact.zipCode = "04890-185"

	firstPerson.print()
	fmt.Println()
	fmt.Println("==========================")
	firstPerson.updateName("José")

	firstPerson.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
