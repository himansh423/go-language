package main

import "fmt"

func main() {
	studentGrades := make(map[string]int)
	studentGrades["John"] = 92
	studentGrades["Jane"] = 96
	studentGrades["Bob"] = 88
	studentGrades["Alice"] = 78

	fmt.Println(("Marks of John"), studentGrades["John"])
	studentGrades["John"] = 95
	fmt.Println(("New Marks of John"), studentGrades["John"])

	delete(studentGrades, "John")
	grades, exist := studentGrades["John"]
	fmt.Println(("Grades of John"), grades, exist)
	fmt.Println(("Marks of John"), studentGrades["John"])

	for index, value := range studentGrades {
		fmt.Println("Index", index, "Value", value)
	}

	person := map[string]string{
		"alice": "123-456-7890",
		"bob":   "123-456-7891",
		"john":  "123-456-7892",
		"jane":  "123-456-7893",
	}

	for index, value := range person {
		fmt.Println("Index", index, "Value", value)
	}

}
