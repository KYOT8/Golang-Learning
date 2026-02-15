package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func subtract(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func divide(x int, y int) int {
	return x / y
}

func main() {
	x := 10
	y := 5

	fmt.Println("Dodawanie:", add(x, y))
	fmt.Println("Odejmowanie:", subtract(x, y))
	fmt.Println("Mnożenie:", multiply(x, y))
	fmt.Println("Dzielenie:", divide(x, y))
}
