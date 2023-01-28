package main

import (
	"fmt"
)

var wiek int

func main() {
	const name, age = "Robert", 34
	wiek = 35
	fmt.Println(age, name, wiek)
	fmt.Printf("Programista %v %v ma %v lata.", name, age, wiek)

	var word string
	var number int

	fmt.Print("\nPodaj słowo: ")
	fmt.Scan(&word)
	fmt.Print("Podaj liczbę: ")
	fmt.Scanln(&number)
	fmt.Printf("Twoje słowo to %v a liczba = %v. Słowo jeszcze raz %v", word, number, word)
}
