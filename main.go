package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/naruto678/aoc-go/globals"
	_ "github.com/naruto678/aoc-go/register"
)

type AOC struct {
	log *log.Logger
}

func main() {
	day := flag.Int("day", -1, "day to run this code for ")
	test := flag.Bool("test", false, "should this run for test  input for real input")
	flag.Parse()

	if *day == -1 {
		panic(fmt.Sprintf("day %d does not exist", *day))
	}
	var file_name string
	if *test {
		file_name = "test.txt"
	} else {
		file_name = "input.txt"
	}

	wd, _ := os.Getwd()
	full_path := strings.Join([]string{wd, "days", fmt.Sprintf("day%d", *day), file_name}, string(os.PathSeparator))
	if _, err := os.Stat(full_path); err == nil {
		content, err := os.ReadFile(full_path)
		if err != nil {
			panic(err)
		}
		computeFirst(content, *day)
		computeSecond(content, *day)
	} else {
		panic(err)
	}
}

func computeFirst(content []byte, day int) {
	fmt.Println(fmt.Sprintf("running the first part for day %d", day))
	stub := globals.FuncMap[fmt.Sprintf("%d-%d", day, 1)]
	stub(content)
}

func computeSecond(content []byte, day int) {
	fmt.Println(fmt.Sprintf("running the second part for day %d", day))
	stub := globals.FuncMap[fmt.Sprintf("%d-%d", day, 2)]
	stub(content)
}
