package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"time"
)

func luckyNums() []int {
	countdown := []string{"1st", "2nd", "3rd", "4th", "5th", "6th"}
	nums := make([]int, 0)
	for _, l := range countdown {
		var numStr string
		fmt.Printf("Enter %s number: ", l)
		fmt.Scanln(&numStr)
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Invalid input: %s is not a number\n", numStr)
			return nil
		}
		nums = append(nums, num)
	}
	sort.Ints(nums)
	return nums
}

func simulator(wybierz string) map[string]float64 {
	rand.Seed(time.Now().UnixNano())
	licznikLosowan := 0
	var numeryUzytkownika []int
	var ileLosow int

	if wybierz == "w" {
		numeryUzytkownika = luckyNums()
		if numeryUzytkownika == nil {
			return nil
		}
		ileLosow = 1
	} else if wybierz == "l" {
		var err error
		fmt.Printf("Podaj ilosc zakładów na jednym kuponie: ")
		var input string
		fmt.Scanln(&input)
		ileLosow, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input: number of bets must be an integer")
			return nil
		}
		numeryUzytkownika = make([]int, ileLosow*6)
		for i := 0; i < ileLosow*6; i += 6 {
			nums := make([]int, 6)
			for j := 0; j < 6; j++ {
				nums[j] = rand.Intn(49) + 1
			}
			sort.Ints(nums)
			copy(numeryUzytkownika[i:i+6], nums)
		}
	} else {
		fmt.Println("Invalid input: select 'w' for own numbers or 'l' for random numbers")
		return nil
	}

	for {
		randomowe := make([]int, 6)
		for i := 0; i < 6; i++ {
			randomowe[i] = rand.Intn(49) + 1
		}
		sort.Ints(randomowe)
		licznikLosowan++

		found := false
		for j := 0; j < len(numeryUzytkownika); j += 6 {
			if reflect.DeepEqual(randomowe, numeryUzytkownika[j:j+6]) {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	dic := map[string]float64{
		"Tygodnie": float64(licznikLosowan / 3),
		"Lata":     float64((licznikLosowan / 3) / 56),
		"koszt":    float64((licznikLosowan * 3) * ileLosow),
	}

	return dic
}

func main() {
	var wybierz string
	fmt.Println("Podaj (w)-własne liczby lub (l) - losowe: ")
	fmt.Scanln(&wybierz)
	dic := simulator(wybierz)

	if dic != nil {
		for c, v := range dic {
			fmt.Printf("%s, %.2f\n", c, v)
		}
	}
}
