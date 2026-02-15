package main

import "fmt"

func main() {

	var imie string
	var nazwisko string
	var wiek int
	var status string

	fmt.Println("Podaj imię:")
	fmt.Scanln(&imie)
	fmt.Println("Podaj nazwisko:")
	fmt.Scanln(&nazwisko)

	for {
		fmt.Println("Podaj wiek:")
		_, err := fmt.Scanln(&wiek)
		if err != nil {
			fmt.Println("Nieprawidłowy wiek")
			continue
		}
		if wiek < 0 {
			fmt.Println("Wiek nie może być ujemny")
			continue
		}
		break

	}
	if wiek < 18 {
		status = "niepełnoletni"
	} else {
		status = "pełnoletni"
	}

	fmt.Println("Witaj", imie, nazwisko, "masz", wiek, "lat", "i jesteś", status)

}
