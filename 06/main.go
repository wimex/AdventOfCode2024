package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fields := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		fields = append(fields, strings.Split(line, ""))
	}

	guard := []int{0, 0}
	orig := []int{0, 0}
	for col, _ := range fields {
		for row, _ := range fields[col] {
			if fields[row][col] == "^" {
				guard[0] = row
				guard[1] = col
				orig[0] = row
				orig[1] = col
			}
		}
	}

	steps := make([][]int, 0)
	steps = append(steps, []int{guard[0], guard[1]})

	direction := []int{-1, 0}
	for true {
		next := []int{guard[0] + direction[0], guard[1] + direction[1]}

		if next[0] < 0 || next[0] >= len(fields) || next[1] < 0 || next[1] >= len(fields[next[0]]) {
			break
		}

		if fields[next[0]][next[1]] == "#" {
			direction, _ = turn(direction)
			continue
		}

		steps = append(steps, next)
		guard = next
	}

	distinct := [][]int{}
	for _, step := range steps {
		found := false
		for _, dist := range distinct {
			if step[0] == dist[0] && step[1] == dist[1] {
				found = true
				break
			}
		}

		if !found {
			distinct = append(distinct, step)
		}
	}

	question2 := 0
	for y, row := range fields {
		for x, _ := range row {
			if fields[y][x] == "#" || fields[y][x] == "^" {
				continue
			}

			fields[y][x] = "#"
			guard = orig
			question2 += checkLoop(fields, guard, []int{-1, 0})
			fields[y][x] = "."
		}
	}

	println("Question 1:", len(distinct))
	println("Question 2:", question2)
}

func turn(direction []int) ([]int, error) {
	if direction[0] == 0 && direction[1] == -1 {
		return []int{-1, 0}, nil
	} else if direction[0] == -1 && direction[1] == 0 {
		return []int{0, 1}, nil
	} else if direction[0] == 0 && direction[1] == 1 {
		return []int{1, 0}, nil
	} else if direction[0] == 1 && direction[1] == 0 {
		return []int{0, -1}, nil
	}

	return nil, errors.New("Invalid direction")
}

func checkLoop(fields [][]string, guard []int, direction []int) int {
	// This seems like a reasonable exit condition to assume that this is a loop. Very demure, very mindful.
	// There are probably better solutions, I already get the correct answer around 10000 iterations.
	exit := len(fields) * len(fields[0])
	for i := 0; i < exit; i++ {
		next := []int{guard[0] + direction[0], guard[1] + direction[1]}

		if next[0] < 0 || next[0] >= len(fields) || next[1] < 0 || next[1] >= len(fields[next[0]]) {
			return 0
		}

		if fields[next[0]][next[1]] == "#" {
			direction, _ = turn(direction)
			continue
		}

		guard = next
	}

	return 1
}
