package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web Service")
	res, err := http.Get("https://jsonplaceholder.typicode.com/users/1")

	if err != nil {
		fmt.Println("Error getting GET Response: ", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("type of response: %T\n", res)
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error reading response: ", err)
		return
	}

	fmt.Println("Response: ", string(data))

}
