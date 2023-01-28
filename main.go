package main

import (
	"fmt"
)

var wiek int

func main() {
	// const name, age = "Robert", 34
	// wiek = 35
	// fmt.Println(age, name, wiek)
	// fmt.Printf("Programista %v %v ma %v lata.", name, age, wiek)
	//
	// var word string
	// var number int
	//
	// fmt.Print("\nPodaj słowo: ")
	// fmt.Scan(&word)
	// fmt.Print("Podaj liczbę: ")
	// fmt.Scanln(&number)
	// fmt.Printf("Twoje słowo to %v a liczba = %v. Słowo jeszcze raz %v", word, number, word)

	var A int64 = 1
	var B int32 // bez przypisania wartosci przyjmuje 0
	C := 3      // kompilator zgaduje typ zmiennej, mozna tylko wew. funkcji

	var AA string = "Misia i miś"
	var BB string
	CC := "Miś mi Misia"

	fmt.Println("Zmienna A", A, "Zmienna B", B, "Zmienna C", C)

	fmt.Printf("Zmienna A %T \n", A)
	fmt.Printf("Zmienna B %T \n", B)
	fmt.Printf("Zmienna C %T \n", C)

	fmt.Println("Zmienna AA:", AA)
	fmt.Println("Zmienna BB:", BB)
	fmt.Println("Zmienna CC:", CC)

	fmt.Printf("Zmienna %T \n", AA)
	fmt.Printf("Zmienna %T \n", BB)
	fmt.Printf("Zmienna %T \n", CC)

	A = int64(B)
	F := 3.14

	fmt.Printf("Typ zmiennej F %T, jej wartość %v \nZmienna B = %v przypisana do A po zmianie typu \n \n", F, F, A)

	// STAŁE

	const stala1 = "stała string" // kompilator zgaduje typ
	const stala2 = 1234           // kompilator zgaduje typ
	const stala3 = -3.14          // kompilator zgaduje typ
	const stala4 int64 = 35

	fmt.Printf("Kolejno typy stałych: \n%T \n%T \n%T \n%T \n", stala1, stala2, stala3, stala4)

	// LICZNIK IOTA DLA STALYCH

	const (
		one = iota
		two
		three
		four
		five
		six
	)
	println(one, two, three, four, five, six)
}
