package main

import (
	"bufio"
	"fmt"
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
	scanner.Scan()

	line := scanner.Text()
	files := []int{}
	ident := 0
	for index, value := range line {
		length, _ := strconv.Atoi(string(value))
		for i := 0; i < length; i++ {
			if index%2 == 0 {
				files = append(files, ident)
			} else {
				files = append(files, -1)
			}
		}

		if index%2 == 0 {
			ident++
		}
	}

	free := 0
	for head := len(files) - 1; head >= 0; head-- {
		if files[head] != -1 {
			for files[free] != -1 {
				free++
			}

			if free < head {
				files[free] = files[head]
				files[head] = -1
			}
		}
	}

	fmt.Println(files)
	fmt.Println("Question 1:", checksum(files))
}

func checksum(files []int) int {
	result := 0
	for index, value := range files {
		if value == -1 {
			continue
		}

		result += value * index
	}

	return result
}
