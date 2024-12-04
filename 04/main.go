package main

import (
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	field := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		array := make([]string, 0)
		for _, char := range line {
			array = append(array, string(char))
		}

		field = append(field, array)
	}

	needle := "XMAS"
	question1 := 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			question1 += search(i, j, 0, -1, 0, needle, field)  // up
			question1 += search(i, j, 0, -1, 1, needle, field)  // up right
			question1 += search(i, j, 0, 0, 1, needle, field)   // right
			question1 += search(i, j, 0, 1, 1, needle, field)   // down right
			question1 += search(i, j, 0, 1, 0, needle, field)   // down
			question1 += search(i, j, 0, 1, -1, needle, field)  // down left
			question1 += search(i, j, 0, 0, -1, needle, field)  // left
			question1 += search(i, j, 0, -1, -1, needle, field) // up left
		}
	}

	println("Question 1:", question1)
}

func search(i int, j int, p int, di int, dj int, needle string, haystack [][]string) int {
	if p == len(needle) {
		return 1
	}

	if i < 0 || i >= len(haystack) || j < 0 || j >= len(haystack[i]) {
		return 0
	}

	if haystack[i][j] != string(needle[p]) {
		return 0
	}

	return search(i+di, j+dj, p+1, di, dj, needle, haystack)
}
