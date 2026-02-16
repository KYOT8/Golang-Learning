package main

import "fmt"

func main() {

	welcome := "Podaj ocene z doreczenia przesylki"
	fmt.Println(welcome)
	var ocena float64
	fmt.Scanln(&ocena)

	for {
		if ocena > 5.0 || ocena < 0.0 {
			fmt.Println("Nieprawidlowa ocena, ocena musi byc w przedziale od 0.0 do 5.0")

			fmt.Println("Podaj ponownie ocene")
			fmt.Scanln(&ocena)
			continue
		} else {
			if ocena >= 4.50 && ocena <= 5.0 {
				fmt.Println("Dziekujemy za wspaniala ocene")
				break
			} else if ocena < 4.50 && ocena >= 2.0 {
				fmt.Println("Przykro nam ze nie jestes wystarczajacyo zadowolony z naszej uslugi")
				break
			} else if ocena < 2.0 && ocena >= 0.0 {
				fmt.Println("Przykro nam z niewystarczajacej oceny, postaramy sie poprawic nasza usluge")
				break
			}
		}
	}
}
