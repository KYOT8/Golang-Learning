package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Witaj w sciecie wycinkow ")

	var listaowocow = []string{"Jabłko", "Banan", "Gruszka", "Pomarańcza"}
	fmt.Printf("typ listy owocow to %T\n", listaowocow)

	listaowocow = append(listaowocow, "Mango", "Ananas") //wykorzystujemy funkcje append do dodania nowych elementow do listy
	fmt.Println(listaowocow)

	listaowocow = append(listaowocow[1:3]) //daje to wycinek z wycinka od zakresu 1-3 indeksu
	fmt.Println(listaowocow)               //1-3 to pierwszy sie nie liczy (bo to zero 0,1,2,3,4), czyli zostaje 1 i 2, a 3 sie nie liczy bo to 4 numer tak naprawde jest 4 elementem, a nie 3

	highScores := make([]int, 4) //tworzymy slice o nazwie highScore, ktory jest slice'em typu int i ma 5 elementow

	highScores[0] = 100
	highScores[1] = 900
	highScores[2] = 300
	highScores[3] = 400
	//highScores[4] = 300

	highScores = append(highScores, 500, 666, 321) //dodajemy nowy element do slice'a highScores

	fmt.Println(highScores)

	sort.Ints(highScores) //sortujemy slice highScores w porzadku rosnacym, funkcja sort.Ints sortuje slice typu int
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) //daje nam to czy slice jest posortowany, zwraca true lub false

}
