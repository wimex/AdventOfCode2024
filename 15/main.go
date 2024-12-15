package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Coordinate struct {
	x int
	y int
}

type Movement struct {
	position  Coordinate
	direction Coordinate
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	warehouse := [][]string{}
	movements := []string{}

	part := 0
	xpos := 0
	ypos := 0
	for scanner.Scan() {
		line := scanner.Text()
		if part == 0 && line == "" {
			part = 1
			continue
		}

		if part == 0 {
			parts := []string{}
			for index, c := range line {
				parts = append(parts, string(c))
				if string(c) == "@" {
					xpos = index
					ypos = len(warehouse)
				}
			}

			warehouse = append(warehouse, parts)
		} else {
			for _, c := range line {
				movements = append(movements, string(c))
			}
		}
	}

	for _, movement := range movements {
		fmt.Println(movement)
		move(warehouse, &xpos, &ypos, movement)
		/*for j := 0; j < len(warehouse); j++ {
			for i := 0; i < len(warehouse[j]); i++ {
				fmt.Print(warehouse[j][i])
			}
			fmt.Println()
		}

		fmt.Println()*/
	}

	question1 := 0
	for j := 0; j < len(warehouse); j++ {
		for i := 0; i < len(warehouse[j]); i++ {
			if warehouse[j][i] == "O" {
				question1 += 100*j + i
			}
		}
	}

	fmt.Println("Question 1:", question1)
}

func move(warehouse [][]string, xpos *int, ypos *int, movement string) {
	dir := direction(movement)
	moves := []Movement{}
	moveable := true

	tx := *xpos
	ty := *ypos
	for true {
		if warehouse[ty][tx] == "#" {
			moveable = false
			break
		}
		if warehouse[ty][tx] == "." {
			moveable = true
			break
		}

		moves = append(moves, Movement{Coordinate{tx, ty}, dir})
		tx = tx + dir.x
		ty = ty + dir.y
	}

	if moveable {
		slices.Reverse(moves)

		for _, step := range moves {
			warehouse[step.position.y+step.direction.y][step.position.x+step.direction.x] = warehouse[step.position.y][step.position.x]
		}

		warehouse[*ypos][*xpos] = "."
		*xpos = *xpos + dir.x
		*ypos = *ypos + dir.y
	}
}

func direction(movement string) Coordinate {
	switch movement {
	case "<":
		return Coordinate{-1, 0}
	case ">":
		return Coordinate{1, 0}
	case "^":
		return Coordinate{0, -1}
	case "v":
		return Coordinate{0, 1}
	}

	return Coordinate{0, 0}
}
