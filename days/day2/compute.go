package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/naruto678/aoc-go/globals"
)

const (
	RED_COUNT   = 12
	GREEN_COUNT = 13
	BLUE_COUNT  = 14
	BLUE        = "blue"
	RED         = "red"
	GREEN       = "green"
)

func computeFirst(content []byte) {
	str_content := string(content)
	count := 0

	for _, line := range strings.Split(str_content, "\n") {
		isPossible := true
		games := strings.Split(line, ": ")
		game, line := games[0], games[1]

		idx, _ := strconv.Atoi(strings.Split(game, " ")[1])
		//fmt.Println(idx)

	outer:
		for _, set := range strings.Split(line, ";") {
			balls := strings.Split(set, ",")
			for _, ball := range balls {
				ball = strings.TrimSpace(ball)

				attrs := strings.Split(ball, " ")
				color := attrs[1]
				num, err := strconv.Atoi(attrs[0])
				if err != nil {
					panic(err)
				}
				//fmt.Println(num, color)
				switch color {
				case BLUE:
					if num > BLUE_COUNT {
						isPossible = false
						break outer

					}
				case GREEN:
					if num > GREEN_COUNT {
						isPossible = false
						break outer
					}

				case RED:
					if num > RED_COUNT {
						isPossible = false
						break outer
					}
				}
			}
		}
		if isPossible {
			count += idx
		}
	}
	fmt.Println(fmt.Sprintf("computed the first part %d", count))
}

func computeSecond(content []byte) {
	str_content := string(content)
	count := 0

	for _, line := range strings.Split(str_content, "\n") {
		line := strings.Split(line, ": ")[1]

		//fmt.Println(idx)
		red_count, blue_count, green_count := 0, 0, 0

		for _, set := range strings.Split(line, ";") {
			balls := strings.Split(set, ",")
			for _, ball := range balls {
				ball = strings.TrimSpace(ball)

				attrs := strings.Split(ball, " ")
				color := attrs[1]
				num, err := strconv.Atoi(attrs[0])
				if err != nil {
					panic(err)
				}
				//fmt.Println(num, color)
				switch color {
				case BLUE:
					blue_count = max(blue_count, num)
				case GREEN:
					green_count = max(green_count, num)

				case RED:
					red_count = max(red_count, num)
				}
			}

		}
		//fmt.Println(blue_count, red_count, green_count)
		count += blue_count * red_count * green_count
	}
	fmt.Println(fmt.Sprintf("computed the second part %d", count))

}

func init() {
	globals.FuncMap["2-1"] = computeFirst
	globals.FuncMap["2-2"] = computeSecond
}
