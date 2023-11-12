package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	choice := 1

	var url string

	if choice == 1 {
		url = "http://www.mbnet.com.pl/dl.txt"
	} else if choice == 2 {
		url = "http://www.mbnet.com.pl/dl_plus.txt"
	} else {
		fmt.Println("Nie rozpoznano wyboru")
		return
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = ioutil.WriteFile("res.txt", body, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	data, err := ioutil.ReadFile("res.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	content := string(data)

	re := regexp.MustCompile(`\d+,\d+,\d+,\d+,\d+,\d+`)
	matches := re.FindAllString(content, -1)

	var numbers [][]int
	for _, match := range matches {
		parts := strings.Split(match, ",")
		var nums []int
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			nums = append(nums, num)
		}
		numbers = append(numbers, nums)
	}

	numberOfDraws := len(numbers)

	if numberOfDraws == 0 {
		fmt.Println("No data found")
		return
	}

	counts := make(map[int][]int)

	for i, nums := range numbers {
		counts[i] = nums
	}

	topCombinations := zbior10NajczesciejPadajacych4(counts)

	// Sort the combinations by frequency
	sort.Slice(topCombinations, func(i, j int) bool {
		return topCombinations[i].Count > topCombinations[j].Count
	})

	file, err := os.Create("most_common_4.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, combination := range topCombinations {
		row := []string{}
		for _, num := range combination.Combination {
			row = append(row, strconv.Itoa(num))
		}
		writer.Write(row)
	}
}

type combinationWithCount struct {
	Combination []int
	Count       int
}

func zbior10NajczesciejPadajacych4(draws map[int][]int) []combinationWithCount {
	combos := make(map[[4]int]int)

	for _, nums := range draws {
		for _, combo := range combinations(nums, 4) {
			var key [4]int
			copy(key[:], combo)
			combos[key]++
		}
	}

	var combinationsWithCount []combinationWithCount

	for combo, count := range combos {
		var combination []int
		combination = append(combination, combo[0], combo[1], combo[2], combo[3])
		combinationsWithCount = append(combinationsWithCount, combinationWithCount{
			Combination: combination,
			Count:       count,
		})
	}

	return combinationsWithCount
}

func combinations(arr []int, k int) [][]int {
	if k == 0 {
		return [][]int{{}}
	}
	if len(arr) == 0 {
		return nil
	}

	head, tail := arr[0], arr[1:]
	excludeHead := combinations(tail, k)
	includeHead := combinations(tail, k-1)
	for i := range includeHead {
		includeHead[i] = append([]int{head}, includeHead[i]...)
	}
	return append(excludeHead, includeHead...)
}
