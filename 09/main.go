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
	files := []string{}
	ident := 0
	for index, value := range line {
		length, _ := strconv.Atoi(string(value))
		for i := 0; i < length; i++ {
			if index%2 == 0 {
				data := strconv.Itoa(ident)
				files = append(files, data)
			} else {
				data := "."
				files = append(files, data)
			}
		}

		if index%2 == 0 {
			if ident == 9 {
				ident = 0
			} else {
				ident++
			}
		}
	}

	free := 0
	for head := len(files) - 1; free <= head; head-- {
		if files[head] != "." {
			for files[free] != "." {
				free++
			}

			files[free] = files[head]
			files[head] = "."

			free++
		}
	}

	fmt.Println(files)
	fmt.Println("Question 1:", checksum(files))
}

func checksum(files []string) int {
	result := 0
	for index, value := range files {
		if value == "." {
			continue
		}

		ident, _ := strconv.Atoi(value)
		result += ident * index
	}

	return result
}
