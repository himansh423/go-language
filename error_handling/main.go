package main
import "fmt"

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
func main() {
	fmt.Println("I'm learning error handling in Go Language");

  answer, _ := divide(10, 1)
	fmt.Println("The division of 10 by 1 is: ", answer);
}