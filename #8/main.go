package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var listaowocow [4]string

	listaowocow[0] = "jablko"
	listaowocow[1] = "gruszka"
	listaowocow[3] = "truskawka"

	fmt.Println("Lista owocow to", listaowocow)       //daje nam to co w liscie
	fmt.Println("Lista owocow to:", len(listaowocow)) //daje to ile jest elementow w liscie

	var listawarzyw = [3]string{"marchewka", "pietruszka", "seler"}

	fmt.Println("Lista warzyw to", listawarzyw)
	fmt.Println("Lista warzyw to:", len(listawarzyw))

}
