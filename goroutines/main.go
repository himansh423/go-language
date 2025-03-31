package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello world...")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("sayHello Function ended..")
}

func sayHi() {
	fmt.Println("Hiii world...")
}
func main() {
	fmt.Println("Learning goroutines....")
	go sayHello()
	sayHi()
}
