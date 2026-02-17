package main

import "fmt"

func main() {

	var zamowienie string
	var znaleziono bool

	cenaPizzy := map[string]float64{
		"Pepperoni":   30.0,
		"Margherita":  20.0,
		"Hawajska":    25.0,
		"Cztery sery": 35.0,
		"Kaprys":      28.0,
		"Romana":      22.0,
		"Weganska":    27.0,
		"Salami":      32.0,
		"Meksykanska": 29.0,
		"Grecka":      26.0,
	}

	wybor_pizzy := []string{"Pepperoni", "Margherita", "Hawajska", "Cztery-Sery", "Kaprys", "Romana", "Weganska", "Salami", "Meksykanska", "Grecka"}

	fmt.Println("Menu:")
	fmt.Println("wybor pizzy: ", wybor_pizzy)

	fmt.Println("Podaj nazwe pizzy ktora chcesz zamowic:")
	fmt.Scanln(&zamowienie)
	for {
		znaleziono = false
		for _, pizza := range wybor_pizzy {
			if pizza == zamowienie {
				znaleziono = true
			}
		}
		if znaleziono {
			break
		} else {
			fmt.Println("Nie mamy takiej pizzy")
			fmt.Println("Podaj ponownie nazwe pizzy ktora chcesz zamowic:")
			fmt.Scanln(&zamowienie)
			znaleziono = false
		}
	}

	cenanapoju := map[string]float64{
		"Woda":     5.0,
		"Sok":      7.0,
		"Cola":     6.0,
		"Fanta":    6.0,
		"Sprite":   6.0,
		"Pepsi":    6.0,
		"7up":      6.0,
		"Lipton":   5.0,
		"Nestea":   5.0,
		"Red Bull": 10.0,
	}
	var napoj string
	wybor_napojow := []string{"Woda", "Sok", "Cola", "Fanta", "Sprite", "Pepsi", "7up", "Lipton", "Nestea", "Red Bull"}
	fmt.Println("Dodatkowo do pizzy mozesz zamowic napoj?")
	fmt.Println("Tak/Nie")
	var odpowiedz string
	fmt.Scanln(&odpowiedz)

	if odpowiedz == "Nie" {
		napoj = "Brak napoju"
	} else if odpowiedz == "Tak" {
		fmt.Println("Proponuje napoj")
		fmt.Println("wybor napojow: ", wybor_napojow)
		fmt.Println("Podaj nazwe napoju ktory chcesz zamowic:")
		fmt.Scanln(&napoj)
		for {
			znaleziono = false
			for _, napoj_w_menu := range wybor_napojow {
				if napoj_w_menu == napoj {
					znaleziono = true
				}
			}
			if znaleziono {
				break
			} else {
				fmt.Println("Nie mamy takiego napoju")
				fmt.Println("Podaj ponownie nazwe napoju ktory chcesz zamowic:")
				fmt.Scanln(&napoj)
				znaleziono = false
			}

		}
	}
	var cena_zamowienia float64
	cena_zamowienia = cenaPizzy[zamowienie] + cenanapoju[napoj]
	if cena_zamowienia > 30.0 {
		fmt.Println("Zamowienie powyzej 30zl, otrzymujesz rabat 10%")
		cena_zamowienia = cena_zamowienia * 0.9
	}
	fmt.Printf("Cena zamowienia: %.2f\n", cena_zamowienia)
}
