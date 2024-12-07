package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var operands1 = map[string]func(int64, int64) int64{
	"+": func(a, b int64) int64 { return a + b },
	"*": func(a, b int64) int64 { return a * b },
}

var operands2 = map[string]func(int64, int64) int64{
	"+": func(a, b int64) int64 { return a + b },
	"*": func(a, b int64) int64 { return a * b },
	"|": func(a, b int64) int64 {
		val1 := strconv.FormatInt(a, 10)
		val2 := strconv.FormatInt(b, 10)

		res, err := strconv.ParseInt(val1+val2, 10, 64)
		if err != nil {
			fmt.Println("Failed to convert to int64: ", val1, val2)
			return -1
		}

		return res
	},
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

	question1 := execute(values, operands1)
	question2 := execute(values, operands2)

	fmt.Println("Question 1: ", question1)
	fmt.Println("Question 2: ", question2)
}

func execute(values map[int64][]int64, operands map[string]func(int64, int64) int64) int64 {
	result := int64(0)
	for test, numbers := range values {
		eval, ops := evaluate(test, numbers[0], numbers[1:], operands, []string{})
		if eval {
			fmt.Println("NUM:", test)
			fmt.Println("TST:", numbers)
			fmt.Println("OPS:", ops)
			fmt.Println("")
			result += test
		}
	}

	return result
}

func evaluate(expected int64, left int64, numbers []int64, operands map[string]func(int64, int64) int64, combiner []string) (bool, []string) {
	if len(numbers) == 0 {
		return false, combiner
	}
	if left > expected {
		return false, combiner
	}

	for operand, function := range operands {
		current := numbers[0]
		result := function(left, current)
		if result == expected && len(numbers) == 1 {
			return true, append(combiner, operand)
		}

		recursion, rr := evaluate(expected, result, numbers[1:], operands, append(combiner, operand))
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
