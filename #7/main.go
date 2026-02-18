package main

import "fmt"

func main() {
	fmt.Println("Witam w swiecie wksaznikow")

	//var ptr *int

	//fmt.Println("watrtosc wskaznika to", ptr)

	myNumber := 23

	var ptr *int = &myNumber

	fmt.Println("wartosc wskaznika to", ptr)
	fmt.Println("wartosc wskaznika to", *ptr)

	*ptr = *ptr + 2

	fmt.Println("nowa wartosc wskaznika to", myNumber)

}
