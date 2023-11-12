package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getDraws(choice int) ([]string, error) {
	var url string
	if choice == 1 {
		url = "http://www.mbnet.com.pl/dl.txt"
	} else if choice == 2 {
		url = "http://www.mbnet.com.pl/dl_plus.txt"
	} else {
		return nil, fmt.Errorf("Invalid choice")
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var draws []string
	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		draws = append(draws, scanner.Text())
	}

	return draws, nil
}

func extractNumbers(draws []string) [][]int {
	var wszystkieLosowania [][]int

	for _, line := range draws {
		fields := strings.Fields(line)
		numbersStr := strings.Split(fields[2], ",")
		var numbers []int
		for _, numStr := range numbersStr {
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}
		wszystkieLosowania = append(wszystkieLosowania, numbers)
	}

	return wszystkieLosowania
}

func zbior10NajczesciejPadajacych5(wszystkieLosowania [][]int) map[string]int {
	wszystkieKombinacje5 := combinations(1, 50, 5)
	zbiorKombinacjiTrzechLiczb := make(map[string]int)

	for _, k := range wszystkieKombinacje5 {
		key := fmt.Sprintf("%v", k)
		for _, losowanie := range wszystkieLosowania {
			if containsAll(losowanie, k) {
				if _, ok := zbiorKombinacjiTrzechLiczb[key]; !ok {
					zbiorKombinacjiTrzechLiczb[key] = 1
				} else {
					zbiorKombinacjiTrzechLiczb[key]++
				}
			}
		}
	}

	return zbiorKombinacjiTrzechLiczb
}

func containsAll(numbers, subset []int) bool {
	for _, num := range subset {
		found := false
		for _, n := range numbers {
			if num == n {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func combinations(start, end, length int) [][]int {
	var result [][]int
	var recur func(index int, current []int)

	recur = func(index int, current []int) {
		if len(current) == length {
			temp := make([]int, length)
			copy(temp, current)
			result = append(result, temp)
			return
		}

		for i := index; i <= end; i++ {
			recur(i+1, append(current, i))
		}
	}

	recur(start, []int{})
	return result
}

func main() {
	var choice int
	fmt.Print("1-Lotto\n2-Plus: ")
	fmt.Scan(&choice)

	draws, err := getDraws(choice)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	wszystkieLosowania := extractNumbers(draws)

	zbiorKombinacjiTrzechLiczb := zbior10NajczesciejPadajacych5(wszystkieLosowania)

	var sDic []struct {
		Key   string
		Value int
	}
	for k, v := range zbiorKombinacjiTrzechLiczb {
		sDic = append(sDic, struct {
			Key   string
			Value int
		}{k, v})
	}

	sort.Slice(sDic, func(i, j int) bool {
		return sDic[i].Value > sDic[j].Value
	})

	var fileName string
	fmt.Print("Podaj nazwę pliku: ")
	fmt.Scan(&fileName)

	file, err := os.Create(fileName + ".csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	losowanieWriter := csv.NewWriter(file)
	defer losowanieWriter.Flush()

	for _, entry := range sDic {
		err := losowanieWriter.Write(strings.Fields(fmt.Sprintf("%v,%d", entry.Key, entry.Value)))
		if err != nil {
			fmt.Println("Error writing to CSV:", err)
			return
		}
	}
}
