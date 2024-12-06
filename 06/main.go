package main

import (
	"bufio"
	"errors"
	"fmt"
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
	for col, _ := range fields {
		for row, _ := range fields[col] {
			if fields[row][col] == "^" {
				guard[0] = row
				guard[1] = col
			}
		}
	}

	steps := make([][]int, 0)
	//steps = append(steps, []int{guard[0], guard[1]})

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

	println("Question 1:", len(distinct))
	fmt.Println(steps)
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
