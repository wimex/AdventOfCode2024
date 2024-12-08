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

	antinodes1 := []coord{}
	antinodes2 := []coord{}
	for _, v := range coords {
		combos := slices.DeleteFunc(combinations.All(v), func(c []coord) bool { return len(c) < 2 })
		for _, c := range combos {
			for _, a := range generateAntinodes(c[0], c[1], width, height, false) {
				if !slices.Contains(antinodes1, a) {
					antinodes1 = append(antinodes1, a)
				}
			}

			for _, a := range generateAntinodes(c[0], c[1], width, height, true) {
				if !slices.Contains(antinodes2, a) {
					antinodes2 = append(antinodes2, a)
				}
			}
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if slices.Contains(antinodes2, coord{j, i}) {
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
	fmt.Println("Question 1:", len(antinodes1))
	fmt.Println("Question 2:", len(antinodes2))
}

func generateAntinodes(a, b coord, width, height int, harmonics bool) []coord {
	distance := coord{b.x - a.x, b.y - a.y}
	slope := getSlope(a, b)
	results := []coord{}

	px1 := a.x
	px2 := b.x
	for true {
		lx := px1 - distance.x
		ly := float64(a.y) - (slope * float64(distance.x))

		rx := px2 + distance.x
		ry := float64(b.y) + (slope * float64(distance.x))

		if ly != float64(int(ly)) || ry != float64(int(ry)) {
			fmt.Println("Not an integer")
			return []coord{}
		}

		added := false
		if lx >= 0 && lx < width && ly >= 0 && int(ly) < height {
			results = append(results, coord{lx, int(ly)})
			added = true
		}

		if rx >= 0 && rx < width && ry >= 0 && int(ry) < height {
			results = append(results, coord{rx, int(ry)})
			added = true
		}

		if !harmonics || !added {
			return results
		}

		px1 = lx
		px2 = rx
	}

	return results
}
func getSlope(a, b coord) float64 {
	if b.x == a.x {
		return 0
	}

	return float64(b.y-a.y) / float64(b.x-a.x)
}
