package main

import "fmt"

func main() {
	// fmt.Println("I'm learning slice in Go Language");
	// numbers := []int{1, 2, 3, 4, 5}
	// number := append(numbers, 6)
	// fmt.Println("Numbers: ", numbers)
	// fmt.Printf("Number has data type: %T\n", numbers)

	numbers := make([]int, 3, 5)
	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
}