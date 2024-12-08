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
		// Take two antennas and generate all possible combinations
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

	// Every antenna is now an antinode as well
	for _, v := range coords {
		for _, c := range v {
			if !slices.Contains(antinodes2, c) {
				antinodes2 = append(antinodes2, c)
			}
		}
	}

	fmt.Println()
	fmt.Println("Question 1:", len(antinodes1))
	fmt.Println("Question 2:", len(antinodes2))
}

func generateAntinodes(a, b coord, width, height int, harmonics bool) []coord {
	distance := coord{b.x - a.x, b.y - a.y} // Distance between the two antennas
	slope := getSlope(a, b)                 // Based on the distance, calculate the slope
	results := []coord{}

	px1 := a.x
	px2 := b.x
	py1 := a.y
	py2 := b.y
	for true {
		// On the X-axis, step to the left, calculate the Y-axis based on the new X and the slope
		lx := px1 - distance.x
		ly := float64(py1) - (slope * float64(distance.x))

		rx := px2 + distance.x
		ry := float64(py2) + (slope * float64(distance.x))

		// Even though the slope is a float, the resulting coordinates should be integers
		if ly != float64(int(ly)) || ry != float64(int(ry)) {
			fmt.Println("Not an integer")
			return []coord{}
		}

		res1 := coord{lx, int(ly)}
		res2 := coord{rx, int(ry)}

		added := false
		if res1.x >= 0 && res1.x < width && res1.y >= 0 && res1.y < height {
			results = append(results, res1)
			added = true
		}

		if res2.x >= 0 && res2.x < width && res2.y >= 0 && res2.y < height {
			results = append(results, coord{rx, int(ry)})
			added = true
		}

		if !harmonics || !added {
			return results
		}

		px1 = res1.x
		py1 = res1.y

		px2 = res2.x
		py2 = res2.y
	}

	return results
}

func getSlope(a, b coord) float64 {
	if b.x == a.x {
		return 0
	}

	return float64(b.y-a.y) / float64(b.x-a.x)
}
