package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

var print = fmt.Print

func main() {

	/*

		// const name, age = "Robert", 34
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

		// USER INPUT

		var userName string
		print("Enter your name:\n") // uzyty alias  ----> var print = fmt.Print
		fmt.Scan(&userName)
		fmt.Printf("Hello %v\n", userName)

		// STRINGI

		myStr := "Python"

		fmt.Println(myStr[0], string(myStr[0]))
		fmt.Println(myStr[1], string(myStr[1]))
		fmt.Println(myStr[2], string(myStr[2]))
		fmt.Println(myStr[3], string(myStr[3]))
		fmt.Println(myStr[4], string(myStr[4]))
		fmt.Println(myStr[5], string(myStr[5]))
		fmt.Printf("Długość: %v \n", len(myStr))

		// stringi są niemutowlane jak w pythonie

		editedMyStr := []byte(myStr)
		editedMyStr[0] = 'Y' // teraz można edytować (podać index)

		fmt.Println(editedMyStr) // []byte(x) -->runy, mozna edytowac, jesli spowrotem do stringa to string(x)
		fmt.Println(string(editedMyStr))

		editedMyStr[0] = 'M'
		fmt.Println(string(editedMyStr))
	*/

	anotherString := "Robert weterynarz"

	fmt.Printf("Lenght with len() = %v \nLenght with utf8.RuneCountInString() = %v \n", len(anotherString), utf8.RuneCountInString(anotherString))
	fmt.Printf("Type with len() = %T \nType with utf8.RuneCountInString() = %T \n", len(anotherString), utf8.RuneCountInString(anotherString))

	anotherStringSplit := strings.Split(anotherString, " ")

	fmt.Printf("%v\n%v \n", anotherStringSplit[0], anotherStringSplit[1])

	// ZMIANA TYPÓW
	// int --> string

	myInt := 100
	myStr := strconv.Itoa(myInt)
	onceAgainInt, _ := strconv.Atoi(myStr) // deklaracja

	fmt.Printf("myInt: %T \nmyStr: %T\nonceAgainInt: %T \n", myInt, myStr, onceAgainInt)

}
