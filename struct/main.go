package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var prince Person
	prince.firstName = "Prince"
	prince.lastName = "Kumar"
	prince.age = 25
	fmt.Println("First Name:", prince.firstName)
	fmt.Println("Last Name:", prince.lastName)
	fmt.Println("Age:", prince.age)

	person1 := Person{
		firstName: "Alice",
		lastName:  "Doe",
		age:       25,
	}
	fmt.Println("First Name:", person1.firstName)
	fmt.Println("Last Name:", person1.lastName)
	fmt.Println("Age:", person1.age)

	var person2 = new(Person)
	person2.firstName = "Bob"
	person2.lastName = "Doe"
	person2.age = 30
	fmt.Println("First Name:", person2)

}
