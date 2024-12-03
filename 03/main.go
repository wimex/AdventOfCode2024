package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var regex = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	var numbers = make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			numbers = append(numbers, []int{num1, num2})
		}
	}

	question1 := 0
	for _, nums := range numbers {
		question1 += nums[0] * nums[1]
	}

	fmt.Println("Question 1:", question1)
}
