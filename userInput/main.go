package main
import "fmt"
import "bufio"
import "os"
func main() {	
	fmt.Println("Enter your name: ")
	// var name string	
	// fmt.Scanln(&name)
	// fmt.Println("Hello, ", name)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello, ", name)
}