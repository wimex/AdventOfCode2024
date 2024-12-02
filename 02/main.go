package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var lines [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		pieces := strings.Split(line, " ")

		var numbers []int
		for _, piece := range pieces {
			num, _ := strconv.Atoi(piece)
			numbers = append(numbers, num)
		}

		lines = append(lines, numbers)
	}

	var question1 []bool
	var question2 []bool
	for _, line := range lines {
		valid := validate(line)
		question1 = append(question1, valid)

		if !valid {
			for i := 0; i < len(line); i++ {
				temp := make([]int, len(line))
				copy(temp, line)

				temp = append(temp[:i], temp[i+1:]...)
				if validate(temp) {
					valid = true
					break
				}
			}
		}

		question2 = append(question2, valid)
	}

	println("Question 1:", filter(question1))
	println("Question 2:", filter(question2))
}

func filter(lines []bool) int {
	result := 0
	for _, line := range lines {
		if line {
			result++
		}
	}

	return result
}

func validate(numbers []int) bool {
	step := 0

	for i := 0; i < len(numbers)-1; i++ {
		num1 := numbers[i]
		num2 := numbers[i+1]
		diff := math.Abs(float64(num1) - float64(num2))
		dire := direction(num1, num2)

		if diff < 1 || diff > 3 {
			return false
		}

		if step == 0 {
			step = dire
			continue
		}

		if step != dire {
			return false
		}
	}

	return true
}

func direction(num1, num2 int) int {
	if num1 == num2 {
		return 0
	}

	if num1 > num2 {
		return -1
	}

	return 1
}
