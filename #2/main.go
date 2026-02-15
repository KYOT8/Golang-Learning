package main

import "fmt"

func main() {
	var username string = "Maciej"
	fmt.Println(username)
	fmt.Printf("Variable type: %T\n", username)

	var isLogged bool = false
	fmt.Println(isLogged)
	fmt.Printf("Variable type: %T\n", isLogged)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable type: %T\n", smallVal)

	var smallFloat float64 = 255.45544511254451885
	fmt.Println(smallFloat)
	fmt.Printf("Variable type: %T\n", smallFloat)
}
