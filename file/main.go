package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error while creating file: ", err)
	// 	return
	// }
	// defer file.Close()

	// content := "Hello World by Himanshu Chauhan"

	// _, errors := io.WriteString(file, content)
	// if err != nil {
	// 	fmt.Println("Error while writing file: ", errors)
	// 	return
	// }

	// fmt.Println("File Succesfully created")

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return
	}
	defer file.Close()
	buffer := make([]byte, 1024)

	for {
		n, error := file.Read(buffer)
		if error == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading file: ", err)
		}

		fmt.Println(string(buffer[:n]))
	}
}
