package day3

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/naruto678/aoc-go/globals"
)

var directions = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{1, -1},
	{1, 0},
	{1, 1},
	{0, 1},
	{0, -1},
}

func parseInput(lines []string) ([][]byte, [][]bool) {
	content := make([][]byte, len(lines))
	visited := make([][]bool, len(lines))
	for idx := range content {
		content[idx] = make([]byte, len(lines[idx]))
		visited[idx] = make([]bool, len(lines[idx]))
	}

	for i := range content {
		for j := range content[i] {
			content[i][j] = lines[i][j]
		}
	}

	return content, visited
}

func computeFirst(content []byte) {
	arr, visited := parseInput(strings.Split(string(content), "\n"))

	var dfs func(int, int)

	isValid := func(row, col int) bool {
		if row < 0 || col < 0 || row >= len(arr) || col >= len(arr[0]) {
			return false
		}
		return true
	}
	dfs = func(row, col int) {

		if !isValid(row, col) {
			return
		}
		visited[row][col] = true
		for _, dir := range directions {
			x := row + dir[0]
			y := col + dir[1]
			if isValid(x, y) && !visited[x][y] && arr[x][y] != '.' {
				dfs(x, y)
			}
		}

	}

	isSymbol := func(letter byte) bool {
		if letter == '.' || unicode.IsDigit(rune(letter)) {
			return false
		}
		return true
	}

	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] != '.' && isSymbol(arr[i][j]) && !visited[i][j] {
				dfs(i, j)
			}
		}
	}

	for i := range arr {
		for j := range arr[i] {
			if isSymbol(arr[i][j]) {
				visited[i][j] = false
			}
		}
	}

	total_sum := 0
	for i := range arr {
		start := 0

		current_sum := 0

		for start < len(arr[i]) {
			if visited[i][start] {
				current_sum = current_sum*10 + int(arr[i][start]-'0')
			} else {
				total_sum += current_sum
				current_sum = 0
			}
			start++
		}
		if current_sum != 0 {
			total_sum += current_sum
		}
	}
	fmt.Println(fmt.Sprintf("computed for the first part %d", total_sum))

}

func computeSecond(content []byte) {
	arr, visited := parseInput(strings.Split(string(content), "\n"))

	var dfs func(int, int)

	isValid := func(row, col int) bool {
		if row < 0 || col < 0 || row >= len(arr) || col >= len(arr[0]) {
			return false
		}
		return true
	}
	dfs = func(row, col int) {

		if !isValid(row, col) {
			return
		}
		visited[row][col] = true
		for _, dir := range directions {
			x := row + dir[0]
			y := col + dir[1]
			if isValid(x, y) && !visited[x][y] && arr[x][y] != '.' {
				dfs(x, y)
			}
		}

	}

	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] == '*' && !visited[i][j] {
				dfs(i, j)
			}
		}
	}

	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] == '*' {
				visited[i][j] = false
			}
		}
	}
	prettyPrint(visited)

	total := 0
	for i := range arr {
		numbers := []int{}
		currentNum := 0
		for start := 0; start < len(arr[i]); start++ {
			if visited[i][start] {
				currentNum = currentNum*10 + int(arr[i][start]-'0')
			} else {
				if currentNum > 0 {
					numbers = append(numbers, currentNum)
					currentNum = 0
				}
			}
		}
		if currentNum > 0 {
			numbers = append(numbers, currentNum)
		}

		if len(numbers) < 2 {
			continue
		}
		fmt.Println(numbers)
		total += numbers[0] * numbers[1]

	}
	fmt.Println(fmt.Sprintf("computed the second part %d", total))

}

func prettyPrint(visited [][]bool) {
	for idx := range visited {
		result := []int{}
		for _, val := range visited[idx] {
			if val {
				result = append(result, 1)
			} else {
				result = append(result, 0)
			}
		}
		fmt.Println(result)
	}
}

func init() {
	globals.FuncMap["3-1"] = computeFirst
	globals.FuncMap["3-2"] = computeSecond
}
