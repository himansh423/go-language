package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	userId    int    `json:"userId"`
	id        int    `json"id"`
	title     string `json:"title"`
	completed bool   `json:"completed"`
}

func performPostRequest() {
	todo := Todo{
		userId:    23,
		title:     "Prince Bhai hai kya",
		completed: true,
		id:        23,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling: ", err)
		return
	}

	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

	res, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", jsonReader)

	if err != nil {
		fmt.Println("Error sending Request: ", err)
		return
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)

	fmt.Println("Respone is: ", string(data))

}
func main() {
	performPostRequest()
}
