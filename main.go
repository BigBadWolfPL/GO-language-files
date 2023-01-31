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

	// SPLIT, JOIN

	anotherString := "Robert weterynarz"

	fmt.Printf("Lenght with len() = %v \nLenght with utf8.RuneCountInString() = %v \n", len(anotherString), utf8.RuneCountInString(anotherString))
	fmt.Printf("Type with len() = %T \nType with utf8.RuneCountInString() = %T \n", len(anotherString), utf8.RuneCountInString(anotherString))

	anotherStringSplit := strings.Split(anotherString, " ")

	fmt.Println("anotherStringSplit: ", anotherStringSplit)
	fmt.Printf("%v\n%v \n", anotherStringSplit[0], anotherStringSplit[1])

	// JOIN

	// OPTION 1:

	anotherStringJoin := strings.Join([]string{anotherStringSplit[0], anotherStringSplit[1]}, " ")
	fmt.Printf("Join: %v \n", anotherStringJoin)

	// OPTION 2:

	someString := "Ala ma kota \n"
	someString += "Kot ma Ale"

	fmt.Println("Some string:\n", someString)

	otherString := anotherStringSplit[0]
	otherString += " programista"

	println("Other String:", otherString)

	otherString += "programista"

	fmt.Println("Other String added:")

	//OPTION 3:
	programista := "programista"
	weterynarz := anotherStringSplit[1]

	newSTRING := &strings.Builder{}
	newSTRING.WriteString(anotherStringSplit[0] + " ")
	newSTRING.WriteString(programista)
	newSTRING.WriteString(" i ")
	newSTRING.WriteString(weterynarz)

	newSTRINGstring := newSTRING.String()

	fmt.Printf("\nMetoda: &strings.Builder{}\nWartość: %v\nTyp: %T\nKonwersja Typu: %T\njeszcze raz konwercja: %T \n", newSTRING, newSTRING, newSTRING.String(), newSTRINGstring)

	// OPCJA 2 ZALECANA!

	// ZMIANA TYPÓW
	// int --> string

	myInt := 100
	myStr := strconv.Itoa(myInt)           // INT to ASCII
	onceAgainInt, _ := strconv.Atoi(myStr) // deklaracja ASCII to INT

	fmt.Printf("myInt: %T \nmyStr: %T\nonceAgainInt: %T \n", myInt, myStr, onceAgainInt)

	// CONTAINS + WARUNEK

	STRING := newSTRING.String() // trzeba zmienić typ bo domyslinie &strings.Builder zwraca typ: *strings.Builder

	if strings.Contains(STRING, "programista") {
		fmt.Println("JESTEŚ PROGRAMISTĄ")
	} else {
		fmt.Println("JESTEŚ WETERYNARZEM")
	}

	i := 1

	switch i {
	case 0:
		fmt.Println("Wartość 0")
	case 1:
		fmt.Println("Wartość 1")
	case 2:
		fmt.Println("Wartość 2")
	case 3:
		fmt.Println("Wartość 3")
	case 15:
		fmt.Println("Wartość 15")
	default:
		fmt.Println("Inna wartość")
	}

	// else if

	var choiceNum int

	fmt.Println("Podaj numer:")
	fmt.Scan(&choiceNum)

	if choiceNum < 5 {
		fmt.Println("Liczba mniejsza od 5")
	} else if choiceNum == 5 {
		fmt.Println("Liczba = 5")
	} else {
		fmt.Println("Liczba większa od 5")
	}

}
