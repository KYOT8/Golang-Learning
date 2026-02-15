package main

import "fmt"

func main() {
	var a, b float64
	var dzialanie string

	fmt.Println("Podaj pierwszą liczbe")
	fmt.Scanln(&a)

	fmt.Println("Podaj drugą liczbe")
	fmt.Scanln(&b)

	fmt.Println("Podaj działanie" + " (+, -, *, /)")
	fmt.Scanln(&dzialanie)
	switch dzialanie {
	case "+":
		fmt.Printf("Wynik: %f", a+b)
	case "-":
		fmt.Printf("Wynik: %f", a-b)
	case "*":
		fmt.Printf("Wynik: %f", a*b)
	case "/":
		if b != 0 {
			fmt.Printf("Wynik: %f", a/b)
		} else {
			fmt.Println("Nie można dzielić przez zero!")
		}
	default:
		fmt.Println("Nieznane działanie!")
	}
}
