package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num + 5

}

func main() {
	// var num int
	// num = 2

	// var ptr *int
	// ptr = &num

	num := 2
	ptr := &num
	fmt.Println("Value of num is", num)
	fmt.Println("Address of num is", &num)
	fmt.Println("Value of num is", *ptr)
	fmt.Println("Address of num is", ptr)

	value := 5
	modifyValueByReference(&value)
	fmt.Println("Value of value is", value)
}
