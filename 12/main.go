package main

import (
	"bufio"
	"fmt"
	"os"
)

var CACHE = map[string]int{}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	plot := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for _, char := range line {
			row = append(row, string(char))
		}

		plot = append(plot, row)
	}

	width := len(plot[0])
	height := len(plot)
	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}

	question1 := 0
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			if visited[j][i] {
				continue
			}

			letter := plot[j][i]
			area, perimeter := calculate(letter, plot, visited, i, j, width, height)
			question1 += area * perimeter
		}
	}

	fmt.Println(question1)
}

func calculate(letter string, plot [][]string, visited [][]bool, x, y, width, height int) (int, int) {
	if x < 0 || x >= width || y < 0 || y >= height {
		return 0, 0
	}

	if visited[y][x] {
		return 0, 0
	}

	if plot[y][x] != letter {
		return 0, 0
	}

	lx := x - 1
	ly := y
	rx := x + 1
	ry := y
	tx := x
	ty := y - 1
	bx := x
	by := y + 1

	visited[y][x] = true
	curr_area := 1
	curr_perimeter := 0

	if lx < 0 || lx >= width || ly < 0 || ly >= height || plot[ly][lx] != letter {
		curr_perimeter++
	}
	if rx < 0 || rx >= width || ry < 0 || ry >= height || plot[ry][rx] != letter {
		curr_perimeter++
	}
	if tx < 0 || tx >= width || ty < 0 || ty >= height || plot[ty][tx] != letter {
		curr_perimeter++
	}
	if bx < 0 || bx >= width || by < 0 || by >= height || plot[by][bx] != letter {
		curr_perimeter++
	}

	a1, p1 := calculate(letter, plot, visited, x+1, y, width, height)
	a2, p2 := calculate(letter, plot, visited, x-1, y, width, height)
	a3, p3 := calculate(letter, plot, visited, x, y+1, width, height)
	a4, p4 := calculate(letter, plot, visited, x, y-1, width, height)

	return curr_area + a1 + a2 + a3 + a4, curr_perimeter + p1 + p2 + p3 + p4
}
