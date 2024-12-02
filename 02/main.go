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
	for _, line := range lines {
		direction := 0
		valid := true

		for i := 0; i < len(line)-1; i++ {
			num1 := line[i]
			num2 := line[i+1]
			diff := math.Abs(float64(num1) - float64(num2))
			dire := func() int {
				if num1 == num2 {
					return 0
				}
				if num1 > num2 {
					return -1
				}
				return 1
			}()

			if diff < 1 || diff > 3 {
				valid = false
				break
			}

			if direction == 0 {
				direction = dire
				continue
			}

			if direction != dire {
				valid = false
				break
			}
		}

		if valid {
			question1 = append(question1, valid)
		}
	}

	println("Question 1:", len(question1))
}
