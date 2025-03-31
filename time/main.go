package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time: ", currentTime)
	fmt.Printf("Type of CurrentTime %T\n", currentTime)

	formatted := currentTime.Format("02-01-2006, Monday")
	fmt.Println("Formatted Time: ", formatted)

	layout_string := "2006-01-02"
	date_str := "2023-11-25"
	formatted_time, _ := time.Parse(layout_string, date_str)
	fmt.Println("formatted data: ", formatted_time)
}
