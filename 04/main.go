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

	question1 := 0
	/*for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			question1 += search(i, j, 0, -1, 0, "XMAS", field)  // up
			question1 += search(i, j, 0, -1, 1, "XMAS", field)  // up right
			question1 += search(i, j, 0, 0, 1, "XMAS", field)   // right
			question1 += search(i, j, 0, 1, 1, "XMAS", field)   // down right
			question1 += search(i, j, 0, 1, 0, "XMAS", field)   // down
			question1 += search(i, j, 0, 1, -1, "XMAS", field)  // down left
			question1 += search(i, j, 0, 0, -1, "XMAS", field)  // left
			question1 += search(i, j, 0, -1, -1, "XMAS", field) // up left
		}
	}*/

	//Two MAS in the shape of an X
	question2 := 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			m1 := search(i, j, 0, -1, -1, "AM", field) + search(i, j, 0, 1, 1, "AS", field) //MAS or
			m2 := search(i, j, 0, -1, -1, "AS", field) + search(i, j, 0, 1, 1, "AM", field) //SAM (left to right)

			m3 := search(i, j, 0, -1, 1, "AM", field) + search(i, j, 0, 1, -1, "AS", field) //MAS or
			m4 := search(i, j, 0, -1, 1, "AS", field) + search(i, j, 0, 1, -1, "AM", field) //SAM (right to left)

			if (m1 == 2 || m2 == 2) && (m3 == 2 || m4 == 2) {
				question2++
			}
		}
	}

	println("Question 1:", question1)
	println("Question 2:", question2)
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
