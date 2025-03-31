package main

import (
	"fmt"
)

func main() {
	// fmt.Println("For loop")

	// for i:=0; i < 5; i++ {
	// 	fmt.Println("Numbers",i)
	// }

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, value := range numbers {
		fmt.Println("Index", index, "Value", value)
	}

	data := "Hello World"
	for index, value := range data {
		fmt.Println("Index", index, "Value", string(value))
	}
}
