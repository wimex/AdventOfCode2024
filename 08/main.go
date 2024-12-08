package main

import (
	"bufio"
	"fmt"
	"github.com/mxschmitt/golang-combinations"
	"os"
	"slices"
)

type coord struct {
	x int
	y int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	width := 0
	height := 0

	coords := make(map[string][]coord)
	for scanner.Scan() {
		line := scanner.Text()
		width = len(line)

		for x, v := range line {
			t := string(v)
			if t == "." {
				continue
			}

			coords[t] = append(coords[t], coord{x, height})
		}

		height++
	}

	antinodes := []coord{}
	for _, v := range coords {
		combos := slices.DeleteFunc(combinations.All(v), func(c []coord) bool { return len(c) < 2 })
		for _, c := range combos {
			distance := coord{c[1].x - c[0].x, c[1].y - c[0].y}
			slope := getSlope(c[0], c[1])

			lx := c[0].x - distance.x
			ly := float64(c[0].y) - (slope * float64(distance.x))

			rx := c[1].x + distance.x
			ry := float64(c[1].y) + (slope * float64(distance.x))

			if ly != float64(int(ly)) || ry != float64(int(ry)) {
				fmt.Println("Not an integer")
				return
			}

			if lx >= 0 && lx < width && ly >= 0 && int(ly) < height && !slices.Contains(antinodes, coord{lx, int(ly)}) {
				antinodes = append(antinodes, coord{lx, int(ly)})
			}

			if rx >= 0 && rx < width && ry >= 0 && int(ry) < height && !slices.Contains(antinodes, coord{rx, int(ry)}) {
				antinodes = append(antinodes, coord{rx, int(ry)})
			}
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if slices.Contains(antinodes, coord{j, i}) {
				fmt.Print("#")
				continue
			}

			for k, v := range coords {
				if slices.Contains(v, coord{j, i}) {
					fmt.Print(k)
					continue
				}
			}

			fmt.Print(".")
		}

		fmt.Println()
	}

	fmt.Println()
	fmt.Println("Question 1:", len(antinodes))
}

func getSlope(a, b coord) float64 {
	if b.x == a.x {
		return 0
	}

	return float64(b.y-a.y) / float64(b.x-a.x)
}
