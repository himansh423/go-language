package main
import "fmt"

func simpleFunction() {
	fmt.Println("This is a simple function");
}

func add(a int, b int) int {
	return a + b;
}
func main() {
	fmt.Println("I'm learning function in Go Language");
	simpleFunction();
	answer := add(5, 10);
	fmt.Println("The sum of 5 and 10 is: ", answer);
}