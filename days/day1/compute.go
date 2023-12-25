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
	letters = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	str_content := string(content)
	for _, val := range strings.Split(str_content, "\n") {
		//first_letter , second_letter : = -1, -1

	}
}

func init() {
	globals.FuncMap["1-1"] = computeFirst
	globals.FuncMap["1-2"] = computeSecond
}
