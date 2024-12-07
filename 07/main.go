package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache = make(map[string]int64)
var operands = map[string]func(int64, int64) int64{
	"+": func(a, b int64) int64 { return a + b },
	"*": func(a, b int64) int64 { return a * b },
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	values := make(map[int64][]int64)
	for scanner.Scan() {
		line := scanner.Text()
		pieces := strings.Split(line, ": ")

		left, err := strconv.ParseInt(pieces[0], 10, 64)
		if err != nil {
			fmt.Println("Failed to parse int: ", pieces[0])
			fmt.Println(err)
			return
		}

		right := strings.Split(pieces[1], " ")
		numbers := make([]int64, len(right))
		for i, v := range right {
			num, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				fmt.Println("Failed to parse int: ", v)
				fmt.Println(err)
				return
			}
			numbers[i] = num
		}

		if _, ok := values[left]; ok {
			fmt.Println("Duplicate key: ", left)
			return
		}

		values[left] = numbers
	}

	question1 := int64(0)
	for test, numbers := range values {
		result, ops := evaluate(test, -1, numbers, []string{})
		if result {
			fmt.Println("Found:", test)
			fmt.Println("Ops:", ops)

			question1 += test
		}
	}

	fmt.Println("Question 1: ", question1)
}

func evaluate(expected int64, left int64, numbers []int64, combiner []string) (bool, []string) {
	if len(numbers) == 0 {
		return false, combiner
	}
	if left > expected {
		return false, combiner
	}

	for operand, function := range operands {
		current := numbers[0]
		pivot := getPivot(operand, left)
		result := function(pivot, current)
		if result == expected {
			return true, append(combiner, operand)
		}

		recursion, rr := evaluate(expected, result, numbers[1:], append(combiner, operand))
		if recursion {
			return true, rr
		}
	}

	return false, combiner
}

func getPivot(operand string, left int64) int64 {
	if left != -1 {
		return left
	}

	if operand == "+" {
		return 0
	}

	if operand == "*" {
		return 1
	}

	return -1
}
