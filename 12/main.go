package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
)

type Coord struct {
	x int
	y int
}

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
	question2 := 0
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			if visited[j][i] {
				continue
			}

			letter := plot[j][i]
			area, perimeter, fence := calculate(letter, plot, visited, i, j, width, height)
			contour := draw(fence, width, height)
			length := corners(contour)

			question1 += area * perimeter
			question2 += area * length
		}
	}

	fmt.Println(question1)
	fmt.Println(question2)
}

func calculate(letter string, plot [][]string, visited [][]bool, x, y, width, height int) (int, int, []Coord) {
	if x < 0 || x >= width || y < 0 || y >= height {
		return 0, 0, []Coord{}
	}

	if visited[y][x] {
		return 0, 0, []Coord{}
	}

	if plot[y][x] != letter {
		return 0, 0, []Coord{}
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

	fence := []Coord{}
	if tx < 0 || tx >= width || ty < 0 || ty >= height || plot[ty][tx] != letter {
		curr_perimeter++
	}
	if rx < 0 || rx >= width || ry < 0 || ry >= height || plot[ry][rx] != letter {
		curr_perimeter++
	}
	if bx < 0 || bx >= width || by < 0 || by >= height || plot[by][bx] != letter {
		curr_perimeter++
	}
	if lx < 0 || lx >= width || ly < 0 || ly >= height || plot[ly][lx] != letter {
		curr_perimeter++
	}

	a1, p1, f1 := calculate(letter, plot, visited, x+1, y, width, height)
	a2, p2, f2 := calculate(letter, plot, visited, x-1, y, width, height)
	a3, p3, f3 := calculate(letter, plot, visited, x, y+1, width, height)
	a4, p4, f4 := calculate(letter, plot, visited, x, y-1, width, height)

	fence = append(fence, Coord{x, y})
	fence = append(fence, f1...)
	fence = append(fence, f2...)
	fence = append(fence, f3...)
	fence = append(fence, f4...)

	return curr_area + a1 + a2 + a3 + a4, curr_perimeter + p1 + p2 + p3 + p4, fence
}

func draw(fence []Coord, width, height int) [][]string {
	result := [][]string{}

	for j := 0; j < height+10; j++ {
		line := []string{}

		for i := 0; i < width+10; i++ {
			contains := lo.ContainsBy(fence, func(coord Coord) bool { return i == coord.x+5 && j == coord.y+5 })
			if contains {
				line = append(line, "*")
			} else {
				line = append(line, " ")
			}
		}

		result = append(result, line)
	}

	for j := 0; j < len(result); j++ {
		for i := 0; i < len(result[j]); i++ {
			fmt.Print(result[j][i])
		}

		fmt.Println()
	}

	return result
}

// Let's assume this piece of code has never happened
func corners(contour [][]string) int {
	count := 0
	for j := 0; j < len(contour); j++ {
		for i := 0; i < len(contour[j]); i++ {
			if contour[j][i] == "*" {
				if contour[j-1][i] == " " && contour[j][i-1] == " " {
					count++
				}
				if contour[j-1][i] == " " && contour[j][i+1] == " " {
					count++
				}
				if contour[j+1][i] == " " && contour[j][i-1] == " " {
					count++
				}
				if contour[j+1][i] == " " && contour[j][i+1] == " " {
					count++
				}

				if contour[j+1][i] == "*" && contour[j][i+1] == "*" && contour[j+1][i+1] == " " {
					count++
				}
				if contour[j-1][i] == "*" && contour[j][i+1] == "*" && contour[j-1][i+1] == " " {
					count++
				}
				if contour[j+1][i] == "*" && contour[j][i-1] == "*" && contour[j+1][i-1] == " " {
					count++
				}
				if contour[j-1][i] == "*" && contour[j][i-1] == "*" && contour[j-1][i-1] == " " {
					count++
				}
			}
		}
	}

	fmt.Println("COUNT", count)
	return count
}
