package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type block struct {
	ident  int
	start  int
	length int
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

	files1 := compact1(files)
	files2 := compact2(files)

	fmt.Println("Question 1:", checksum(files1))
	fmt.Println("Question 2:", checksum(files2))
}

func compact1(files []int) []int {
	result := []int{}
	for _, data := range files {
		result = append(result, data)
	}

	free := 0
	for head := len(result) - 1; head >= 0; head-- {
		if result[head] != -1 {
			for result[free] != -1 {
				free++
			}

			if free < head {
				result[free] = result[head]
				result[head] = -1
			}
		}
	}

	return result
}

func compact2(input []int) []int {
	// Check the input and collect continous blocks of the same value
	surface := []block{}
	for i := 0; i < len(input)-1; i++ {
		length := 0
		for j := i; j < len(input) && input[i] == input[j]; j++ {
			length++
		}

		surface = append(surface, block{input[i], i, length})
		i += length - 1
	}

	result := []int{}
	head := len(surface) - 1
	for true {
		//Find the next file from the right
		for head >= 0 && surface[head].ident == -1 {
			head--
		}

		//If no more files are found, we have finished
		if head <= 0 {
			break
		}

		//Find a free space from the left that fits the file
		found := false
		free := 0
		for ; free < head; free++ { //Files are only moved to the left, not to the right
			if surface[free].ident == -1 && surface[free].length >= surface[head].length {
				found = true
				break
			}
		}

		//If a suitable space was found, move the file
		if found {
			//Insert a new remaining free space after the current free space
			if surface[head].length != surface[free].length {
				surface = slices.Insert(surface, free+1, block{-1, surface[free].start + surface[head].length, surface[free].length - surface[head].length})
				head++ //A new block has been inserted so the head moves
			}

			//Move the file
			surface[free] = block{surface[head].ident, surface[free].start, surface[head].length}

			//Mark the file as free space
			surface[head].ident = -1
		} else {
			//The file can not be fit anywere, go to the next file
			head--
		}

		//Merge free spaces that are next to each other
		for i := 0; i < len(surface)-2; i++ {
			if surface[i].ident == -1 && surface[i+1].ident == -1 {
				surface[i].length += surface[i+1].length
				surface = slices.Delete(surface, i+1, i+2)
				head-- //The indexes have changed, move head to the left
				i--    //Step back so if a page was merged, it will be checked again
			}
		}
	}

	//Convert the surface to a file list
	for _, block := range surface {
		for i := 0; i < block.length; i++ {
			result = append(result, block.ident)
		}
	}

	return result
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
