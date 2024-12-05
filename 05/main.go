package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		pieces := strings.Split(line, "|")

		num1, _ := strconv.Atoi(pieces[0])
		num2, _ := strconv.Atoi(pieces[1])

		if _, ok := rules[num1]; ok {
			rules[num1] = append(rules[num1], num2)
		} else {
			rules[num1] = []int{num2}
		}
	}

	pages := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		pieces := strings.Split(line, ",")

		var numbers []int
		for _, piece := range pieces {
			num, _ := strconv.Atoi(piece)
			numbers = append(numbers, num)
		}

		pages = append(pages, numbers)
	}

	question1 := 0
	for _, page := range pages {
		valid := true
		for index, num := range page {
			if before, ok := rules[num]; ok {
				if evaluate(index, before, page) == false {
					valid = false
					break
				}
			}
		}

		if valid {
			middle := (len(page) + 1) / 2
			question1 += page[middle-1]
			fmt.Println("Middle: ", page[middle])
		}

		fmt.Println(page)
		fmt.Println("Valid: ", valid)
	}

	println("Question 1: ", question1)
}

func evaluate(index int, needles []int, haystack []int) bool {
	for i := index; i >= 0; i-- {
		num := haystack[i]
		for _, needle := range needles {
			if needle == num {
				return false
			}
		}
	}

	return true
}
