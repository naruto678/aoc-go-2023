package day1

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/naruto678/aoc-go/globals"
)

var computeFirst = func(content []byte) {
	str_content := string(content)
	sum := 0
	for _, val := range strings.Split(str_content, "\n") {
		first_letter, second_letter := -1, -1
		for idx := range val {
			isDigit := unicode.IsDigit(rune(val[idx]))
			if isDigit && first_letter == -1 {
				first_letter = int(val[idx] - '0')
				second_letter = first_letter
			} else if isDigit {
				second_letter = int(val[idx] - '0')
			}
		}

		sum += first_letter*10 + second_letter
	}
	fmt.Println(fmt.Sprintf("first part %d", sum))
}

var computeSecond = func(content []byte) {
	sum := 0
	str_content := string(content)

	letterMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	getDecimalValue := func(letter string) int {
		result, found := letterMap[letter]
		if !found {
			panic("calling letter map with wrong values" + letter)
		}
		return result
	}

	for _, val := range strings.Split(str_content, "\n") {

		for letter := range letterMap {
			val = strings.ReplaceAll(val, letter, fmt.Sprintf("%s%d%s", letter, getDecimalValue(letter), letter))
		}
		//fmt.Println(val)
		first, second := -1, -1
		for idx := range val {
			isDigit := unicode.IsDigit(rune(val[idx]))
			if isDigit {
				if first == -1 {
					first = int(val[idx] - '0')
					second = first
				} else {
					second = int(val[idx] - '0')
				}
			}
		}
		sum += 10*first + second

	}
	fmt.Println(fmt.Sprintf("computing second %d", sum))
}

func init() {
	globals.FuncMap["1-1"] = computeFirst
	globals.FuncMap["1-2"] = computeSecond
}
