package main

import "fmt"
import "time"


func main() {
	fmt.Println("Witaj w naucze czasu w golangu")

	PresentTime := time.Now()
	fmt.Println("Aktualna data i czas:", PresentTime)

	fmt.Println(PresentTime.Format("2006-01-02 15:04:05"))

	createdate := time.Date(2020, time.October, 10, 23, 23, 0, 0, time.Local)
	fmt.Println(createdate)
	fmt.Println(createdate.Format("2006-01-02 15:04:05"))		

}