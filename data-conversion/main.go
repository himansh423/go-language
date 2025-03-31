package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Println("Value of num is", num)
	fmt.Printf("Type of num is %T\n", num)

	var data float64 = float64(num)
	fmt.Println("Value of data is", data)
	fmt.Printf("Type of data is %T\n", data)

	num = 123
	str := strconv.Itoa(num)
	fmt.Println("Value of str is", str)

	number_string := "123"
	number, _ := strconv.Atoi(number_string)

	fmt.Println("Value of number is", number)
	fmt.Printf("Type of number is %T\n", number)

	num_string := "123.456"
	num_float, _ := strconv.ParseFloat(num_string, 64)
	fmt.Println("Value of num_float is", num_float)
	fmt.Printf("Type of num_float is %T\n", num_float)

}
