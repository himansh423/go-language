package main
import "fmt"
func main() {
	fmt.Println("I'm learning array in Go Language");
	var name[5]string
	name[0] = "John"	
	name[1] = "Doe"
	name[2] = "Himanshu"
	name[3] = "Rahul"
	name[4] = "Raj"

	fmt.Println("Name: ", name)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers: ", numbers)

	fmt.Println("Length of name array: ", len(name))
	fmt.Println("Length of numbers array: ", len(numbers))

	fmt.Println("First element of name array: ", name[0])
}