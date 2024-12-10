package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"strconv"
)

type coord struct {
	x int
	y int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	coords := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		coord := []int{}

		for _, v := range line {
			num, _ := strconv.Atoi(string(v))
			coord = append(coord, num)
		}

		coords = append(coords, coord)
	}

	question1 := [][]coord{}
	question2 := [][]coord{}
	for j := 0; j < len(coords); j++ {
		for i := 0; i < len(coords[j]); i++ {
			if coords[j][i] != 0 {
				continue
			}

			paths, valid := hike(coords, coord{i, j}, 0)
			unique := lo.Uniq(paths)
			if valid {
				question1 = append(question1, unique)
				question2 = append(question2, paths)
			}
		}
	}

	fmt.Println("Question 1:", lo.SumBy(question1, func(c []coord) int { return len(c) }))
	fmt.Println("Question 2:", lo.SumBy(question2, func(c []coord) int { return len(c) }))
}

func hike(coords [][]int, position coord, height int) ([]coord, bool) {
	if position.y < 0 || position.y >= len(coords) || position.x < 0 || position.x >= len(coords[position.y]) {
		return []coord{position}, false
	}

	if coords[position.y][position.x] != height {
		return []coord{position}, false
	}

	if height == 9 {
		return []coord{position}, true
	}

	score1, valid1 := hike(coords, coord{position.x - 1, position.y}, height+1)
	score2, valid2 := hike(coords, coord{position.x + 1, position.y}, height+1)
	score3, valid3 := hike(coords, coord{position.x, position.y - 1}, height+1)
	score4, valid4 := hike(coords, coord{position.x, position.y + 1}, height+1)

	scores := []coord{}
	if valid1 {
		scores = append(scores, score1...)
	}
	if valid2 {
		scores = append(scores, score2...)
	}
	if valid3 {
		scores = append(scores, score3...)
	}
	if valid4 {
		scores = append(scores, score4...)
	}

	return scores, true
}
