package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	fmt.Println("We are learning JSON")
	person := Person{Name: "John", Age: 30, IsAdult: true}
	fmt.Println("Person Data is: ", person)

	// convert person into JSON Encoding (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling: ", err)
		return
	}

	fmt.Println("Person Data is: ", string(jsonData))

	// Decoding (unmarshalling)

	var personData Person

	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error unmarshalling: ", err)
		return
	}

	fmt.Println("Person Data is: ", personData)
}
