package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,banana,mango"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two three four five two three"
	count := strings.Count(str, "two")
	fmt.Println(count)

	str1 := "Hello"
	str2 := "World"
	result := strings.Join([]string{str1, str2}, " ")
	fmt.Println(result)
}
