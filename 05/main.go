package main

import (
	"bufio"
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
	invalids := make([][]int, 0)
	for _, page := range pages {
		valid := true
		for index, num := range page {
			if before, ok := rules[num]; ok {
				if evaluate(index, before, page) >= 0 {
					valid = false
					invalids = append(invalids, page)

					break
				}
			}
		}

		if valid {
			middle := (len(page) + 1) / 2
			question1 += page[middle-1]
		}
	}

	question2 := 0
	for _, invalid := range invalids {
		ordered := reorder(invalid, rules)
		middle := (len(ordered) + 1) / 2
		question2 += ordered[middle-1]
	}

	println("Question 1: ", question1)
	println("Question 2: ", question2)
}

func evaluate(index int, needles []int, haystack []int) int {
	for i := index; i >= 0; i-- {
		num := haystack[i]
		for _, needle := range needles {
			if needle == num {
				return i
			}
		}
	}

	return -1
}

func reorder(invalid []int, rules map[int][]int) []int {
	for index := 0; index < len(invalid); index++ { // Go through the numbers in the current list
	restart:
		num := invalid[index]

		// Check if the current number has a rule
		if rule, ok := rules[num]; ok {

			// It has a rule, go through it. The current number should come BEFORE the numbers in the rule
			for _, after := range rule {

				// In the array, we are at <index>. Let's go back and see that any of the numbers in the rule come before
				for i := index; i >= 0; i-- {

					// If it does, swap the numbers
					if invalid[i] == after {
						temp := invalid[i]
						invalid[i] = invalid[index]
						invalid[index] = temp
						index = 0

						goto restart // The horror... the array has been modified, let's again if the order is correct
					}
				}
			}
		}
	}

	return invalid
}
